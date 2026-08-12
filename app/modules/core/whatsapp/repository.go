package whatsapp

import (
	"ERP-System/config"
)

func CreateWaTemplate(data *WaTemplate) error {
	return config.DB.Create(data).Error
}

func GetAllWaTemplate() ([]WaTemplate, error) {
	var list []WaTemplate
	err := config.DB.Find(&list).Error
	return list, err
}

func GetWaTemplateByID(id uint) (*WaTemplate, error) {
	var data WaTemplate
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateWaTemplate(data *WaTemplate) error {
	return config.DB.Save(data).Error
}

func DeleteWaTemplate(id uint) error {
	return config.DB.Delete(&WaTemplate{}, id).Error
}
