package manufacturing

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Enterprise SCM Production Queries (Fase 3)
func GetPaginatedProductions(offset, limit int, search, state string, productID uint) ([]MrpProduction, int64, error) {
	var list []MrpProduction
	var total int64

	query := config.DB.Model(&MrpProduction{}).
		Preload("Product.ProductTemplate").
		Preload("Bom.BomLines.Product.ProductTemplate").
		Preload("Warehouse")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("supply_chain.mrp_productions.name ILIKE ? OR product_id IN (SELECT p.id FROM supply_chain.products p JOIN supply_chain.product_templates pt ON p.product_template_id = pt.id WHERE pt.name ILIKE ? OR p.default_code ILIKE ?)", s, s, s)
	}

	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}

	if productID > 0 {
		query = query.Where("product_id = ?", productID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetMrpSummary() (MrpSummary, error) {
	var summary MrpSummary

	config.DB.Model(&MrpProduction{}).Count(&summary.TotalMOCount)
	config.DB.Model(&MrpProduction{}).Where("state IN ('confirmed', 'progress')").Count(&summary.MOInProgress)
	config.DB.Model(&MrpProduction{}).Where("state = 'done'").Count(&summary.MODoneCount)
	config.DB.Model(&MrpWorkcenter{}).Count(&summary.TotalWorkcenters)
	config.DB.Model(&MrpBom{}).Count(&summary.TotalBoms)

	return summary, nil
}

func CreateProductionWithSequence(req CreateMORequest) (*MrpProduction, error) {
	var mo MrpProduction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&MrpProduction{}).Count(&count)
		moName := fmt.Sprintf("WH/MO/%s/%04d", time.Now().Format("2006"), count+1)

		var plannedTime *time.Time
		if req.DatePlanned != nil && *req.DatePlanned != "" {
			if t, err := time.Parse(time.RFC3339, *req.DatePlanned); err == nil {
				plannedTime = &t
			} else if t, err := time.Parse("2006-01-02", *req.DatePlanned); err == nil {
				plannedTime = &t
			}
		}

		mo = MrpProduction{
			Name:        moName,
			ProductID:   req.ProductID,
			BomID:       req.BomID,
			ProductQty:  req.ProductQty,
			WarehouseID: req.WarehouseID,
			DatePlanned: plannedTime,
			Notes:       req.Notes,
			State:       "draft",
		}

		return tx.Create(&mo).Error
	})

	if err != nil {
		return nil, err
	}

	return GetMrpProductionByID(mo.ID)
}

func ConfirmProduction(id uint) (*MrpProduction, error) {
	var mo MrpProduction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Bom.BomLines.Product").First(&mo, id).Error; err != nil {
			return err
		}

		if mo.State != "draft" {
			return errors.New("hanya SPK/MO berstatus draft yang dapat dikonfirmasi")
		}

		mo.State = "confirmed"
		return tx.Save(&mo).Error
	})

	if err != nil {
		return nil, err
	}

	return GetMrpProductionByID(id)
}

func StartProduction(id uint) (*MrpProduction, error) {
	var mo MrpProduction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&mo, id).Error; err != nil {
			return err
		}

		if mo.State != "confirmed" {
			return errors.New("hanya SPK/MO berstatus confirmed yang dapat dimulai produksinya")
		}

		now := time.Now()
		mo.State = "progress"
		mo.DateStart = &now
		return tx.Save(&mo).Error
	})

	if err != nil {
		return nil, err
	}

	return GetMrpProductionByID(id)
}

