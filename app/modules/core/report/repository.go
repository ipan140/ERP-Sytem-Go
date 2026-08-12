package report

import (
	"ERP-System/config"
)

func CreateReport(data *Report) error {
	return config.DB.Create(data).Error
}

func GetAllReport() ([]Report, error) {
	var list []Report
	err := config.DB.Find(&list).Error
	return list, err
}

func GetReportByID(id uint) (*Report, error) {
	var data Report
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateReport(data *Report) error {
	return config.DB.Save(data).Error
}

func DeleteReport(id uint) error {
	return config.DB.Delete(&Report{}, id).Error
}
