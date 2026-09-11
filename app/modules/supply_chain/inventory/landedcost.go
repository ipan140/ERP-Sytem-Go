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