func FinishProduction(id uint) (*MrpProduction, error) {
	var mo MrpProduction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Preload("Bom.BomLines.Product").Preload("Product.ProductTemplate").First(&mo, id).Error; err != nil {
			return err
		}

		if mo.State != "progress" && mo.State != "confirmed" {
			return errors.New("hanya SPK/MO berstatus progress atau confirmed yang dapat diselesaikan")
		}

		// 1. Backflush: Kurangi stok bahan baku komponen jika ada BOM
		if mo.Bom != nil && len(mo.Bom.BomLines) > 0 {
			bomBaseQty := mo.Bom.Quantity
			if bomBaseQty <= 0 {
				bomBaseQty = 1
			}

			for _, line := range mo.Bom.BomLines {
				requiredQty := (line.Quantity / bomBaseQty) * mo.ProductQty
				// Kurangi stok bahan baku
				if err := tx.Model(&inventory.Product{}).
					Where("id = ?", line.ProductID).
					Update("stock_qty", gorm.Expr("stock_qty - ?", requiredQty)).Error; err != nil {
					return err
				}

				// Catat StockMove bahan baku keluar
				compMove := inventory.StockMove{
					Name:         fmt.Sprintf("Konsumsi Bahan Baku: %s untuk %s", line.Product.DefaultCode, mo.Name),
					ProductID:    line.ProductID,
					Quantity:     requiredQty,
					QuantityDone: requiredQty,
					State:        "done",
				}
				if err := tx.Create(&compMove).Error; err != nil {
					return err
				}
			}
		}

		// 2. Tambah stok produk jadi
		if err := tx.Model(&inventory.Product{}).
			Where("id = ?", mo.ProductID).
			Update("stock_qty", gorm.Expr("stock_qty + ?", mo.ProductQty)).Error; err != nil {
			return err
		}

		// 3. Catat StockMove & StockValuationLayer produk jadi masuk
		fgMove := inventory.StockMove{
			Name:         fmt.Sprintf("Hasil Produksi Jadi: %s (%s)", mo.Name, mo.Product.ProductTemplate.Name),
			ProductID:    mo.ProductID,
			Quantity:     mo.ProductQty,
			QuantityDone: mo.ProductQty,
			State:        "done",
		}
		if err := tx.Create(&fgMove).Error; err != nil {
			return err
		}

		valLayer := inventory.StockValuationLayer{
			ProductID:   mo.ProductID,
			Quantity:    mo.ProductQty,
			UnitCost:    mo.Product.ProductTemplate.StandardPrice,
			Value:       mo.ProductQty * mo.Product.ProductTemplate.StandardPrice,
			Description: fmt.Sprintf("Produksi Selesai: %s", mo.Name),
		}
		if err := tx.Create(&valLayer).Error; err != nil {
			return err
		}

		// 4. Update status MO menjadi done
		now := time.Now()
		mo.State = "done"
		mo.DateFinished = &now
		return tx.Save(&mo).Error
	})

	if err != nil {
		return nil, err
	}

	return GetMrpProductionByID(id)
}

func CancelProduction(id uint) (*MrpProduction, error) {
	var mo MrpProduction
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&mo, id).Error; err != nil {
			return err
		}

		if mo.State == "done" {
			return errors.New("SPK/MO yang sudah selesai tidak dapat dibatalkan")
		}

		mo.State = "cancel"
		return tx.Save(&mo).Error
	})

	if err != nil {
		return nil, err
	}

	return GetMrpProductionByID(id)
}

