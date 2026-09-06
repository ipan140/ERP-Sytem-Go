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

// WhatsApp Template Repository
func CreateWaTemplate(data *WaTemplate) error {
	return config.DB.Create(data).Error
}

func GetAllWaTemplates() ([]WaTemplate, error) {
	var list []WaTemplate
	err := config.DB.Order("id desc").Find(&list).Error
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

// WhatsApp Config Repository
func GetWaConfig() (*WaConfig, error) {
	// Prioritaskan membaca dari .env file
	envWabaID := config.GetEnv("META_WABA_ID", "")
	envPhoneID := config.GetEnv("META_PHONE_NUMBER_ID", "")
	envToken := config.GetEnv("META_ACCESS_TOKEN", "")
	envVersion := config.GetEnv("META_API_VERSION", "v20.0")
	envWebhook := config.GetEnv("META_WEBHOOK_VERIFY_TOKEN", "")

	if envWabaID != "" || envPhoneID != "" || envToken != "" {
		return &WaConfig{
			PhoneNumberID: envPhoneID,
			WabaID:        envWabaID,
			AccessToken:   envToken,
			ApiVersion:    envVersion,
			WebhookToken:  envWebhook,
			IsActive:      true,
		}, nil
	}

	var cfg WaConfig
	err := config.DB.First(&cfg).Error
	if err != nil {
		return &WaConfig{
			PhoneNumberID: "",
			WabaID:        "",
			AccessToken:   "",
			ApiVersion:    "v20.0",
			WebhookToken:  "",
			IsActive:      false,
		}, nil
	}
	return &cfg, nil
}

func SaveWaConfig(cfg *WaConfig) error {
	var existing WaConfig
	if err := config.DB.First(&existing).Error; err != nil {
		return config.DB.Create(cfg).Error
	}
	cfg.ID = existing.ID
	return config.DB.Save(cfg).Error
}
