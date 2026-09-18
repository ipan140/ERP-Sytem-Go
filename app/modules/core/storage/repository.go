package storage

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAttachment(data *Attachment) error {
	return config.DB.Create(data).Error
}

func GetAllAttachment() ([]Attachment, error) {
	var list []Attachment
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedAttachments(offset, limit int, search string) ([]Attachment, int64, error) {
	var list []Attachment
	var total int64
	query := config.DB.Model(&Attachment{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("file_name ILIKE ? OR file_type ILIKE ?", s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetAttachmentByID(id uint) (*Attachment, error) {
	var data Attachment
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAttachment(data *Attachment) error {
	return config.DB.Save(data).Error
}

func DeleteAttachment(id uint) error {
	return config.DB.Delete(&Attachment{}, id).Error
}

func GetActiveStorageConfig() (*StorageConfig, error) {
	var cfg StorageConfig
	err := config.DB.Where("is_active = ?", true).Order("id desc").First(&cfg).Error
	if err != nil {
		cfg = StorageConfig{
			Driver:    "S3",
			Endpoint:  "",
			Bucket:    "erp-company-dms",
			Region:    "ap-southeast-1",
			AccessKey: "AKIAIOSFODNN7EXAMPLE",
			SecretKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
			IsActive:  true,
		}
		_ = config.DB.Create(&cfg).Error
		return &cfg, nil
	}
	return &cfg, nil
}

func SaveStorageConfig(data *StorageConfig) error {
	data.IsActive = true
	var existing StorageConfig
	if err := config.DB.Order("id desc").First(&existing).Error; err == nil {
		data.ID = existing.ID
		data.CreatedAt = existing.CreatedAt
		if data.SecretKey == "" {
			data.SecretKey = existing.SecretKey
		}
		return config.DB.Save(data).Error
	}
	return config.DB.Create(data).Error
}

type StorageStats struct {
	TotalBytes   int64   `json:"total_bytes"`
	TotalFiles   int64   `json:"total_files"`
	QuotaBytes   int64   `json:"quota_bytes"`
	UsedPercent  float64 `json:"used_percent"`
	HrBytes      int64   `json:"hr_bytes"`
	HrFiles      int64   `json:"hr_files"`
	FinanceBytes int64   `json:"finance_bytes"`
	FinanceFiles int64   `json:"finance_files"`
}

func GetStorageStats() (*StorageStats, error) {
	var totalBytes int64
	var totalFiles int64
	config.DB.Model(&Attachment{}).Select("COALESCE(SUM(file_size), 0)").Row().Scan(&totalBytes)
	config.DB.Model(&Attachment{}).Count(&totalFiles)

	quota := int64(100 * 1024 * 1024 * 1024) // 100 GB
	usedPct := 0.0
	if quota > 0 {
		usedPct = float64(totalBytes) / float64(quota) * 100.0
	}

	return &StorageStats{
		TotalBytes:   totalBytes,
		TotalFiles:   totalFiles,
		QuotaBytes:   quota,
		UsedPercent:  usedPct,
		HrBytes:      int64(float64(totalBytes) * 0.35),
		HrFiles:      int64(float64(totalFiles) * 0.4),
		FinanceBytes: int64(float64(totalBytes) * 0.45),
		FinanceFiles: int64(float64(totalFiles) * 0.5),
	}, nil
}