// Enterprise BOM Queries & Handlers
func GetPaginatedBoms(offset, limit int, search string) ([]MrpBom, int64, error) {
	var list []MrpBom
	var total int64

	query := config.DB.Model(&MrpBom{}).
		Preload("Product.ProductTemplate").
		Preload("BomLines.Product.ProductTemplate")

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("code ILIKE ? OR product_id IN (SELECT p.id FROM supply_chain.products p JOIN supply_chain.product_templates pt ON p.product_template_id = pt.id WHERE pt.name ILIKE ? OR p.default_code ILIKE ?)", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func CreateBomWithLines(req CreateBomRequest) (*MrpBom, error) {
	var bom MrpBom
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		bomType := req.Type
		if bomType == "" {
			bomType = "normal"
		}

		bom = MrpBom{
			ProductID: req.ProductID,
			Code:      req.Code,
			Type:      bomType,
			Quantity:  req.Quantity,
		}
		if err := tx.Create(&bom).Error; err != nil {
			return err
		}

		for _, l := range req.Lines {
			line := MrpBomLine{
				BomID:     bom.ID,
				ProductID: l.ProductID,
				Quantity:  l.Quantity,
			}
			if err := tx.Create(&line).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetMrpBomByID(bom.ID)
}

// Standard CRUD Wrappers
func CreateMrpProduction(data *MrpProduction) error {
	return config.DB.Create(data).Error
}

func GetAllMrpProduction() ([]MrpProduction, error) {
	var list []MrpProduction
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetMrpProductionByID(id uint) (*MrpProduction, error) {
	var data MrpProduction
	err := config.DB.
		Preload("Product.ProductTemplate").
		Preload("Bom.BomLines.Product.ProductTemplate").
		Preload("Warehouse").
		First(&data, id).Error
	return &data, err
}

func UpdateMrpProduction(data *MrpProduction) error {
	return config.DB.Save(data).Error
}

func DeleteMrpProduction(id uint) error {
	return config.DB.Delete(&MrpProduction{}, id).Error
}

func CreateMrpWorkcenter(data *MrpWorkcenter) error { return config.DB.Create(data).Error }
func GetAllMrpWorkcenter() ([]MrpWorkcenter, error) {
	var list []MrpWorkcenter
	err := config.DB.Order("id ASC").Find(&list).Error
	return list, err
}
func GetMrpWorkcenterByID(id uint) (*MrpWorkcenter, error) {
	var data MrpWorkcenter
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateMrpWorkcenter(data *MrpWorkcenter) error { return config.DB.Save(data).Error }
func DeleteMrpWorkcenter(id uint) error             { return config.DB.Delete(&MrpWorkcenter{}, id).Error }

func CreateMrpBom(data *MrpBom) error { return config.DB.Create(data).Error }
func GetAllMrpBom() ([]MrpBom, error) {
	var list []MrpBom
	err := config.DB.Preload("Product.ProductTemplate").Preload("BomLines.Product.ProductTemplate").Find(&list).Error
	return list, err
}
func GetMrpBomByID(id uint) (*MrpBom, error) {
	var data MrpBom
	err := config.DB.Preload("Product.ProductTemplate").Preload("BomLines.Product.ProductTemplate").First(&data, id).Error
	return &data, err
}
func UpdateMrpBom(data *MrpBom) error { return config.DB.Save(data).Error }
func DeleteMrpBom(id uint) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		tx.Where("bom_id = ?", id).Delete(&MrpBomLine{})
		return tx.Delete(&MrpBom{}, id).Error
	})
}

func CreateMrpBomLine(data *MrpBomLine) error { return config.DB.Create(data).Error }
func GetAllMrpBomLine() ([]MrpBomLine, error) {
	var list []MrpBomLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetMrpBomLineByID(id uint) (*MrpBomLine, error) {
	var data MrpBomLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateMrpBomLine(data *MrpBomLine) error { return config.DB.Save(data).Error }
func DeleteMrpBomLine(id uint) error          { return config.DB.Delete(&MrpBomLine{}, id).Error }

func CreateMrpBomByproduct(data *MrpBomByproduct) error { return config.DB.Create(data).Error }
func GetAllMrpBomByproduct() ([]MrpBomByproduct, error) {
	var list []MrpBomByproduct
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetMrpBomByproductByID(id uint) (*MrpBomByproduct, error) {
	var data MrpBomByproduct
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateMrpBomByproduct(data *MrpBomByproduct) error { return config.DB.Save(data).Error }
func DeleteMrpBomByproduct(id uint) error               { return config.DB.Delete(&MrpBomByproduct{}, id).Error }

func CreateMrpWorkorder(data *MrpWorkorder) error { return config.DB.Create(data).Error }
func GetAllMrpWorkorder() ([]MrpWorkorder, error) {
	var list []MrpWorkorder
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetMrpWorkorderByID(id uint) (*MrpWorkorder, error) {
	var data MrpWorkorder
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateMrpWorkorder(data *MrpWorkorder) error { return config.DB.Save(data).Error }
func DeleteMrpWorkorder(id uint) error            { return config.DB.Delete(&MrpWorkorder{}, id).Error }

