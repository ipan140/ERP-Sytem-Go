package storage

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"time"
)

type Attachment struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	FileName     string    `gorm:"type:varchar(255);not null" json:"file_name"`
	FileType     string    `gorm:"type:varchar(100)" json:"file_type"`
	FileSize     int64     `json:"file_size"`
	FilePath     string    `gorm:"type:text;not null" json:"file_path"`
	UploadedByID uint      `json:"uploaded_by_id"` // Siapa yang mengunggah
	UploadedBy *auth.User `gorm:"foreignKey:UploadedByID" json:"uploadedby,omitempty"` // Cross-module relation
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}


func (Attachment) TableName() string {
	return "setting.attachments"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Attachment{})
}
