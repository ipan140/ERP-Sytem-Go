package purchase

import (
	"ERP-System/config"
)

func CreatePurchaseOrder(data *PurchaseOrder) error {
	return config.DB.Create(data).Error
}

func GetAllPurchaseOrder() ([]PurchaseOrder, error) {
	var list []PurchaseOrder
	err := config.DB.Find(&list).Error
	return list, err
}

func GetPurchaseOrderByID(id uint) (*PurchaseOrder, error) {
	var data PurchaseOrder
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdatePurchaseOrder(data *PurchaseOrder) error {
	return config.DB.Save(data).Error
}

func DeletePurchaseOrder(id uint) error {
	return config.DB.Delete(&PurchaseOrder{}, id).Error
}
