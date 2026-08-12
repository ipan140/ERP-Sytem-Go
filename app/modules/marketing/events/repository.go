package events

import (
	"ERP-System/config"
)

func CreateEvent(data *Event) error {
	return config.DB.Create(data).Error
}

func GetAllEvent() ([]Event, error) {
	var list []Event
	err := config.DB.Find(&list).Error
	return list, err
}

func GetEventByID(id uint) (*Event, error) {
	var data Event
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateEvent(data *Event) error {
	return config.DB.Save(data).Error
}

func DeleteEvent(id uint) error {
	return config.DB.Delete(&Event{}, id).Error
}
