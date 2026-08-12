package ecommerce

import (
	"ERP-System/config"
)

func CreateCart(data *Cart) error {
	return config.DB.Create(data).Error
}

func GetAllCart() ([]Cart, error) {
	var list []Cart
	err := config.DB.Find(&list).Error
	return list, err
}

func GetCartByID(id uint) (*Cart, error) {
	var data Cart
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateCart(data *Cart) error {
	return config.DB.Save(data).Error
}

func DeleteCart(id uint) error {
	return config.DB.Delete(&Cart{}, id).Error
}
