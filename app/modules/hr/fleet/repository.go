package fleet

import (
	"ERP-System/config"
)

func CreateFleetVehicle(data *FleetVehicle) error {
	return config.DB.Create(data).Error
}

func GetAllFleetVehicle() ([]FleetVehicle, error) {
	var list []FleetVehicle
	err := config.DB.Find(&list).Error
	return list, err
}

func GetFleetVehicleByID(id uint) (*FleetVehicle, error) {
	var data FleetVehicle
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateFleetVehicle(data *FleetVehicle) error {
	return config.DB.Save(data).Error
}

func DeleteFleetVehicle(id uint) error {
	return config.DB.Delete(&FleetVehicle{}, id).Error
}
