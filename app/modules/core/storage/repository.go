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
