package recruitment

func CreateApplicantService(data *Applicant) error {
	return CreateApplicant(data)
}

func GetAllApplicantService() ([]Applicant, error) {
	return GetAllApplicant()
}

func GetPaginatedApplicantService(offset, limit int, search string, jobPositionID string, stageID string, state string) ([]Applicant, int64, error) {
	return GetPaginatedApplicants(offset, limit, search, jobPositionID, stageID, state)
}

func GetApplicantByIDService(id uint) (*Applicant, error) {
	return GetApplicantByID(id)
}

func UpdateApplicantService(data *Applicant) error {
	return UpdateApplicant(data)
}

func DeleteApplicantService(id uint) error {
	return DeleteApplicant(id)
}
