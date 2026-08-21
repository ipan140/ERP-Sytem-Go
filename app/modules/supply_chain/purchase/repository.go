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

func CreatePurchaseRequisition(data *PurchaseRequisition) error { return config.DB.Create(data).Error }
func GetAllPurchaseRequisition() ([]PurchaseRequisition, error) {
	var list []PurchaseRequisition
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPurchaseRequisitionByID(id uint) (*PurchaseRequisition, error) {
	var data PurchaseRequisition
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePurchaseRequisition(data *PurchaseRequisition) error { return config.DB.Save(data).Error }
func DeletePurchaseRequisition(id uint) error {
	return config.DB.Delete(&PurchaseRequisition{}, id).Error
}

func CreateProductSupplierInfo(data *ProductSupplierInfo) error { return config.DB.Create(data).Error }
func GetAllProductSupplierInfo() ([]ProductSupplierInfo, error) {
	var list []ProductSupplierInfo
	err := config.DB.Find(&list).Error
	return list, err
}
func GetProductSupplierInfoByID(id uint) (*ProductSupplierInfo, error) {
	var data ProductSupplierInfo
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateProductSupplierInfo(data *ProductSupplierInfo) error { return config.DB.Save(data).Error }
func DeleteProductSupplierInfo(id uint) error {
	return config.DB.Delete(&ProductSupplierInfo{}, id).Error
}

func CreatePurchaseOrderLine(data *PurchaseOrderLine) error { return config.DB.Create(data).Error }
func GetAllPurchaseOrderLine() ([]PurchaseOrderLine, error) {
	var list []PurchaseOrderLine
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPurchaseOrderLineByID(id uint) (*PurchaseOrderLine, error) {
	var data PurchaseOrderLine
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePurchaseOrderLine(data *PurchaseOrderLine) error { return config.DB.Save(data).Error }
func DeletePurchaseOrderLine(id uint) error                 { return config.DB.Delete(&PurchaseOrderLine{}, id).Error }
