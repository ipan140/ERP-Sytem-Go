package invoicing

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateInvoice(data *Invoice) error {
	return config.DB.Create(data).Error
}

func GetAllInvoice() ([]Invoice, error) {
	var list []Invoice
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetInvoiceByID(id uint) (*Invoice, error) {
	var data Invoice
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateInvoice(data *Invoice) error {
	return config.DB.Save(data).Error
}

func DeleteInvoice(id uint) error {
	return config.DB.Delete(&Invoice{}, id).Error
}

func CreatePaymentTermLine(data *PaymentTermLine) error { return config.DB.Create(data).Error }
func GetAllPaymentTermLine() ([]PaymentTermLine, error) {
	var list []PaymentTermLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPaymentTermLineByID(id uint) (*PaymentTermLine, error) {
	var data PaymentTermLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePaymentTermLine(data *PaymentTermLine) error { return config.DB.Save(data).Error }
func DeletePaymentTermLine(id uint) error               { return config.DB.Delete(&PaymentTermLine{}, id).Error }

func CreateTaxRepartitionLine(data *TaxRepartitionLine) error { return config.DB.Create(data).Error }
func GetAllTaxRepartitionLine() ([]TaxRepartitionLine, error) {
	var list []TaxRepartitionLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetTaxRepartitionLineByID(id uint) (*TaxRepartitionLine, error) {
	var data TaxRepartitionLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateTaxRepartitionLine(data *TaxRepartitionLine) error { return config.DB.Save(data).Error }
func DeleteTaxRepartitionLine(id uint) error {
	return config.DB.Delete(&TaxRepartitionLine{}, id).Error
}
