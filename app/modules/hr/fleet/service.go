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

func CreateVehicleLogContractService(data *VehicleLogContract) error {
	return CreateVehicleLogContract(data)
}
func GetAllVehicleLogContractService() ([]VehicleLogContract, error) {
	return GetAllVehicleLogContract()
}
func GetVehicleLogContractByIDService(id uint) (*VehicleLogContract, error) {
	return GetVehicleLogContractByID(id)
}
func UpdateVehicleLogContractService(data *VehicleLogContract) error {
	return UpdateVehicleLogContract(data)
}
func DeleteVehicleLogContractService(id uint) error { return DeleteVehicleLogContract(id) }

func CreateVehicleLogFuelService(data *VehicleLogFuel) error        { return CreateVehicleLogFuel(data) }
func GetAllVehicleLogFuelService() ([]VehicleLogFuel, error)        { return GetAllVehicleLogFuel() }
func GetVehicleLogFuelByIDService(id uint) (*VehicleLogFuel, error) { return GetVehicleLogFuelByID(id) }
func UpdateVehicleLogFuelService(data *VehicleLogFuel) error        { return UpdateVehicleLogFuel(data) }
func DeleteVehicleLogFuelService(id uint) error                     { return DeleteVehicleLogFuel(id) }

func CreateVehicleLogServicesService(data *VehicleLogServices) error {
	return CreateVehicleLogServices(data)
}
func GetAllVehicleLogServicesService() ([]VehicleLogServices, error) {
	return GetAllVehicleLogServices()
}
func GetVehicleLogServicesByIDService(id uint) (*VehicleLogServices, error) {
	return GetVehicleLogServicesByID(id)
}
func UpdateVehicleLogServicesService(data *VehicleLogServices) error {
	return UpdateVehicleLogServices(data)
}
func DeleteVehicleLogServicesService(id uint) error { return DeleteVehicleLogServices(id) }
