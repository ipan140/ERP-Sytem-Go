package repairs

import (
	"ERP-System/config"
)

func CreateRepairOrder(data *RepairOrder) error {
	return config.DB.Create(data).Error
}

func GetAllRepairOrder() ([]RepairOrder, error) {
	var list []RepairOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetRepairOrderByID(id uint) (*RepairOrder, error) {
	var data RepairOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateRepairOrder(data *RepairOrder) error {
	return config.DB.Save(data).Error
}

func DeleteRepairOrder(id uint) error {
	return config.DB.Delete(&RepairOrder{}, id).Error
}
