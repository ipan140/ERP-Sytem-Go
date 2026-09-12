package inventory

import (
	"ERP-System/app/modules/finance/accounting"
	"ERP-System/config"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type LandedCostInput struct {
	ProductID   uint    `json:"product_id"`
	SplitMethod string  `json:"split_method"`
	Amount      float64 `json:"amount"`
}

type CreateLandedCostRequest struct {
	PickingID uint              `json:"picking_id"`
	Costs     []LandedCostInput `json:"costs"`
}

func ProcessLandedCost(req CreateLandedCostRequest) error {
	return config.DB.Transaction(func(tx *gorm.DB) error {
		var picking StockPicking
		if err := tx.First(&picking, req.PickingID).Error; err != nil {
			return err
		}

		lcName := fmt.Sprintf("LC/%s/%%04d", time.Now().Format("20060102"), picking.ID)
		lc := StockLandedCost{
			Name:      lcName,
			Date:      time.Now(),
			PickingID: picking.ID,
		}
		if err := tx.Create(&lc).Error; err != nil {
			return err
		}

		var totalLC float64
		for _, c := range req.Costs {
			line := StockLandedCostLine{
				CostID:      lc.ID,
				ProductID:   c.ProductID,
				SplitMethod: c.SplitMethod,
				PriceUnit:   c.Amount,
			}
			tx.Create(&line)
			totalLC += c.Amount
		}

		var moves []StockMove
		tx.Where("picking_id = ? AND state = 'done'", picking.ID).Find(&moves)

		if len(moves) == 0 {
			return nil
		}

		splitAmount := totalLC / float64(len(moves))

		var inventoryAccount accounting.Account
		tx.Where("name = 'Persediaan Barang'").First(&inventoryAccount)
		if inventoryAccount.ID == 0 {
			inventoryAccount = accounting.Account{Code: "1-1401", Name: "Persediaan Barang", Type: "asset"}
			tx.Create(&inventoryAccount)
		}

		var lcAccount accounting.Account
		tx.Where("name = 'Beban Landed Cost'").First(&lcAccount)
		if lcAccount.ID == 0 {
			lcAccount = accounting.Account{Code: "6-1002", Name: "Beban Landed Cost", Type: "expense"}
			tx.Create(&lcAccount)
		}

		var journal accounting.Journal
		tx.Where("code = 'STJ'").First(&journal)
		if journal.ID == 0 {
			journal = accounting.Journal{Code: "STJ", Name: "Stock Journal", Type: "general"}
			tx.Create(&journal)
		}

		for _, m := range moves {
			valLayer := StockValuationLayer{
				ProductID:   m.ProductID,
				Quantity:    0,
				UnitCost:    0,
				Value:       splitAmount,
				Description: fmt.Sprintf("Landed Cost: %s untuk %s", lcName, picking.Name),
				CreatedAt:   time.Now(),
			}
			tx.Create(&valLayer)

			entry := accounting.JournalEntry{
				Name:      fmt.Sprintf("STJ/LC/%s/%d", time.Now().Format("200601"), m.ID),
				JournalID: journal.ID,
				Date:      time.Now(),
				State:     "posted",
				CreatedAt: time.Now(),
			}
			tx.Create(&entry)

			tx.Create(&accounting.JournalItem{
				EntryID:   entry.ID,
				AccountID: inventoryAccount.ID,
				Name:      valLayer.Description,
				Debit:     splitAmount,
				Credit:    0,
			})

			tx.Create(&accounting.JournalItem{
				EntryID:   entry.ID,
				AccountID: lcAccount.ID,
				Name:      valLayer.Description,
				Debit:     0,
				Credit:    splitAmount,
			})
		}

		return nil
	})
}

// FASE 15: Auto-Replenishment & Reordering Rules
func CreateStockWarehouseOrderpoint(req CreateOrderpointRequest) (*StockWarehouseOrderpoint, error) {
	opName := fmt.Sprintf("OP/%s/%04d", time.Now().Format("2006"), time.Now().Unix()%10000)
	op := StockWarehouseOrderpoint{
		Name:          opName,
		ProductID:     req.ProductID,
		WarehouseID:   req.WarehouseID,
		ProductMinQty: req.ProductMinQty,
		ProductMaxQty: req.ProductMaxQty,
		QtyMultiple:   req.QtyMultiple,
		Active:        true,
		CreatedAt:     time.Now(),
	}
	if op.QtyMultiple <= 0 {
		op.QtyMultiple = 1
	}
	if err := config.DB.Create(&op).Error; err != nil {
		return nil, err
	}
	config.DB.Preload("Product.ProductTemplate").Preload("Warehouse").First(&op, op.ID)
	return &op, nil
}

func GetAllStockWarehouseOrderpoints() ([]StockWarehouseOrderpoint, error) {
	var list []StockWarehouseOrderpoint
	err := config.DB.Preload("Product.ProductTemplate").Preload("Warehouse").Find(&list).Error
	return list, err
}

func DeleteStockWarehouseOrderpoint(id uint) error {
	return config.DB.Delete(&StockWarehouseOrderpoint{}, id).Error
}

type AutoReplenishResult struct {
	OrderpointsChecked int      `json:"orderpoints_checked"`
	DraftPOsCreated    int      `json:"draft_pos_created"`
	GeneratedPONames   []string `json:"generated_po_names"`
}

func RunAutoReplenishment() (*AutoReplenishResult, error) {
	var orderpoints []StockWarehouseOrderpoint
	if err := config.DB.Where("active = true").Find(&orderpoints).Error; err != nil {
		return nil, err
	}

	result := &AutoReplenishResult{
		OrderpointsChecked: len(orderpoints),
		GeneratedPONames:   make([]string, 0),
	}

	for _, op := range orderpoints {
		var prod Product
		if err := config.DB.First(&prod, op.ProductID).Error; err != nil {
			continue
		}

		// Check if stock_qty is less than min_qty
		if prod.StockQty <= op.ProductMinQty {
			qtyNeeded := op.ProductMaxQty - prod.StockQty
			if qtyNeeded <= 0 {
				continue
			}

			// Find default vendor from product_supplier_infos
			var supplierInfo struct {
				PartnerID uint
				Price     float64
			}
			_ = config.DB.Table("supply_chain.product_supplier_infos").
				Select("partner_id, price").
				Where("product_id = ?", op.ProductID).
				Order("price asc").
				Limit(1).
				Scan(&supplierInfo).Error

			partnerID := supplierInfo.PartnerID
			if partnerID == 0 {
				// Fallback to first vendor partner if none specified
				config.DB.Table("core.partners").Select("id").Where("is_vendor = true").Limit(1).Scan(&partnerID)
			}
			if partnerID == 0 {
				continue
			}

			priceUnit := supplierInfo.Price
			if priceUnit <= 0 {
				var pt ProductTemplate
				config.DB.First(&pt, prod.ProductTemplateID)
				priceUnit = pt.StandardPrice
			}

			// Generate Draft Purchase Order
			poName := fmt.Sprintf("PO/AUTO/%s/%04d", time.Now().Format("2006"), time.Now().Unix()%10000)
			po := struct {
				Name          string    `gorm:"type:varchar(100);not null"`
				PartnerID     uint      `json:"partner_id"`
				State         string    `json:"state"`
				AmountUntaxed float64   `json:"amount_untaxed"`
				AmountTotal   float64   `json:"amount_total"`
				DateOrder     time.Time `json:"date_order"`
				Notes         string    `json:"notes"`
				CreatedAt     time.Time `json:"created_at"`
			}{
				Name:          poName,
				PartnerID:     partnerID,
				State:         "draft",
				AmountUntaxed: qtyNeeded * priceUnit,
				AmountTotal:   qtyNeeded * priceUnit,
				DateOrder:     time.Now(),
				Notes:         fmt.Sprintf("Auto-replenishment dari Aturan Reordering %s (Stok Saat Ini: %.2f, Min: %.2f)", op.Name, prod.StockQty, op.ProductMinQty),
				CreatedAt:     time.Now(),
			}

			if err := config.DB.Table("supply_chain.purchase_orders").Create(&po).Error; err == nil {
				// Create PO Line
				var poID uint
				config.DB.Table("supply_chain.purchase_orders").Select("id").Where("name = ?", poName).Scan(&poID)
				if poID > 0 {
					line := struct {
						OrderID       uint    `json:"order_id"`
						ProductID     uint    `json:"product_id"`
						Name          string  `json:"name"`
						Quantity      float64 `json:"quantity"`
						PriceUnit     float64 `json:"price_unit"`
						PriceSubtotal float64 `json:"price_subtotal"`
					}{
						OrderID:       poID,
						ProductID:     prod.ID,
						Name:          fmt.Sprintf("Pengadaan Otomatis: %s", prod.DefaultCode),
						Quantity:      qtyNeeded,
						PriceUnit:     priceUnit,
						PriceSubtotal: qtyNeeded * priceUnit,
					}
					config.DB.Table("supply_chain.purchase_order_lines").Create(&line)
					result.DraftPOsCreated++
					result.GeneratedPONames = append(result.GeneratedPONames, poName)
				}
			}
		}
	}

	return result, nil
}
