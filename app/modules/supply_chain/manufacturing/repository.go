package manufacturing

import (
	"ERP-System/config"
)

func CreateManufacturingOrder(data *ManufacturingOrder) error {
	return config.DB.Create(data).Error
}

func GetAllManufacturingOrder() ([]ManufacturingOrder, error) {
	var list []ManufacturingOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetManufacturingOrderByID(id uint) (*ManufacturingOrder, error) {
	var data ManufacturingOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateManufacturingOrder(data *ManufacturingOrder) error {
	return config.DB.Save(data).Error
}

func DeleteManufacturingOrder(id uint) error {
	return config.DB.Delete(&ManufacturingOrder{}, id).Error
}
