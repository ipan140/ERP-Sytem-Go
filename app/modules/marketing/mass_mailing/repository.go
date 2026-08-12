package mass_mailing

import (
	"ERP-System/config"
)

func CreateMailingCampaign(data *MailingCampaign) error {
	return config.DB.Create(data).Error
}

func GetAllMailingCampaign() ([]MailingCampaign, error) {
	var list []MailingCampaign
	err := config.DB.Find(&list).Error
	return list, err
}

func GetMailingCampaignByID(id uint) (*MailingCampaign, error) {
	var data MailingCampaign
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateMailingCampaign(data *MailingCampaign) error {
	return config.DB.Save(data).Error
}

func DeleteMailingCampaign(id uint) error {
	return config.DB.Delete(&MailingCampaign{}, id).Error
}
