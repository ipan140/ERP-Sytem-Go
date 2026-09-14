package documents

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateFinanceDocument(data *FinanceDocument) error {
	return config.DB.Create(data).Error
}

func GetAllFinanceDocument() ([]FinanceDocument, error) {
	var list []FinanceDocument
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedFinanceDocuments(offset int, limit int, search string) ([]FinanceDocument, int64, error) {
	var list []FinanceDocument
	var total int64

	query := config.DB.Model(&FinanceDocument{})
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetFinanceDocumentByID(id uint) (*FinanceDocument, error) {
	var data FinanceDocument
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateFinanceDocument(data *FinanceDocument) error {
	return config.DB.Save(data).Error
}

func DeleteFinanceDocument(id uint) error {
	return config.DB.Delete(&FinanceDocument{}, id).Error
}
