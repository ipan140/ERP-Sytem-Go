package appraisals

func CreateAppraisalService(data *Appraisal) error {
	return CreateAppraisal(data)
}

func GetAllAppraisalService() ([]Appraisal, error) {
	return GetAllAppraisal()
}

func GetPaginatedAppraisalsService(offset, limit int, search string, employeeID string, state string) ([]Appraisal, int64, error) {
	return GetPaginatedAppraisals(offset, limit, search, employeeID, state)
}

func GetAppraisalByIDService(id uint) (*Appraisal, error) {
	return GetAppraisalByID(id)
}

func UpdateAppraisalService(data *Appraisal) error {
	return UpdateAppraisal(data)
}

func DeleteAppraisalService(id uint) error {
	return DeleteAppraisal(id)
}
