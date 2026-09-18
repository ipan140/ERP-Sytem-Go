package voip

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateCallRecord(data *CallRecord) error {
	return config.DB.Create(data).Error
}

func GetAllCallRecord() ([]CallRecord, error) {
	var list []CallRecord
	err := config.DB.Order("id desc").Find(&list).Error
	return list, err
}

func GetAllVoipExtensions() ([]VoipExtension, error) {
	var list []VoipExtension
	err := config.DB.Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultVoipExtensions()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultVoipExtensions() error {
	defaults := []VoipExtension{
		{Ext: "101", User: "Siti Aminah", Dept: "Human Resources & General Affairs", Device: "Yealink T46U SIP", IP: "192.168.10.101", Status: "AVAILABLE"},
		{Ext: "102", User: "Budi Santoso", Dept: "Finance & Accounting", Device: "Grandstream GXP2170", IP: "192.168.10.102", Status: "AVAILABLE"},
		{Ext: "103", User: "Joko Prabowo", Dept: "Supply Chain & Warehouse Cikarang", Device: "Fanvil X4U Industrial", IP: "192.168.20.103", Status: "ON CALL"},
		{Ext: "104", User: "Rina Kusuma", Dept: "Sales & Marketing Enterprise", Device: "WebRTC Softphone Desktop", IP: "192.168.10.104", Status: "AVAILABLE"},
		{Ext: "201", User: "Helpdesk Support Tier-1", Dept: "Customer Service & IT Support", Device: "Call Center Queue SIP", IP: "192.168.10.201", Status: "AVAILABLE"},
	}
	for _, ext := range defaults {
		var existing VoipExtension
		if err := config.DB.Where("ext = ?", ext.Ext).First(&existing).Error; err != nil {
			_ = config.DB.Create(&ext).Error
		}
	}
	return nil
}

func CreateVoipExtension(data *VoipExtension) error {
	return config.DB.Create(data).Error
}

func GetCallRecordByID(id uint) (*CallRecord, error) {
	var data CallRecord
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateCallRecord(data *CallRecord) error {
	return config.DB.Save(data).Error
}

func DeleteCallRecord(id uint) error {
	return config.DB.Delete(&CallRecord{}, id).Error
}
