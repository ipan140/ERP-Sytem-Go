package whatsapp

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateWaTemplate(data *WaTemplate) error {
	return config.DB.Create(data).Error
}

func GetAllWaTemplate() ([]WaTemplate, error) {
	var list []WaTemplate
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedWaTemplate(offset, limit int, search string) ([]WaTemplate, int64, error) {
	var list []WaTemplate
	var total int64
	db := config.DB.Model(&WaTemplate{})
	if search != "" {
		db = db.Where("name ILIKE ?", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Preload(clause.Associations).Order("id asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetWaTemplateByID(id uint) (*WaTemplate, error) {
	var data WaTemplate
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateWaTemplate(data *WaTemplate) error {
	return config.DB.Save(data).Error
}

func DeleteWaTemplate(id uint) error {
	return config.DB.Delete(&WaTemplate{}, id).Error
}
