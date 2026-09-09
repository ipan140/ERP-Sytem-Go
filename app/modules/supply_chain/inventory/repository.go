package inventory

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"ERP-System/config"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProduct(data *Product) error {
	return config.DB.Create(data).Error
}

func GetAllProduct() ([]Product, error) {
	var list []Product
	err := config.DB.Preload(clause.Associations).
		Preload("ProductTemplate.Category").
		Preload("ProductTemplate.UoM").
		Find(&list).Error
	return list, err
}

func GetPaginatedProducts(offset, limit int, search string, categoryID, warehouseID uint, stockStatus string) ([]Product, int64, error) {
	var list []Product
	var total int64

	query := config.DB.Model(&Product{}).
		Joins("LEFT JOIN supply_chain.product_templates ON supply_chain.product_templates.id = supply_chain.products.product_template_id")

	if search != "" {
		searchPattern := "%" + strings.ToLower(search) + "%"
		query = query.Where(
			"LOWER(supply_chain.products.default_code) LIKE ? OR LOWER(supply_chain.products.barcode) LIKE ? OR LOWER(supply_chain.product_templates.name) LIKE ?",
			searchPattern, searchPattern, searchPattern,
		)
	}

	if categoryID > 0 {
		query = query.Where("supply_chain.product_templates.category_id = ?", categoryID)
	}

	if warehouseID > 0 {
		query = query.Where(
			"supply_chain.products.id IN (SELECT sq.product_id FROM supply_chain.stock_quants sq JOIN supply_chain.stock_locations sl ON sl.id = sq.location_id WHERE sl.warehouse_id = ?)",
			warehouseID,
		)
	}

	if stockStatus != "" {
		switch stockStatus {
		case "out_of_stock":
			query = query.Where("supply_chain.products.stock_qty <= 0")
		case "low_stock":
			query = query.Where("supply_chain.products.stock_qty > 0 AND supply_chain.products.stock_qty <= 10")
		case "in_stock":
			query = query.Where("supply_chain.products.stock_qty > 10")
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.
		Preload("ProductTemplate").
		Preload("ProductTemplate.Category").
		Preload("ProductTemplate.UoM").
		Order("supply_chain.products.id DESC").
		Offset(offset).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}

func GetInventorySummary() (InventorySummary, error) {
	var summary InventorySummary

	config.DB.Model(&Product{}).Count(&summary.TotalSKU)

	type AggResult struct {
		OnHand   float64
		Reserved float64
	}
	var agg AggResult
	config.DB.Model(&Product{}).Select("COALESCE(SUM(stock_qty), 0) as on_hand, COALESCE(SUM(reserved_qty), 0) as reserved").Scan(&agg)
	summary.TotalOnHand = agg.OnHand
	summary.TotalReserved = agg.Reserved

	config.DB.Model(&Product{}).Where("stock_qty <= 0").Count(&summary.OutOfStockCount)
	config.DB.Model(&Product{}).Where("stock_qty > 0 AND stock_qty <= 10").Count(&summary.LowStockCount)

	var valResult struct {
		TotalVal float64
	}
	config.DB.Table("supply_chain.products p").
		Select("COALESCE(SUM(p.stock_qty * COALESCE(NULLIF(pt.standard_price, 0), pt.list_price, 0)), 0) as total_val").
		Joins("LEFT JOIN supply_chain.product_templates pt ON pt.id = p.product_template_id").
		Scan(&valResult)
	summary.TotalValuation = valResult.TotalVal

	return summary, nil
}

func ApplyStockAdjustment(req StockAdjustmentRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var prod Product
		if err := tx.First(&prod, req.ProductID).Error; err != nil {
			return err
		}

		diff := req.CountedQty - prod.StockQty
		prod.StockQty = req.CountedQty
		if err := tx.Save(&prod).Error; err != nil {
			return err
		}

		var targetLocID *uint
		if req.LocationID > 0 {
			targetLocID = &req.LocationID
		} else {
			var loc StockLocation
			if err := tx.First(&loc).Error; err == nil {
				targetLocID = &loc.ID
			}
		}

		// Record stock move for audit trail
		docNum := fmt.Sprintf("INV-ADJ/%s/%04d", time.Now().Format("20060102"), prod.ID)
		move := StockMove{
			Name:           fmt.Sprintf("Penyesuaian Fisik (Opname) - %s (Selisih: %.2f)", req.Reason, diff),
			ProductID:      prod.ID,
			Quantity:       req.CountedQty,
			QuantityDone:   req.CountedQty,
			LocationID:     targetLocID,
			LocationDestID: targetLocID,
			State:          "done",
			CreatedAt:      time.Now(),
		}
		if err := tx.Create(&move).Error; err != nil {
			return err
		}

		// Record valuation layer
		var pt ProductTemplate
		if err := tx.First(&pt, prod.ProductTemplateID).Error; err == nil {
			unitCost := pt.StandardPrice
			if unitCost == 0 {
				unitCost = pt.ListPrice
			}
			valLayer := StockValuationLayer{
				ProductID:   prod.ID,
				Quantity:    diff,
				UnitCost:    unitCost,
				Value:       diff * unitCost,
				Description: fmt.Sprintf("Penyesuaian Stok: %s (%s)", docNum, req.Reason),
				CreatedAt:   time.Now(),
			}
			tx.Create(&valLayer)
		}

		return nil
	})
}

