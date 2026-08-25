package appointments

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAppointment(data *Appointment) error {
	return config.DB.Create(data).Error
}

func GetAllAppointment() ([]Appointment, error) {
	var list []Appointment
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetAppointmentByID(id uint) (*Appointment, error) {
	var data Appointment
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAppointment(data *Appointment) error {
	return config.DB.Save(data).Error
}

func DeleteAppointment(id uint) error {
	return config.DB.Delete(&Appointment{}, id).Error
}
