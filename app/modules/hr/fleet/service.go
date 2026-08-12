package fleet

func CreateFleetVehicleService(data *FleetVehicle) error {
	return CreateFleetVehicle(data)
}

func GetAllFleetVehicleService() ([]FleetVehicle, error) {
	return GetAllFleetVehicle()
}

func GetFleetVehicleByIDService(id uint) (*FleetVehicle, error) {
	return GetFleetVehicleByID(id)
}

func UpdateFleetVehicleService(data *FleetVehicle) error {
	return UpdateFleetVehicle(data)
}

func DeleteFleetVehicleService(id uint) error {
	return DeleteFleetVehicle(id)
}
