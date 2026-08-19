package barcode

import (
	"ERP-System/config"
)

func CreateBarcodeConfig(data *BarcodeNomenclature) error {
	return config.DB.Create(data).Error
}

func GetAllBarcodeConfig() ([]BarcodeNomenclature, error) {
	var list []BarcodeNomenclature
	err := config.DB.Find(&list).Error
	return list, err
}

func GetBarcodeConfigByID(id uint) (*BarcodeNomenclature, error) {
	var data BarcodeNomenclature
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateBarcodeConfig(data *BarcodeNomenclature) error {
	return config.DB.Save(data).Error
}

func DeleteBarcodeConfig(id uint) error {
	return config.DB.Delete(&BarcodeNomenclature{}, id).Error
}