func ApplyInternalTransfer(req InternalTransferRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var prod Product
		if err := tx.First(&prod, req.ProductID).Error; err != nil {
			return err
		}

		if prod.StockQty < req.Quantity {
			return errors.New("stok produk tidak mencukupi untuk ditransfer")
		}

		transferCode := fmt.Sprintf("WH/INT/%s/%d", time.Now().Format("20060102150405"), prod.ID)
		var srcLoc, destLoc StockLocation
		var srcLocID, destLocID *uint
		if req.SourceWarehouseID > 0 && tx.Where("warehouse_id = ?", req.SourceWarehouseID).First(&srcLoc).Error == nil {
			srcLocID = &srcLoc.ID
		}
		if req.DestWarehouseID > 0 && tx.Where("warehouse_id = ?", req.DestWarehouseID).First(&destLoc).Error == nil {
			destLocID = &destLoc.ID
		}

		picking := StockPicking{
			Name:           transferCode,
			LocationID:     srcLocID,
			LocationDestID: destLocID,
			State:          "done",
			ScheduledDate: time.Now(),
			CreatedAt:     time.Now(),
		}
		if err := tx.Create(&picking).Error; err != nil {
			return err
		}

		move := StockMove{
			Name:           fmt.Sprintf("Transfer Internal Antar-Gudang %s (Catatan: %s)", transferCode, req.Notes),
			PickingID:      &picking.ID,
			ProductID:      prod.ID,
			Quantity:       req.Quantity,
			QuantityDone:   req.Quantity,
			LocationID:     srcLocID,
			LocationDestID: destLocID,
			State:          "done",
			CreatedAt:      time.Now(),
		}
		if err := tx.Create(&move).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetProductByID(id uint) (*Product, error) {
	var data Product
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateProduct(data *Product) error {
	return config.DB.Save(data).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&Product{}, id).Error
}

func CreateProductCategory(data *ProductCategory) error { return config.DB.Create(data).Error }
func GetAllProductCategory() ([]ProductCategory, error) {
	var list []ProductCategory
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetProductCategoryByID(id uint) (*ProductCategory, error) {
	var data ProductCategory
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateProductCategory(data *ProductCategory) error { return config.DB.Save(data).Error }
func DeleteProductCategory(id uint) error               { return config.DB.Delete(&ProductCategory{}, id).Error }

func CreateUoMCategory(data *UoMCategory) error { return config.DB.Create(data).Error }
func GetAllUoMCategory() ([]UoMCategory, error) {
	var list []UoMCategory
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetUoMCategoryByID(id uint) (*UoMCategory, error) {
	var data UoMCategory
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateUoMCategory(data *UoMCategory) error { return config.DB.Save(data).Error }
func DeleteUoMCategory(id uint) error           { return config.DB.Delete(&UoMCategory{}, id).Error }

func CreateUoM(data *UoM) error { return config.DB.Create(data).Error }
func GetAllUoM() ([]UoM, error) { var list []UoM; err := config.DB.Preload(clause.Associations).Find(&list).Error; return list, err }
func GetUoMByID(id uint) (*UoM, error) {
	var data UoM
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateUoM(data *UoM) error { return config.DB.Save(data).Error }
func DeleteUoM(id uint) error   { return config.DB.Delete(&UoM{}, id).Error }

func CreateProductTemplate(data *ProductTemplate) error { return config.DB.Create(data).Error }
func GetAllProductTemplate() ([]ProductTemplate, error) {
	var list []ProductTemplate
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetProductTemplateByID(id uint) (*ProductTemplate, error) {
	var data ProductTemplate
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateProductTemplate(data *ProductTemplate) error { return config.DB.Save(data).Error }
func DeleteProductTemplate(id uint) error               { return config.DB.Delete(&ProductTemplate{}, id).Error }

func CreateProductAttribute(data *ProductAttribute) error { return config.DB.Create(data).Error }
func GetAllProductAttribute() ([]ProductAttribute, error) {
	var list []ProductAttribute
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetProductAttributeByID(id uint) (*ProductAttribute, error) {
	var data ProductAttribute
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateProductAttribute(data *ProductAttribute) error { return config.DB.Save(data).Error }
func DeleteProductAttribute(id uint) error                { return config.DB.Delete(&ProductAttribute{}, id).Error }

func CreateProductAttributeValue(data *ProductAttributeValue) error {
	return config.DB.Create(data).Error
}
func GetAllProductAttributeValue() ([]ProductAttributeValue, error) {
	var list []ProductAttributeValue
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetProductAttributeValueByID(id uint) (*ProductAttributeValue, error) {
	var data ProductAttributeValue
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateProductAttributeValue(data *ProductAttributeValue) error {
	return config.DB.Save(data).Error
}
func DeleteProductAttributeValue(id uint) error {
	return config.DB.Delete(&ProductAttributeValue{}, id).Error
}

func CreateStockWarehouse(data *StockWarehouse) error { return config.DB.Create(data).Error }
func GetAllStockWarehouse() ([]StockWarehouse, error) {
	var list []StockWarehouse
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockWarehouseByID(id uint) (*StockWarehouse, error) {
	var data StockWarehouse
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockWarehouse(data *StockWarehouse) error { return config.DB.Save(data).Error }
func DeleteStockWarehouse(id uint) error              { return config.DB.Delete(&StockWarehouse{}, id).Error }

func CreateStockLocation(data *StockLocation) error { return config.DB.Create(data).Error }
func GetAllStockLocation() ([]StockLocation, error) {
	var list []StockLocation
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockLocationByID(id uint) (*StockLocation, error) {
	var data StockLocation
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockLocation(data *StockLocation) error { return config.DB.Save(data).Error }
func DeleteStockLocation(id uint) error             { return config.DB.Delete(&StockLocation{}, id).Error }

func CreateStockPicking(data *StockPicking) error { return config.DB.Create(data).Error }
func GetAllStockPicking() ([]StockPicking, error) {
	var list []StockPicking
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockPickingByID(id uint) (*StockPicking, error) {
	var data StockPicking
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockPicking(data *StockPicking) error { return config.DB.Save(data).Error }
func DeleteStockPicking(id uint) error            { return config.DB.Delete(&StockPicking{}, id).Error }

func CreateStockLot(data *StockLot) error { return config.DB.Create(data).Error }
func GetAllStockLot() ([]StockLot, error) {
	var list []StockLot
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockLotByID(id uint) (*StockLot, error) {
	var data StockLot
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockLot(data *StockLot) error { return config.DB.Save(data).Error }
func DeleteStockLot(id uint) error        { return config.DB.Delete(&StockLot{}, id).Error }

func CreateStockQuant(data *StockQuant) error { return config.DB.Create(data).Error }
func GetAllStockQuant() ([]StockQuant, error) {
	var list []StockQuant
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockQuantByID(id uint) (*StockQuant, error) {
	var data StockQuant
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockQuant(data *StockQuant) error { return config.DB.Save(data).Error }
func DeleteStockQuant(id uint) error          { return config.DB.Delete(&StockQuant{}, id).Error }

func CreateStockPutawayRule(data *StockPutawayRule) error { return config.DB.Create(data).Error }
func GetAllStockPutawayRule() ([]StockPutawayRule, error) {
	var list []StockPutawayRule
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockPutawayRuleByID(id uint) (*StockPutawayRule, error) {
	var data StockPutawayRule
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockPutawayRule(data *StockPutawayRule) error { return config.DB.Save(data).Error }
func DeleteStockPutawayRule(id uint) error                { return config.DB.Delete(&StockPutawayRule{}, id).Error }

func CreateStockValuationLayer(data *StockValuationLayer) error { return config.DB.Create(data).Error }
func GetAllStockValuationLayer() ([]StockValuationLayer, error) {
	var list []StockValuationLayer
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetStockValuationLayerByID(id uint) (*StockValuationLayer, error) {
	var data StockValuationLayer
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateStockValuationLayer(data *StockValuationLayer) error { return config.DB.Save(data).Error }
func DeleteStockValuationLayer(id uint) error {
	return config.DB.Delete(&StockValuationLayer{}, id).Error
}
