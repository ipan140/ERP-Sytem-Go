package iot

import (
	"ERP-System/config"
	"time"
)

type IoTDevice struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	DeviceName   string    `gorm:"type:varchar(255);default:''" json:"device_name"`
	Name         string    `gorm:"type:varchar(255);default:''" json:"name"`
	Model        string    `gorm:"type:varchar(100);default:''" json:"model"`
	Location     string    `gorm:"type:varchar(255);default:''" json:"location"`
	IPAddress    string    `gorm:"type:varchar(50);default:'192.168.1.200'" json:"ip_address"`
	Port         int       `gorm:"default:4370" json:"port"`
	SensorType   string    `gorm:"type:varchar(100);default:'Face 3D + Fingerprint'" json:"sensor_type"`
	Status       string    `gorm:"type:varchar(50);default:'Online'" json:"status"`
	PingMs       int       `gorm:"default:15" json:"ping_ms"`
	TodayRecords string    `gorm:"type:varchar(100);default:'0 Punch'" json:"today_records"`
	Icon         string    `gorm:"type:varchar(50);default:'📟'" json:"icon"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type AttendanceLog struct {
	ID                 uint      `gorm:"primaryKey" json:"id"`
	Timestamp          string    `gorm:"type:varchar(50)" json:"timestamp"`
	EmployeeName       string    `gorm:"type:varchar(255)" json:"employee_name"`
	NIK                string    `gorm:"type:varchar(100)" json:"nik"`
	Department         string    `gorm:"type:varchar(100)" json:"department"`
	DeviceLocation     string    `gorm:"type:varchar(255)" json:"device_location"`
	PunchType          string    `gorm:"type:varchar(50);default:'Check-In'" json:"punch_type"`
	VerificationMethod string    `gorm:"type:varchar(100);default:'Sidik Jari'" json:"verification_method"`
	CreatedAt          time.Time `json:"created_at"`
}

func (IoTDevice) TableName() string {
	return "setting.iot_devices"
}

func (AttendanceLog) TableName() string {
	return "setting.iot_attendance_logs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &IoTDevice{}, &AttendanceLog{})
}