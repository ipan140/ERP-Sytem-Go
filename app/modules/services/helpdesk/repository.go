package helpdesk

import (
	"ERP-System/config"
)

func CreateTicket(data *Ticket) error {
	return config.DB.Create(data).Error
}

func GetAllTicket() ([]Ticket, error) {
	var list []Ticket
	err := config.DB.Find(&list).Error
	return list, err
}

func GetTicketByID(id uint) (*Ticket, error) {
	var data Ticket
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateTicket(data *Ticket) error {
	return config.DB.Save(data).Error
}

func DeleteTicket(id uint) error {
	return config.DB.Delete(&Ticket{}, id).Error
}
