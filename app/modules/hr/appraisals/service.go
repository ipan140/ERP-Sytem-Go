package appraisals

func CreateAppraisalService(data *Appraisal) error {
	return CreateAppraisal(data)
}

func GetAllAppraisalService() ([]Appraisal, error) {
	return GetAllAppraisal()
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
