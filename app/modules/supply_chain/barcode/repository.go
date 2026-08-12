package barcode

import (
	"ERP-System/config"
)

func CreateBarcodeConfig(data *BarcodeConfig) error {
	return config.DB.Create(data).Error
}

func GetAllBarcodeConfig() ([]BarcodeConfig, error) {
	var list []BarcodeConfig
	err := config.DB.Find(&list).Error
	return list, err
}

func GetBarcodeConfigByID(id uint) (*BarcodeConfig, error) {
	var data BarcodeConfig
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateBarcodeConfig(data *BarcodeConfig) error {
	return config.DB.Save(data).Error
}

func DeleteBarcodeConfig(id uint) error {
	return config.DB.Delete(&BarcodeConfig{}, id).Error
}
