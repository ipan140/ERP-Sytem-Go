package field_service

func CreateFieldServiceTaskService(data *FieldServiceTask) error {
	return CreateFieldServiceTask(data)
}

func GetAllFieldServiceTaskService() ([]FieldServiceTask, error) {
	return GetAllFieldServiceTask()
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
