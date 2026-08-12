package quality

func CreateQualityCheckService(data *QualityCheck) error {
	return CreateQualityCheck(data)
}

func GetAllQualityCheckService() ([]QualityCheck, error) {
	return GetAllQualityCheck()
}

func GetQualityCheckByIDService(id uint) (*QualityCheck, error) {
	return GetQualityCheckByID(id)
}

func UpdateQualityCheckService(data *QualityCheck) error {
	return UpdateQualityCheck(data)
}

func DeleteQualityCheckService(id uint) error {
	return DeleteQualityCheck(id)
}
