package discuss

import (
	"ERP-System/config"
)

func CreateChannel(data *Channel) error {
	return config.DB.Create(data).Error
}

func GetAllChannel() ([]Channel, error) {
	var list []Channel
	err := config.DB.Find(&list).Error
	return list, err
}

func GetChannelByID(id uint) (*Channel, error) {
	var data Channel
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateChannel(data *Channel) error {
	return config.DB.Save(data).Error
}

func DeleteChannel(id uint) error {
	return config.DB.Delete(&Channel{}, id).Error
}
