package sales_core

import (
	"ERP-System/config"
)

func CreateSaleOrder(data *SaleOrder) error {
	return config.DB.Create(data).Error
}

func GetAllSaleOrder() ([]SaleOrder, error) {
	var list []SaleOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSaleOrderByID(id uint) (*SaleOrder, error) {
	var data SaleOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSaleOrder(data *SaleOrder) error {
	return config.DB.Save(data).Error
}

func DeleteSaleOrder(id uint) error {
	return config.DB.Delete(&SaleOrder{}, id).Error
}
