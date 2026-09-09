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

func GetPaginatedMaintenanceRequestsService(page, limit int, search, state, reqType string) ([]MaintenanceRequest, int64, error) {
	return GetPaginatedMaintenanceRequests(page, limit, search, state, reqType)
}

func GetMaintenanceSummaryService() (*MaintenanceSummary, error) {
	return GetMaintenanceSummary()
}

func CreateMaintenanceRequestWithSequenceService(req *CreateMaintenanceRequestDto) (*MaintenanceRequest, error) {
	return CreateMaintenanceRequestWithSequence(req)
}

func UpdateMaintenanceStateService(id uint, state string, duration *float64, notes *string) (*MaintenanceRequest, error) {
	return UpdateMaintenanceState(id, state, duration, notes)
}

func GetPaginatedEquipmentsService(page, limit int, search string) ([]MaintenanceEquipment, int64, error) {
	return GetPaginatedEquipments(page, limit, search)
}

func CreateMaintenanceEquipmentService(dto *CreateMaintenanceEquipmentDto) (*MaintenanceEquipment, error) {
	return CreateMaintenanceEquipment(dto)
}

func DeleteMaintenanceEquipmentService(id uint) error {
	return DeleteMaintenanceEquipment(id)
}

