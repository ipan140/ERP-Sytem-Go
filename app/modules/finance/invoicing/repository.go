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

func GetPaginatedInvoices(offset int, limit int, search string, status string) ([]Invoice, int64, error) {
	var list []Invoice
	var total int64

	query := config.DB.Model(&Invoice{}).Preload(clause.Associations)

	if status != "" && status != "all" && status != "All" && status != "Semua" {
		query = query.Where("invoices.state = ?", status)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN partners ON partners.id = invoices.partner_id").
			Where("invoices.name ILIKE ? OR partners.name ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("invoices.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
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
