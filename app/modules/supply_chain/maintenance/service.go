package maintenance

func CreateMaintenanceRequestService(data *MaintenanceRequest) error {
	return CreateMaintenanceRequest(data)
}

func GetAllMaintenanceRequestService() ([]MaintenanceRequest, error) {
	return GetAllMaintenanceRequest()
}

func GetMaintenanceRequestByIDService(id uint) (*MaintenanceRequest, error) {
	return GetMaintenanceRequestByID(id)
}

func UpdateMaintenanceRequestService(data *MaintenanceRequest) error {
	return UpdateMaintenanceRequest(data)
}

func DeleteMaintenanceRequestService(id uint) error {
	return DeleteMaintenanceRequest(id)
}
