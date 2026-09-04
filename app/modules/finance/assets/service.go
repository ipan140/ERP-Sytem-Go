package assets

import (
	"ERP-System/config"
	"fmt"
	"math"
	"time"

	"gorm.io/gorm"
)

// Category Services
func GetAllCategoriesService() ([]AssetCategory, error) {
	return GetAllCategoriesRepo()
}

func CreateCategoryService(cat *AssetCategory) error {
	if cat.DefaultUsefulLife <= 0 {
		cat.DefaultUsefulLife = 48
	}
	return CreateCategoryRepo(cat)
}

func UpdateCategoryService(cat *AssetCategory) error {
	return UpdateCategoryRepo(cat)
}

func DeleteCategoryService(id uint) error {
	return DeleteCategoryRepo(id)
}

// Asset Services
func GetAllAssetsService() ([]FixedAsset, float64, float64, float64, error) {
	list, err := GetAllAssetsRepo()
	if err != nil {
		return nil, 0, 0, 0, err
	}

	var totalAcq, totalAccum, totalNBV float64
	for _, it := range list {
		totalAcq += it.AcquisitionCost
		totalAccum += it.AccumulatedDepr
		totalNBV += it.NetBookValue
	}
	return list, totalAcq, totalAccum, totalNBV, nil
}

func CreateAssetService(asset *FixedAsset) error {
	if asset.AcquisitionDate.IsZero() {
		asset.AcquisitionDate = time.Now()
	}
	if asset.UsefulLifeMonths <= 0 {
		asset.UsefulLifeMonths = 48 // Default 4 tahun
	}
	if asset.DepreciationMethod == "" {
		asset.DepreciationMethod = "Garis Lurus"
	}

	// Hitung depresiasi per bulan: (Harga - Residu) / Bulan
	depreciableAmount := asset.AcquisitionCost - asset.ResidualValue
	if depreciableAmount > 0 {
		asset.MonthlyDepreciation = math.Round(depreciableAmount / float64(asset.UsefulLifeMonths))
	}
	asset.NetBookValue = asset.AcquisitionCost - asset.AccumulatedDepr
	asset.Status = "Aktif"

	return CreateAssetRepo(asset)
}

func UpdateAssetService(asset *FixedAsset) error {
	depreciableAmount := asset.AcquisitionCost - asset.ResidualValue
	if asset.UsefulLifeMonths > 0 && depreciableAmount > 0 {
		asset.MonthlyDepreciation = math.Round(depreciableAmount / float64(asset.UsefulLifeMonths))
	}
	asset.NetBookValue = asset.AcquisitionCost - asset.AccumulatedDepr
	return UpdateAssetRepo(asset)
}

func DeleteAssetService(id uint) error {
	return DeleteAssetRepo(id)
}

func ExecuteMonthlyDepreciationService() (float64, error) {
	list, err := GetActiveAssetsForDepreciationRepo()
	if err != nil {
		return 0, err
	}

	var totalDepreciated float64
	for _, it := range list {
		if it.NetBookValue > it.ResidualValue {
			deprNow := it.MonthlyDepreciation
			if it.NetBookValue-deprNow < it.ResidualValue {
				deprNow = it.NetBookValue - it.ResidualValue
			}
			it.AccumulatedDepr += deprNow
			it.NetBookValue -= deprNow
			_ = UpdateAssetRepo(&it)
			totalDepreciated += deprNow
		}
	}

	// Otomatis Posting ke General Ledger (GL) jika ada nilai depresiasi
	if totalDepreciated > 0 {
		now := time.Now()
		entryRef := fmt.Sprintf("DEPR/%d/%02d", now.Year(), now.Month())

		// Cari akun GL Beban Depresiasi (6-3000) dan Akumulasi Penyusutan (1-2001)
		type accountMini struct {
			ID      uint
			Code    string
			Balance float64
		}
		var deprExpenseAcc, accumDeprAcc accountMini
		config.DB.Table("accounts").Where("code = ?", "6-3000").First(&deprExpenseAcc)
		config.DB.Table("accounts").Where("code = ?", "1-2001").First(&accumDeprAcc)

		type journalEntryMini struct {
			ID        uint      `gorm:"primaryKey"`
			Name      string
			Date      time.Time
			State     string
			CreatedAt time.Time
		}
		type journalItemMini struct {
			EntryID   uint
			AccountID uint
			Name      string
			Debit     float64
			Credit    float64
		}

		jEntry := journalEntryMini{
			Name:      entryRef,
			Date:      now,
			State:     "posted",
			CreatedAt: now,
		}
		if err := config.DB.Table("journal_entries").Create(&jEntry).Error; err == nil {
			if deprExpenseAcc.ID > 0 {
				config.DB.Table("journal_items").Create(&journalItemMini{
					EntryID:   jEntry.ID,
					AccountID: deprExpenseAcc.ID,
					Name:      fmt.Sprintf("Beban Penyusutan Aset Tetap Periode %s %d", now.Month().String(), now.Year()),
					Debit:     totalDepreciated,
					Credit:    0,
				})
				config.DB.Table("accounts").Where("id = ?", deprExpenseAcc.ID).
					UpdateColumn("balance", gorm.Expr("balance + ?", totalDepreciated))
			}
			if accumDeprAcc.ID > 0 {
				config.DB.Table("journal_items").Create(&journalItemMini{
					EntryID:   jEntry.ID,
					AccountID: accumDeprAcc.ID,
					Name:      fmt.Sprintf("Akumulasi Penyusutan Aset Tetap Periode %s %d", now.Month().String(), now.Year()),
					Debit:     0,
					Credit:    totalDepreciated,
				})
				// Akun kontra-aset bersaldo negatif atau bertambah kredit
				config.DB.Table("accounts").Where("id = ?", accumDeprAcc.ID).
					UpdateColumn("balance", gorm.Expr("balance - ?", totalDepreciated))
			}
		}
	}

	return totalDepreciated, nil
}
