package sms_marketing

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSmsCampaign(data *SmsCampaign) error {
	return config.DB.Create(data).Error
}

func GetAllSmsCampaign() ([]SmsCampaign, error) {
	var list []SmsCampaign
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetSmsCampaignByID(id uint) (*SmsCampaign, error) {
	var data SmsCampaign
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSmsCampaign(data *SmsCampaign) error {
	return config.DB.Save(data).Error
}

func DeleteSmsCampaign(id uint) error {
	return config.DB.Delete(&SmsCampaign{}, id).Error
}
