package plm

import (
	"ERP-System/config"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPaginatedEcos(page, limit int, search, state string) ([]PlmEco, int64, error) {
	var list []PlmEco
	var total int64

	db := config.DB.Model(&PlmEco{})

	if search != "" {
		s := "%" + strings.ToLower(search) + "%"
		db = db.Where("LOWER(name) LIKE ? OR LOWER(code) LIKE ?", s, s)
	}

	if state != "" && state != "all" {
		db = db.Where("state = ?", state)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Preload(clause.Associations).
		Preload("Product.ProductTemplate").
		Order("id desc").
		Offset(offset).
		Limit(limit).
		Find(&list).Error

	return list, total, err
}

func GetPlmSummary() (*PlmSummary, error) {
	var summary PlmSummary

	config.DB.Model(&PlmEco{}).Count(&summary.TotalEco)
	config.DB.Model(&PlmEco{}).Where("state = ?", "draft").Count(&summary.DraftCount)
	config.DB.Model(&PlmEco{}).Where("state = ?", "progress").Count(&summary.ProgressCount)
	config.DB.Model(&PlmEco{}).Where("state = ?", "approved").Count(&summary.ApprovedCount)
	config.DB.Model(&PlmEco{}).Where("state = ?", "done").Count(&summary.DoneCount)

	return &summary, nil
}

func CreateEcoWithSequence(req *CreateEcoRequest) (*PlmEco, error) {
	var eco PlmEco
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		year := time.Now().Format("2006")
		prefix := fmt.Sprintf("ECO/%s/", year)

		var last PlmEco
		var nextNum int64 = 1
		if err := tx.Unscoped().Where("code LIKE ?", prefix+"%").Order("code desc").First(&last).Error; err == nil {
			var lastNum int64
			fmt.Sscanf(last.Code, prefix+"%d", &lastNum)
			nextNum = lastNum + 1
		} else {
			var count int64
			tx.Model(&PlmEco{}).Count(&count)
			nextNum = count + 1
		}

		code := fmt.Sprintf("ECO/%s/%04d", year, nextNum)

		eco = PlmEco{
			Code:      code,
			Name:      req.Name,
			TypeID:    req.TypeID,
			ProductID: req.ProductID,
			OldBomID:  req.OldBomID,
			NewBomID:  req.NewBomID,
			Reason:    req.Reason,
			State:     "draft",
		}

		if req.EffectiveDate != nil && *req.EffectiveDate != "" {
			parsedDate, err := time.Parse(time.RFC3339, *req.EffectiveDate)
			if err != nil {
				parsedDate, err = time.Parse("2006-01-02", *req.EffectiveDate)
			}
			if err == nil {
				eco.EffectiveDate = &parsedDate
			}
		}

		return tx.Create(&eco).Error
	})

	if err != nil {
		return nil, err
	}

	config.DB.Preload(clause.Associations).Preload("Product.ProductTemplate").First(&eco, eco.ID)
	return &eco, nil
}

func UpdateEcoState(id uint, state string, approverID *uint) (*PlmEco, error) {
	var eco PlmEco
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&eco, id).Error; err != nil {
			return err
		}

		eco.State = state
		if state == "approved" {
			now := time.Now()
			eco.ApprovedAt = &now
			eco.ApproverID = approverID
		} else if state == "done" {
			// FASE 10: Apply to Live BOM
			// Archive Old BOM
			if eco.OldBomID != nil {
				if err := tx.Exec("UPDATE supply_chain.mrp_boms SET active = false WHERE id = ?", *eco.OldBomID).Error; err != nil {
					return err
				}
			}
			// Set New BOM as Active and bump version
			if eco.NewBomID != nil {
				var oldVersion int
				tx.Raw("SELECT version FROM supply_chain.mrp_boms WHERE id = ?", *eco.OldBomID).Scan(&oldVersion)
				if oldVersion == 0 {
					oldVersion = 1
				}
				if err := tx.Exec("UPDATE supply_chain.mrp_boms SET active = true, version = ? WHERE id = ?", oldVersion+1, *eco.NewBomID).Error; err != nil {
					return err
				}
			}
		}

		return tx.Save(&eco).Error
	})

	if err != nil {
		return nil, err
	}

	config.DB.Preload(clause.Associations).Preload("Product.ProductTemplate").First(&eco, eco.ID)
	return &eco, nil
}

func GetAllEcoTypes() ([]PlmEcoType, error) {
	var types []PlmEcoType
	err := config.DB.Find(&types).Error
	// Seed default types if empty
	if len(types) == 0 {
		defaultTypes := []PlmEcoType{
			{Name: "Perubahan Resep (BoM Change)"},
			{Name: "Substitusi Bahan Baku (Material Substitution)"},
			{Name: "Penyempurnaan Proses (Routing Change)"},
			{Name: "Peningkatan Mutu (Quality Improvement)"},
		}
		for _, t := range defaultTypes {
			config.DB.Create(&t)
		}
		config.DB.Find(&types)
	}
	return types, err
}

func GetEcoByID(id uint) (*PlmEco, error) {
	var eco PlmEco
	err := config.DB.Preload(clause.Associations).Preload("Product.ProductTemplate").First(&eco, id).Error
	return &eco, err
}

func DeleteEco(id uint) error {
	return config.DB.Delete(&PlmEco{}, id).Error
}
