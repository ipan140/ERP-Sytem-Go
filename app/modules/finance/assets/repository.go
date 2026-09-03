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
