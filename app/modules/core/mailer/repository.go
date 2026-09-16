package mailer

import (
	"ERP-System/config"
	"gorm.io/gorm/clause"
)

func CreateEmailLog(data *EmailLog) error {
	return config.DB.Create(data).Error
}

func GetAllEmailLog() ([]EmailLog, error) {
	var list []EmailLog
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedEmailLog(offset, limit int, search string) ([]EmailLog, int64, error) {
	var list []EmailLog
	var total int64
	db := config.DB.Model(&EmailLog{})
	if search != "" {
		db = db.Where("recipient ILIKE ? OR subject ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Preload(clause.Associations).Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetEmailLogByID(id uint) (*EmailLog, error) {
	var data EmailLog
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateEmailLog(data *EmailLog) error {
	return config.DB.Save(data).Error
}

func DeleteEmailLog(id uint) error {
	return config.DB.Delete(&EmailLog{}, id).Error
}
