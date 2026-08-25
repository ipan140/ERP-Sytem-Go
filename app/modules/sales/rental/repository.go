package rental

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateRentalOrder(data *RentalOrder) error {
	return config.DB.Create(data).Error
}

func GetAllRentalOrder() ([]RentalOrder, error) {
	var list []RentalOrder
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetRentalOrderByID(id uint) (*RentalOrder, error) {
	var data RentalOrder
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateRentalOrder(data *RentalOrder) error {
	return config.DB.Save(data).Error
}

func DeleteRentalOrder(id uint) error {
	return config.DB.Delete(&RentalOrder{}, id).Error
}

func CreateRentalOrderLine(data *RentalOrderLine) error { return config.DB.Create(data).Error }
func GetAllRentalOrderLine() ([]RentalOrderLine, error) {
	var list []RentalOrderLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetRentalOrderLineByID(id uint) (*RentalOrderLine, error) {
	var data RentalOrderLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateRentalOrderLine(data *RentalOrderLine) error { return config.DB.Save(data).Error }
func DeleteRentalOrderLine(id uint) error               { return config.DB.Delete(&RentalOrderLine{}, id).Error }
