package documents

import (
	"ERP-System/config"
)

func CreateFinanceDocument(data *FinanceDocument) error {
	return config.DB.Create(data).Error
}

func GetAllFinanceDocument() ([]FinanceDocument, error) {
	var list []FinanceDocument
	err := config.DB.Find(&list).Error
	return list, err
}

func GetFinanceDocumentByID(id uint) (*FinanceDocument, error) {
	var data FinanceDocument
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateFinanceDocument(data *FinanceDocument) error {
	return config.DB.Save(data).Error
}

func DeleteFinanceDocument(id uint) error {
	return config.DB.Delete(&FinanceDocument{}, id).Error
}
