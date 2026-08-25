package appraisals

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAppraisal(data *Appraisal) error {
	return config.DB.Create(data).Error
}

func GetAllAppraisal() ([]Appraisal, error) {
	var list []Appraisal
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetAppraisalByID(id uint) (*Appraisal, error) {
	var data Appraisal
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAppraisal(data *Appraisal) error {
	return config.DB.Save(data).Error
}

func DeleteAppraisal(id uint) error {
	return config.DB.Delete(&Appraisal{}, id).Error
}
