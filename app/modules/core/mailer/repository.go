package mailer

import (
	"ERP-System/config"
)

func CreateEmailLog(data *EmailLog) error {
	return config.DB.Create(data).Error
}

func GetAllEmailLog() ([]EmailLog, error) {
	var list []EmailLog
	err := config.DB.Find(&list).Error
	return list, err
}

func GetEmailLogByID(id uint) (*EmailLog, error) {
	var data EmailLog
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateEmailLog(data *EmailLog) error {
	return config.DB.Save(data).Error
}

func DeleteEmailLog(id uint) error {
	return config.DB.Delete(&EmailLog{}, id).Error
}
