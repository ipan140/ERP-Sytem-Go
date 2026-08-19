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

func CreateHelpdeskSLA(data *HelpdeskSLA) error { return config.DB.Create(data).Error }
func GetAllHelpdeskSLA() ([]HelpdeskSLA, error) { var list []HelpdeskSLA; err := config.DB.Find(&list).Error; return list, err }
func GetHelpdeskSLAByID(id uint) (*HelpdeskSLA, error) { var data HelpdeskSLA; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateHelpdeskSLA(data *HelpdeskSLA) error { return config.DB.Save(data).Error }
func DeleteHelpdeskSLA(id uint) error { return config.DB.Delete(&HelpdeskSLA{}, id).Error }

func CreateHelpdeskCannedResponse(data *HelpdeskCannedResponse) error { return config.DB.Create(data).Error }
func GetAllHelpdeskCannedResponse() ([]HelpdeskCannedResponse, error) { var list []HelpdeskCannedResponse; err := config.DB.Find(&list).Error; return list, err }
func GetHelpdeskCannedResponseByID(id uint) (*HelpdeskCannedResponse, error) { var data HelpdeskCannedResponse; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateHelpdeskCannedResponse(data *HelpdeskCannedResponse) error { return config.DB.Save(data).Error }
func DeleteHelpdeskCannedResponse(id uint) error { return config.DB.Delete(&HelpdeskCannedResponse{}, id).Error }

