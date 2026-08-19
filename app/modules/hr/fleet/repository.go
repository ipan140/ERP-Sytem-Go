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

func CreateVehicleLogContract(data *VehicleLogContract) error { return config.DB.Create(data).Error }
func GetAllVehicleLogContract() ([]VehicleLogContract, error) { var list []VehicleLogContract; err := config.DB.Find(&list).Error; return list, err }
func GetVehicleLogContractByID(id uint) (*VehicleLogContract, error) { var data VehicleLogContract; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateVehicleLogContract(data *VehicleLogContract) error { return config.DB.Save(data).Error }
func DeleteVehicleLogContract(id uint) error { return config.DB.Delete(&VehicleLogContract{}, id).Error }

func CreateVehicleLogFuel(data *VehicleLogFuel) error { return config.DB.Create(data).Error }
func GetAllVehicleLogFuel() ([]VehicleLogFuel, error) { var list []VehicleLogFuel; err := config.DB.Find(&list).Error; return list, err }
func GetVehicleLogFuelByID(id uint) (*VehicleLogFuel, error) { var data VehicleLogFuel; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateVehicleLogFuel(data *VehicleLogFuel) error { return config.DB.Save(data).Error }
func DeleteVehicleLogFuel(id uint) error { return config.DB.Delete(&VehicleLogFuel{}, id).Error }

func CreateVehicleLogServices(data *VehicleLogServices) error { return config.DB.Create(data).Error }
func GetAllVehicleLogServices() ([]VehicleLogServices, error) { var list []VehicleLogServices; err := config.DB.Find(&list).Error; return list, err }
func GetVehicleLogServicesByID(id uint) (*VehicleLogServices, error) { var data VehicleLogServices; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateVehicleLogServices(data *VehicleLogServices) error { return config.DB.Save(data).Error }
func DeleteVehicleLogServices(id uint) error { return config.DB.Delete(&VehicleLogServices{}, id).Error }
