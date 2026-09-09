package appointments

import (
	"ERP-System/config"
	"gorm.io/gorm/clause"
)

func CreateAppointment(data *Appointment) error {
	return config.DB.Create(data).Error
}

func GetAllAppointment() ([]Appointment, error) {
	var list []Appointment
	err := config.DB.Preload(clause.Associations).Order("date DESC, id DESC").Find(&list).Error
	return list, err
}

func GetPaginatedAppointments(offset int, limit int, search string, state string, employeeID uint, partnerID uint, companyID uint) ([]Appointment, int64, error) {
	var list []Appointment
	var total int64

	query := config.DB.Model(&Appointment{}).Preload(clause.Associations)

	if companyID > 0 {
		query = query.Where("company_id = ?", companyID)
	}
	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if partnerID > 0 {
		query = query.Where("partner_id = ?", partnerID)
	}
	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR notes ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("date DESC, id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
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
