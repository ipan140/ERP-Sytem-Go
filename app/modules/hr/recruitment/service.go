package recruitment

func CreateApplicantService(data *Applicant) error {
	return CreateApplicant(data)
}

func GetAllApplicantService() ([]Applicant, error) {
	return GetAllApplicant()
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
