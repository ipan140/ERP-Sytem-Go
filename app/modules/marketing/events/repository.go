package events

import (
	"errors"

	"ERP-System/config"
	"gorm.io/gorm/clause"
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

func CreateEventTicket(data *EventTicket) error {
	// 1. Jika customer_id bernilai 0, ubah jadi nil agar tidak melanggar foreign key
	if data.CustomerID != nil && *data.CustomerID == 0 {
		data.CustomerID = nil
	}

	// 2. Validasi Kuota Kapasitas Acara
	var event Event
	if err := config.DB.First(&event, data.EventID).Error; err != nil {
		return err
	}

	var currentCount int64
	config.DB.Model(&EventTicket{}).Where("event_id = ?", data.EventID).Count(&currentCount)

	if event.MaxCapacity > 0 && int(currentCount) >= event.MaxCapacity {
		return errors.New("Kapasitas acara sudah penuh! Tidak dapat menerbitkan tiket baru.")
	}

	return config.DB.Create(data).Error
}
func GetAllEventTicket() ([]EventTicket, error) {
	var list []EventTicket
	err := config.DB.Preload("Event").Preload("Customer").Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetEventTicketByID(id uint) (*EventTicket, error) {
	var data EventTicket
	err := config.DB.Preload("Event").Preload("Customer").Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateEventTicket(data *EventTicket) error { return config.DB.Save(data).Error }
func DeleteEventTicket(id uint) error           { return config.DB.Delete(&EventTicket{}, id).Error }

func ScanTicketByBarcode(barcode string) (*EventTicket, error) {
	var ticket EventTicket
	if err := config.DB.Preload(clause.Associations).Where("barcode = ?", barcode).First(&ticket).Error; err != nil {
		return nil, errors.New("Tiket dengan barcode ini tidak ditemukan!")
	}

	if ticket.IsScanned {
		return &ticket, errors.New("TIKET SUDAH PERNAH DIGUNAKAN! Presensi ditolak (Mencegah duplikasi).")
	}

	ticket.IsScanned = true
	if err := config.DB.Save(&ticket).Error; err != nil {
		return nil, err
	}

	return &ticket, nil
}
