package attendances

import (
	"ERP-System/config"
)

func CreateAttendance(data *Attendance) error {
	return config.DB.Create(data).Error
}

func GetAllAttendance() ([]Attendance, error) {
	var list []Attendance
	err := config.DB.Find(&list).Error
	return list, err
}

func GetAttendanceByID(id uint) (*Attendance, error) {
	var data Attendance
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateAttendance(data *Attendance) error {
	return config.DB.Save(data).Error
}

func DeleteAttendance(id uint) error {
	return config.DB.Delete(&Attendance{}, id).Error
}

func CreateOvertime(data *Overtime) error { return config.DB.Create(data).Error }
func GetAllOvertime() ([]Overtime, error) { var list []Overtime; err := config.DB.Find(&list).Error; return list, err }
func GetOvertimeByID(id uint) (*Overtime, error) { var data Overtime; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateOvertime(data *Overtime) error { return config.DB.Save(data).Error }
func DeleteOvertime(id uint) error { return config.DB.Delete(&Overtime{}, id).Error }
