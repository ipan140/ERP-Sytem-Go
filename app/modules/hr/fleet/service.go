package fleet

func CreateVehicleService(data *Vehicle) error {
	return CreateVehicle(data)
}

func GetAllVehicleService() ([]Vehicle, error) {
	return GetAllVehicle()
}

func GetVehicleByIDService(id uint) (*Vehicle, error) {
	return GetVehicleByID(id)
}

func UpdateVehicleService(data *Vehicle) error {
	return UpdateVehicle(data)
}

func DeleteVehicleService(id uint) error {
	return DeleteVehicle(id)
}
