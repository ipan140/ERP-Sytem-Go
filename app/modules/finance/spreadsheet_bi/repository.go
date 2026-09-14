package spreadsheet_bi

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSpreadsheet(data *Spreadsheet) error {
	return config.DB.Create(data).Error
}

func GetAllSpreadsheet() ([]Spreadsheet, error) {
	var list []Spreadsheet
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedSpreadsheets(offset int, limit int, search string) ([]Spreadsheet, int64, error) {
	var list []Spreadsheet
	var total int64

	query := config.DB.Model(&Spreadsheet{})
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

func GetSpreadsheetByID(id uint) (*Spreadsheet, error) {
	var data Spreadsheet
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSpreadsheet(data *Spreadsheet) error {
	return config.DB.Save(data).Error
}

func DeleteSpreadsheet(id uint) error {
	return config.DB.Delete(&Spreadsheet{}, id).Error
}
