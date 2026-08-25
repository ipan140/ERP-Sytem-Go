package dashboards

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateDashboard(data *Dashboard) error {
	return config.DB.Create(data).Error
}

func GetAllDashboard() ([]Dashboard, error) {
	var list []Dashboard
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetDashboardByID(id uint) (*Dashboard, error) {
	var data Dashboard
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateDashboard(data *Dashboard) error {
	return config.DB.Save(data).Error
}

func DeleteDashboard(id uint) error {
	return config.DB.Delete(&Dashboard{}, id).Error
}
