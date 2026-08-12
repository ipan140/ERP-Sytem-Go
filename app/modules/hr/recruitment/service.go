package recruitment

func CreateJobApplicantService(data *JobApplicant) error {
	return CreateJobApplicant(data)
}

func GetAllJobApplicantService() ([]JobApplicant, error) {
	return GetAllJobApplicant()
}

func GetJobApplicantByIDService(id uint) (*JobApplicant, error) {
	return GetJobApplicantByID(id)
}

func UpdateJobApplicantService(data *JobApplicant) error {
	return UpdateJobApplicant(data)
}

func DeleteJobApplicantService(id uint) error {
	return DeleteJobApplicant(id)
}
