package crm

import (
	"ERP-System/config"
)

func CreateLead(data *Lead) error {
	return config.DB.Create(data).Error
}

func GetAllLead() ([]Lead, error) {
	var list []Lead
	err := config.DB.Find(&list).Error
	return list, err
}

func GetLeadByID(id uint) (*Lead, error) {
	var data Lead
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateLead(data *Lead) error {
	return config.DB.Save(data).Error
}

func DeleteLead(id uint) error {
	return config.DB.Delete(&Lead{}, id).Error
}
