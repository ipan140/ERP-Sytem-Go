package surveys

func CreateSurveyService(data *Survey) error {
	return CreateSurvey(data)
}

func GetAllSurveyService() ([]Survey, error) {
	return GetAllSurvey()
}

func GetSurveyByIDService(id uint) (*Survey, error) {
	return GetSurveyByID(id)
}

func UpdateSurveyService(data *Survey) error {
	return UpdateSurvey(data)
}

func DeleteSurveyService(id uint) error {
	return DeleteSurvey(id)
}

func RecordSurveyResponseService(id uint, rating int) (*Survey, error) {
	return RecordSurveyResponse(id, rating)
}

