package website_builder

import (
	"ERP-System/config"
)

func CreatePage(data *Page) error {
	return config.DB.Create(data).Error
}

func GetAllPage() ([]Page, error) {
	var list []Page
	err := config.DB.Find(&list).Error
	return list, err
}

func GetPageByID(id uint) (*Page, error) {
	var data Page
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdatePage(data *Page) error {
	return config.DB.Save(data).Error
}

func DeletePage(id uint) error {
	return config.DB.Delete(&Page{}, id).Error
}
