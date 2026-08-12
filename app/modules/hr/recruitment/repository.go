package recruitment

import (
	"ERP-System/config"
)

func CreateJobApplicant(data *JobApplicant) error {
	return config.DB.Create(data).Error
}

func GetAllJobApplicant() ([]JobApplicant, error) {
	var list []JobApplicant
	err := config.DB.Find(&list).Error
	return list, err
}

func GetJobApplicantByID(id uint) (*JobApplicant, error) {
	var data JobApplicant
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateJobApplicant(data *JobApplicant) error {
	return config.DB.Save(data).Error
}

func DeleteJobApplicant(id uint) error {
	return config.DB.Delete(&JobApplicant{}, id).Error
}
