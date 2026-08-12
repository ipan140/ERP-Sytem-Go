package iot

import (
	"ERP-System/config"
)

func CreateIoTDevice(data *IoTDevice) error {
	return config.DB.Create(data).Error
}

func GetAllIoTDevice() ([]IoTDevice, error) {
	var list []IoTDevice
	err := config.DB.Find(&list).Error
	return list, err
}

func GetIoTDeviceByID(id uint) (*IoTDevice, error) {
	var data IoTDevice
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateIoTDevice(data *IoTDevice) error {
	return config.DB.Save(data).Error
}

func DeleteIoTDevice(id uint) error {
	return config.DB.Delete(&IoTDevice{}, id).Error
}
