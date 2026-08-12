package mailer

func CreateEmailLogService(data *EmailLog) error {
	return CreateEmailLog(data)
}

func GetAllEmailLogService() ([]EmailLog, error) {
	return GetAllEmailLog()
}

func GetEmailLogByIDService(id uint) (*EmailLog, error) {
	return GetEmailLogByID(id)
}

func UpdateEmailLogService(data *EmailLog) error {
	return UpdateEmailLog(data)
}

func DeleteEmailLogService(id uint) error {
	return DeleteEmailLog(id)
}
