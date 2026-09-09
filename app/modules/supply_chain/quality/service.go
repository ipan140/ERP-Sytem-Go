package quality

func GetPaginatedQualityChecksService(offset, limit int, search, result string, productID uint) ([]QualityCheck, int64, error) {
	return GetPaginatedQualityChecks(offset, limit, search, result, productID)
}

func GetQualitySummaryService() (QualitySummary, error) {
	return GetQualitySummary()
}

func CreateQualityCheckWithSequenceService(req CreateQualityCheckRequest) (*QualityCheck, error) {
	return CreateQualityCheckWithSequence(req)
}

func ProcessQualityCheckService(id uint, req ProcessQCRequest, inspectorID *uint) (*QualityCheck, error) {
	return ProcessQualityCheck(id, req, inspectorID)
}

func GetPaginatedQualityPointsService(offset, limit int, search string) ([]QualityPoint, int64, error) {
	return GetPaginatedQualityPoints(offset, limit, search)
}

func CreateQualityPointService(data *QualityPoint) error {
	return CreateQualityPoint(data)
}

func GetAllQualityPointsService() ([]QualityPoint, error) {
	return GetAllQualityPoints()
}

func DeleteQualityPointService(id uint) error {
	return DeleteQualityPoint(id)
}

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
