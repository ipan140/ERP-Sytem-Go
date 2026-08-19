package plm

import (
	"ERP-System/config"
)

func CreateBom(data *PlmEco) error {
	return config.DB.Create(data).Error
}

func GetAllBom() ([]PlmEco, error) {
	var list []PlmEco
	err := config.DB.Find(&list).Error
	return list, err
}

func GetBomByID(id uint) (*PlmEco, error) {
	var data PlmEco
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateBom(data *PlmEco) error {
	return config.DB.Save(data).Error
}

func DeleteBom(id uint) error {
	return config.DB.Delete(&PlmEco{}, id).Error
}
