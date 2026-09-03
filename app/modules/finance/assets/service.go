package assets

import (
	"math"
	"time"
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
			if it.NetBookValue - deprNow < it.ResidualValue {
				deprNow = it.NetBookValue - it.ResidualValue
			}
			it.AccumulatedDepr += deprNow
			it.NetBookValue -= deprNow
			_ = UpdateAssetRepo(&it)
			totalDepreciated += deprNow
		}
	}
	return totalDepreciated, nil
}
