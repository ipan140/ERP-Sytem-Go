package iot

func CreateIoTDeviceService(data *IoTDevice) error {
	return CreateIoTDevice(data)
}

func GetAllIoTDeviceService() ([]IoTDevice, error) {
	return GetAllIoTDevice()
}

func GetIoTDeviceByIDService(id uint) (*IoTDevice, error) {
	return GetIoTDeviceByID(id)
}

func UpdateIoTDeviceService(data *IoTDevice) error {
	return UpdateIoTDevice(data)
}

func DeleteIoTDeviceService(id uint) error {
	return DeleteIoTDevice(id)
}
