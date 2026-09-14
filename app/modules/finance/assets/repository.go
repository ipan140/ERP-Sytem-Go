package assets

import (
	"ERP-System/config"
)

func GetAllCategoriesRepo() ([]AssetCategory, error) {
	var list []AssetCategory
	err := config.DB.Order("id asc").Find(&list).Error
	return list, err
}

func GetCategoryByIDRepo(id uint) (*AssetCategory, error) {
	var item AssetCategory
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateCategoryRepo(cat *AssetCategory) error {
	return config.DB.Create(cat).Error
}

func UpdateCategoryRepo(cat *AssetCategory) error {
	return config.DB.Save(cat).Error
}

func DeleteCategoryRepo(id uint) error {
	return config.DB.Delete(&AssetCategory{}, id).Error
}

func GetAllAssetsRepo() ([]FixedAsset, error) {
	var list []FixedAsset
	err := config.DB.Order("id asc").Find(&list).Error
	return list, err
}

func GetPaginatedAssetsRepo(offset int, limit int, search string, category string) ([]FixedAsset, int64, float64, float64, float64, error) {
	var list []FixedAsset
	var total int64

	var summary struct {
		TotalAcq   float64
		TotalAccum float64
		TotalNBV   float64
	}
	config.DB.Model(&FixedAsset{}).Select("COALESCE(SUM(acquisition_cost), 0) as total_acq, COALESCE(SUM(accumulated_depr), 0) as total_accum, COALESCE(SUM(net_book_value), 0) as total_nbv").Scan(&summary)

	query := config.DB.Model(&FixedAsset{})
	if category != "" && category != "all" && category != "All" && category != "Semua" {
		query = query.Where("category ILIKE ?", "%"+category+"%")
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR code ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, 0, 0, 0, err
	}

	err := query.Order("id asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, summary.TotalAcq, summary.TotalAccum, summary.TotalNBV, err
}

func GetAssetByIDRepo(id uint) (*FixedAsset, error) {
	var item FixedAsset
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateAssetRepo(asset *FixedAsset) error {
	return config.DB.Create(asset).Error
}

func UpdateAssetRepo(asset *FixedAsset) error {
	return config.DB.Save(asset).Error
}

func DeleteAssetRepo(id uint) error {
	return config.DB.Delete(&FixedAsset{}, id).Error
}

func GetActiveAssetsForDepreciationRepo() ([]FixedAsset, error) {
	var list []FixedAsset
	err := config.DB.Where("status = ?", "Aktif").Find(&list).Error
	return list, err
}
