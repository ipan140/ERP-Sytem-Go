package inventory

import (
	"ERP-System/config"
)

func CreateProduct(data *Product) error {
	return config.DB.Create(data).Error
}

func GetAllProduct() ([]Product, error) {
	var list []Product
	err := config.DB.Find(&list).Error
	return list, err
}

func GetProductByID(id uint) (*Product, error) {
	var data Product
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateProduct(data *Product) error {
	return config.DB.Save(data).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&Product{}, id).Error
}
