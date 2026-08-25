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
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
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
