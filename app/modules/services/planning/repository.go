package planning

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateShift(data *Shift) error {
	return config.DB.Create(data).Error
}

func GetAllShift() ([]Shift, error) {
	var list []Shift
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetShiftByID(id uint) (*Shift, error) {
	var data Shift
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateShift(data *Shift) error {
	return config.DB.Save(data).Error
}

func DeleteShift(id uint) error {
	return config.DB.Delete(&Shift{}, id).Error
}
