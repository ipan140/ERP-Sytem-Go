package plm

func GetPaginatedEcosService(page, limit int, search, state string) ([]PlmEco, int64, error) {
	return GetPaginatedEcos(page, limit, search, state)
}

func GetPlmSummaryService() (*PlmSummary, error) {
	return GetPlmSummary()
}

func CreateEcoWithSequenceService(req *CreateEcoRequest) (*PlmEco, error) {
	return CreateEcoWithSequence(req)
}

func UpdateEcoStateService(id uint, state string, approverID *uint) (*PlmEco, error) {
	return UpdateEcoState(id, state, approverID)
}

func GetAllEcoTypesService() ([]PlmEcoType, error) {
	return GetAllEcoTypes()
}

func GetEcoByIDService(id uint) (*PlmEco, error) {
	return GetEcoByID(id)
}

func DeleteEcoService(id uint) error {
	return DeleteEco(id)
}
