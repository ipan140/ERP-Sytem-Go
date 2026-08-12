package crm

func CreateLeadService(data *Lead) error {
	return CreateLead(data)
}

func GetAllLeadService() ([]Lead, error) {
	return GetAllLead()
}

func GetLeadByIDService(id uint) (*Lead, error) {
	return GetLeadByID(id)
}

func UpdateLeadService(data *Lead) error {
	return UpdateLead(data)
}

func DeleteLeadService(id uint) error {
	return DeleteLead(id)
}
