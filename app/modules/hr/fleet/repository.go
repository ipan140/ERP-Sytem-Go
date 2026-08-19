package fleet

import (
	"ERP-System/config"
)

func CreateVehicle(data *Vehicle) error {
	return config.DB.Create(data).Error
}

func GetAllVehicle() ([]Vehicle, error) {
	var list []Vehicle
	err := config.DB.Find(&list).Error
	return list, err
}

func GetVehicleByID(id uint) (*Vehicle, error) {
	var data Vehicle
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateVehicle(data *Vehicle) error {
	return config.DB.Save(data).Error
}

func DeleteVehicle(id uint) error {
	return config.DB.Delete(&Vehicle{}, id).Error
}
