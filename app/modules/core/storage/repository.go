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
