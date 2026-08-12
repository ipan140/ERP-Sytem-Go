package plm

import (
	"ERP-System/config"
)

func CreateBom(data *Bom) error {
	return config.DB.Create(data).Error
}

func GetAllBom() ([]Bom, error) {
	var list []Bom
	err := config.DB.Find(&list).Error
	return list, err
}

func GetBomByID(id uint) (*Bom, error) {
	var data Bom
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateBom(data *Bom) error {
	return config.DB.Save(data).Error
}

func DeleteBom(id uint) error {
	return config.DB.Delete(&Bom{}, id).Error
}
