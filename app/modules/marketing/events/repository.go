package events

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateEvent(data *Event) error {
	return config.DB.Create(data).Error
}

func GetAllEvent() ([]Event, error) {
	var list []Event
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetEventByID(id uint) (*Event, error) {
	var data Event
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateEvent(data *Event) error {
	return config.DB.Save(data).Error
}

func DeleteEvent(id uint) error {
	return config.DB.Delete(&Event{}, id).Error
}

func CreateEventTicket(data *EventTicket) error { return config.DB.Create(data).Error }
func GetAllEventTicket() ([]EventTicket, error) {
	var list []EventTicket
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetEventTicketByID(id uint) (*EventTicket, error) {
	var data EventTicket
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateEventTicket(data *EventTicket) error { return config.DB.Save(data).Error }
func DeleteEventTicket(id uint) error           { return config.DB.Delete(&EventTicket{}, id).Error }
