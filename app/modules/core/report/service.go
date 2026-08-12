package report

func CreateReportService(data *Report) error {
	return CreateReport(data)
}

func GetAllReportService() ([]Report, error) {
	return GetAllReport()
}

func GetReportByIDService(id uint) (*Report, error) {
	return GetReportByID(id)
}

func UpdateReportService(data *Report) error {
	return UpdateReport(data)
}

func DeleteReportService(id uint) error {
	return DeleteReport(id)
}
