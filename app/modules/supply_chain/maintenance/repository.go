package maintenance

import (
	"ERP-System/config"
)

func CreateMaintenanceRequest(data *MaintenanceRequest) error {
	return config.DB.Create(data).Error
}

func GetAllMaintenanceRequest() ([]MaintenanceRequest, error) {
	var list []MaintenanceRequest
	err := config.DB.Find(&list).Error
	return list, err
}

func GetMaintenanceRequestByID(id uint) (*MaintenanceRequest, error) {
	var data MaintenanceRequest
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateMaintenanceRequest(data *MaintenanceRequest) error {
	return config.DB.Save(data).Error
}

func DeleteMaintenanceRequest(id uint) error {
	return config.DB.Delete(&MaintenanceRequest{}, id).Error
}
