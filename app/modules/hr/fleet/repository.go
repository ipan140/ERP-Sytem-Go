package fleet

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateVehicle(data *Vehicle) error {
	return config.DB.Create(data).Error
}

func GetAllVehicle() ([]Vehicle, error) {
	var list []Vehicle
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedVehicles(offset, limit int, search string, employeeID string, state string) ([]Vehicle, int64, error) {
	var list []Vehicle
	var total int64

	query := config.DB.Model(&Vehicle{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.vehicles.employee_id = ?", employeeID)
	}

	if state != "" && state != "all" {
		query = query.Where("hrd.vehicles.state = ?", state)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.vehicles.employee_id").
			Where("hrd.vehicles.model_name ILIKE ? OR hrd.vehicles.license_plate ILIKE ? OR hrd.employees.name ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.vehicles.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetVehicleByID(id uint) (*Vehicle, error) {
	var data Vehicle
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateVehicle(data *Vehicle) error {
	return config.DB.Save(data).Error
}

func DeleteVehicle(id uint) error {
	return config.DB.Delete(&Vehicle{}, id).Error
}

func CreateVehicleLogContract(data *VehicleLogContract) error { return config.DB.Create(data).Error }
func GetAllVehicleLogContract() ([]VehicleLogContract, error) {
	var list []VehicleLogContract
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetVehicleLogContractByID(id uint) (*VehicleLogContract, error) {
	var data VehicleLogContract
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateVehicleLogContract(data *VehicleLogContract) error { return config.DB.Save(data).Error }
func DeleteVehicleLogContract(id uint) error {
	return config.DB.Delete(&VehicleLogContract{}, id).Error
}

func CreateVehicleLogFuel(data *VehicleLogFuel) error { return config.DB.Create(data).Error }
func GetAllVehicleLogFuel() ([]VehicleLogFuel, error) {
	var list []VehicleLogFuel
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetVehicleLogFuelByID(id uint) (*VehicleLogFuel, error) {
	var data VehicleLogFuel
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateVehicleLogFuel(data *VehicleLogFuel) error { return config.DB.Save(data).Error }
func DeleteVehicleLogFuel(id uint) error              { return config.DB.Delete(&VehicleLogFuel{}, id).Error }

func CreateVehicleLogServices(data *VehicleLogServices) error { return config.DB.Create(data).Error }
func GetAllVehicleLogServices() ([]VehicleLogServices, error) {
	var list []VehicleLogServices
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetVehicleLogServicesByID(id uint) (*VehicleLogServices, error) {
	var data VehicleLogServices
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateVehicleLogServices(data *VehicleLogServices) error { return config.DB.Save(data).Error }
func DeleteVehicleLogServices(id uint) error {
	return config.DB.Delete(&VehicleLogServices{}, id).Error
}
