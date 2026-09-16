package website_builder

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreatePage(data *Page) error {
	return config.DB.Create(data).Error
}

func GetAllPage() ([]Page, error) {
	var list []Page
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedPages(offset, limit int, search string) ([]Page, int64, error) {
	var list []Page
	var total int64
	query := config.DB.Model(&Page{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ?", s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetPageByID(id uint) (*Page, error) {
	var data Page
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdatePage(data *Page) error {
	return config.DB.Save(data).Error
}

func DeletePage(id uint) error {
	return config.DB.Delete(&Page{}, id).Error
}
