package invoicing

import (
	"ERP-System/config"
)

func CreateInvoice(data *Invoice) error {
	return config.DB.Create(data).Error
}

func GetAllInvoice() ([]Invoice, error) {
	var list []Invoice
	err := config.DB.Find(&list).Error
	return list, err
}

func GetInvoiceByID(id uint) (*Invoice, error) {
	var data Invoice
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateInvoice(data *Invoice) error {
	return config.DB.Save(data).Error
}

func DeleteInvoice(id uint) error {
	return config.DB.Delete(&Invoice{}, id).Error
}
