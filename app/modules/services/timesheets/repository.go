package timesheets

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateTimesheet(data *Timesheet) error {
	return config.DB.Create(data).Error
}

func GetAllTimesheet() ([]Timesheet, error) {
	var list []Timesheet
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetTimesheetByID(id uint) (*Timesheet, error) {
	var data Timesheet
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateTimesheet(data *Timesheet) error {
	return config.DB.Save(data).Error
}

func DeleteTimesheet(id uint) error {
	return config.DB.Delete(&Timesheet{}, id).Error
}
