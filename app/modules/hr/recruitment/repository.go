package recruitment

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateApplicant(data *Applicant) error {
	return config.DB.Create(data).Error
}

func GetAllApplicant() ([]Applicant, error) {
	var list []Applicant
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetApplicantByID(id uint) (*Applicant, error) {
	var data Applicant
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateApplicant(data *Applicant) error {
	return config.DB.Save(data).Error
}

func DeleteApplicant(id uint) error {
	return config.DB.Delete(&Applicant{}, id).Error
}
