package marketing_automation

import (
	"ERP-System/config"
)

func CreateAutomationCampaign(data *AutomationCampaign) error {
	return config.DB.Create(data).Error
}

func GetAllAutomationCampaign() ([]AutomationCampaign, error) {
	var list []AutomationCampaign
	err := config.DB.Find(&list).Error
	return list, err
}

func GetAutomationCampaignByID(id uint) (*AutomationCampaign, error) {
	var data AutomationCampaign
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateAutomationCampaign(data *AutomationCampaign) error {
	return config.DB.Save(data).Error
}

func DeleteAutomationCampaign(id uint) error {
	return config.DB.Delete(&AutomationCampaign{}, id).Error
}
