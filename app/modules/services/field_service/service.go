package field_service

func CreateFieldServiceTaskService(data *FieldServiceTask) error {
	return CreateFieldServiceTask(data)
}

func GetAllFieldServiceTaskService() ([]FieldServiceTask, error) {
	return GetAllFieldServiceTask()
}

func GetPaginatedFieldServiceTaskService(offset int, limit int, search string, state string, priority string, employeeID uint, companyID uint) ([]FieldServiceTask, int64, error) {
	return GetPaginatedFieldServiceTasks(offset, limit, search, state, priority, employeeID, companyID)
}

func GetFieldServiceTaskByIDService(id uint) (*FieldServiceTask, error) {
	return GetFieldServiceTaskByID(id)
}

func UpdateFieldServiceTaskService(data *FieldServiceTask) error {
	return UpdateFieldServiceTask(data)
}

func DeleteFieldServiceTaskService(id uint) error {
	return DeleteFieldServiceTask(id)
}

// Fase 2: e-BAST Validation Gate
func ValidateBastFieldServiceTaskService(id uint, validatorID uint) error {
	return ValidateBastFieldServiceTask(id, validatorID)
}

// Fase 4: GPS Geotagging
func RecordGPSCheckInService(id uint, lat float64, lng float64) (*FieldServiceTask, error) {
	return RecordGPSCheckIn(id, lat, lng)
}

func RecordGPSCheckOutService(id uint, lat float64, lng float64) (*FieldServiceTask, error) {
	return RecordGPSCheckOut(id, lat, lng)
}

