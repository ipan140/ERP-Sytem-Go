package point_of_sale

func CreatePosSessionService(data *PosSession) error {
	return CreatePosSession(data)
}

func GetAllPosSessionService() ([]PosSession, error) {
	return GetAllPosSession()
}

func GetPosSessionByIDService(id uint) (*PosSession, error) {
	return GetPosSessionByID(id)
}

func UpdatePosSessionService(data *PosSession) error {
	return UpdatePosSession(data)
}

func DeletePosSessionService(id uint) error {
	return DeletePosSession(id)
}
