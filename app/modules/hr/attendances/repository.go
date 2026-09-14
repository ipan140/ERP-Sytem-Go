package attendances

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAttendance(data *Attendance) error {
	return config.DB.Create(data).Error
}

func GetAllAttendance() ([]Attendance, error) {
	var list []Attendance
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedAttendances(offset, limit int, search string, employeeID string, date string) ([]Attendance, int64, error) {
	var list []Attendance
	var total int64

	query := config.DB.Model(&Attendance{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.attendances.employee_id = ?", employeeID)
	}

	if date != "" {
		query = query.Where("DATE(hrd.attendances.check_in) = ?", date)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN employees ON employees.id = hrd.attendances.employee_id").
			Where("employees.name ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.attendances.check_in DESC, hrd.attendances.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetAttendanceByID(id uint) (*Attendance, error) {
	var data Attendance
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAttendance(data *Attendance) error {
	return config.DB.Save(data).Error
}

func DeleteAttendance(id uint) error {
	return config.DB.Delete(&Attendance{}, id).Error
}

func CreateOvertime(data *Overtime) error { return config.DB.Create(data).Error }
func GetAllOvertime() ([]Overtime, error) {
	var list []Overtime
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetOvertimeByID(id uint) (*Overtime, error) {
	var data Overtime
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateOvertime(data *Overtime) error { return config.DB.Save(data).Error }
func DeleteOvertime(id uint) error        { return config.DB.Delete(&Overtime{}, id).Error }
