package base

import (
	"ERP-System/config"
)

func CreateCurrency(data *Currency) error {
	return config.DB.Create(data).Error
}

func GetAllCurrency() ([]Currency, error) {
	var list []Currency
	err := config.DB.Find(&list).Error
	return list, err
}

func GetCurrencyByID(id uint) (*Currency, error) {
	var data Currency
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateCurrency(data *Currency) error {
	return config.DB.Save(data).Error
}

func DeleteCurrency(id uint) error {
	return config.DB.Delete(&Currency{}, id).Error
}
