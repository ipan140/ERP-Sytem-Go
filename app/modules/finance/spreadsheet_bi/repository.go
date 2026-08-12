package spreadsheet_bi

import (
	"ERP-System/config"
)

func CreateSpreadsheet(data *Spreadsheet) error {
	return config.DB.Create(data).Error
}

func GetAllSpreadsheet() ([]Spreadsheet, error) {
	var list []Spreadsheet
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSpreadsheetByID(id uint) (*Spreadsheet, error) {
	var data Spreadsheet
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSpreadsheet(data *Spreadsheet) error {
	return config.DB.Save(data).Error
}

func DeleteSpreadsheet(id uint) error {
	return config.DB.Delete(&Spreadsheet{}, id).Error
}
