package surveys

import (
	"ERP-System/config"
)

func CreateSurvey(data *Survey) error {
	return config.DB.Create(data).Error
}

func GetAllSurvey() ([]Survey, error) {
	var list []Survey
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSurveyByID(id uint) (*Survey, error) {
	var data Survey
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSurvey(data *Survey) error {
	return config.DB.Save(data).Error
}

func DeleteSurvey(id uint) error {
	return config.DB.Delete(&Survey{}, id).Error
}
