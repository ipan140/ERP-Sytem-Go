package iot

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateIoTDevice(data *IoTDevice) error {
	if data.DeviceName == "" && data.Name != "" {
		data.DeviceName = data.Name
	}
	if data.Name == "" && data.DeviceName != "" {
		data.Name = data.DeviceName
	}
	return config.DB.Create(data).Error
}

func GetAllIoTDevice() ([]IoTDevice, error) {
	var list []IoTDevice
	err := config.DB.Preload(clause.Associations).Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultIoTDevices()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultIoTDevices() error {
	defaults := []IoTDevice{
		{
			DeviceName:   "ZKTeco uFace 800 - Lobi Utama",
			Name:         "ZKTeco uFace 800 - Lobi Utama",
			Model:        "ZKTeco uFace800 Plus",
			Location:     "Kantor Pusat HQ - Lobi Utama Lt. 1",
			IPAddress:    "192.168.1.201",
			Port:         4370,
			SensorType:   "Face 3D + Fingerprint",
			Status:       "Online",
			PingMs:       12,
			TodayRecords: "194 Punch",
			Icon:         "👤",
		},
		{
			DeviceName:   "ZKTeco SilkBio - Pintu Masuk Gudang",
			Name:         "ZKTeco SilkBio - Pintu Masuk Gudang",
			Model:        "ZKTeco SilkBio-101TC",
			Location:     "Kawasan Industri Cikarang Barat MM2100",
			IPAddress:    "192.168.2.202",
			Port:         4370,
			SensorType:   "Fingerprint Biometrik",
			Status:       "Online",
			PingMs:       24,
			TodayRecords: "72 Punch",
			Icon:         "🏭",
		},
		{
			DeviceName:   "Solution X105 - Cabang Surabaya",
			Name:         "Solution X105 - Cabang Surabaya",
			Model:        "Solution X105 Standalone",
			Location:     "Kantor Perwakilan Surabaya Lt. 2",
			IPAddress:    "10.10.1.50",
			Port:         4370,
			SensorType:   "RFID Card + Fingerprint",
			Status:       "Online",
			PingMs:       38,
			TodayRecords: "46 Punch",
			Icon:         "🏢",
		},
		{
			DeviceName:   "Sensor IoT DHT22 - Cold Storage Farmasi",
			Name:         "Sensor IoT DHT22 - Cold Storage Farmasi",
			Model:        "ESP32 MQTT Modbus RTU",
			Location:     "Gudang Cikarang - Ruang Dingin A-1",
			IPAddress:    "192.168.2.215",
			Port:         1883,
			SensorType:   "Sensor Suhu & Kelembaban (DHT22)",
			Status:       "Online",
			PingMs:       15,
			TodayRecords: "Suhu 4.2°C • Kelembaban 45%",
			Icon:         "❄️",
		},
	}
	for _, d := range defaults {
		var existing IoTDevice
		if err := config.DB.Where("device_name = ? OR name = ?", d.DeviceName, d.Name).First(&existing).Error; err != nil {
			_ = config.DB.Create(&d).Error
		}
	}
	return nil
}

func GetAllAttendanceLogs() ([]AttendanceLog, error) {
	var list []AttendanceLog
	err := config.DB.Order("id desc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultAttendanceLogs()
		_ = config.DB.Order("id desc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultAttendanceLogs() error {
	defaults := []AttendanceLog{
		{Timestamp: "17:31:04", EmployeeName: "Ahmad Fauzi", NIK: "EMP-0012", Department: "Logistik & Gudang", DeviceLocation: "Gudang Cikarang", PunchType: "Check-Out", VerificationMethod: "Sidik Jari (98%)"},
		{Timestamp: "17:30:45", EmployeeName: "Siti Rahmawati", NIK: "EMP-0045", Department: "HR & Kepegawaian", DeviceLocation: "Lobi Utama HQ", PunchType: "Check-Out", VerificationMethod: "Face Recognition 3D"},
		{Timestamp: "17:28:12", EmployeeName: "Budi Santoso", NIK: "EMP-0008", Department: "Finance & Accounting", DeviceLocation: "Lobi Utama HQ", PunchType: "Check-Out", VerificationMethod: "Face Recognition 3D"},
		{Timestamp: "08:42:19", EmployeeName: "Rian Hidayat", NIK: "EMP-0098", Department: "Sales & Marketing", DeviceLocation: "Cabang Surabaya", PunchType: "Check-In", VerificationMethod: "RFID Card + Finger"},
		{Timestamp: "08:35:50", EmployeeName: "Dewi Lestari", NIK: "EMP-0034", Department: "Supply Chain Ops", DeviceLocation: "Lobi Utama HQ", PunchType: "Check-In", VerificationMethod: "Face Recognition 3D"},
		{Timestamp: "08:14:02", EmployeeName: "Hendro Wijaya", NIK: "EMP-0019", Department: "Security & Facility", DeviceLocation: "Gudang Cikarang", PunchType: "Check-In", VerificationMethod: "Sidik Jari (100%)"},
	}
	for _, l := range defaults {
		_ = config.DB.Create(&l).Error
	}
	return nil
}

func GetIoTDeviceByID(id uint) (*IoTDevice, error) {
	var data IoTDevice
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateIoTDevice(data *IoTDevice) error {
	return config.DB.Save(data).Error
}

func DeleteIoTDevice(id uint) error {
	return config.DB.Delete(&IoTDevice{}, id).Error
}
