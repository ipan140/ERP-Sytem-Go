package voip

func CreateCallRecordService(data *CallRecord) error {
	return CreateCallRecord(data)
}

func GetAllCallRecordService() ([]CallRecord, error) {
	return GetAllCallRecord()
}

func GetCallRecordByIDService(id uint) (*CallRecord, error) {
	return GetCallRecordByID(id)
}

func UpdateCallRecordService(data *CallRecord) error {
	return UpdateCallRecord(data)
}

func DeleteCallRecordService(id uint) error {
	return DeleteCallRecord(id)
}
