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

// StorageConfig menyimpan konfigurasi Cloud Storage (S3, MinIO, Local)
type StorageConfig struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Driver    string    `gorm:"type:varchar(50);default:'S3'" json:"driver"` // S3, LOCAL, GCS
	Endpoint  string    `gorm:"type:varchar(255)" json:"endpoint"`
	Bucket    string    `gorm:"type:varchar(100);default:'erp-company-dms'" json:"bucket"`
	Region    string    `gorm:"type:varchar(50);default:'ap-southeast-1'" json:"region"`
	AccessKey string    `gorm:"type:varchar(255)" json:"access_key"`
	SecretKey string    `gorm:"type:varchar(255)" json:"secret_key"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (StorageConfig) TableName() string {
	return "setting.storage_configs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Attachment{}, &StorageConfig{})
}
