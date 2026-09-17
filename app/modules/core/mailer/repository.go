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

func GetActiveSmtpConfig() (*SmtpConfig, error) {
	var cfg SmtpConfig
	err := config.DB.Where("is_active = ?", true).Order("id desc").First(&cfg).Error
	if err != nil {
		// Buat default jika belum ada di database
		cfg = SmtpConfig{
			Host:       "smtp.gmail.com",
			Port:       587,
			Username:   "no-reply@erp-enterprise.co.id",
			Password:   "",
			SenderName: "ERP Notification System",
			Encryption: "STARTTLS",
			IsActive:   true,
		}
		if createErr := config.DB.Create(&cfg).Error; createErr != nil {
			return nil, createErr
		}
		return &cfg, nil
	}
	return &cfg, nil
}

func SaveSmtpConfig(data *SmtpConfig) error {
	data.IsActive = true
	var existing SmtpConfig
	if err := config.DB.Order("id desc").First(&existing).Error; err == nil {
		data.ID = existing.ID
		data.CreatedAt = existing.CreatedAt
		// Jika password dikosongkan dari payload, gunakan password lama
		if data.Password == "" {
			data.Password = existing.Password
		}
		return config.DB.Save(data).Error
	}
	return config.DB.Create(data).Error
}
