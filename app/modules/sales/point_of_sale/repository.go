package point_of_sale

import (
	"ERP-System/config"
)

func CreatePosSession(data *PosSession) error {
	return config.DB.Create(data).Error
}

func GetAllPosSession() ([]PosSession, error) {
	var list []PosSession
	err := config.DB.Find(&list).Error
	return list, err
}

func GetPosSessionByID(id uint) (*PosSession, error) {
	var data PosSession
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdatePosSession(data *PosSession) error {
	return config.DB.Save(data).Error
}

func DeletePosSession(id uint) error {
	return config.DB.Delete(&PosSession{}, id).Error
}
