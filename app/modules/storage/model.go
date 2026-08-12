package storage

import (
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
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Attachment{})
}
