package timesheets

func CreateTimesheetService(data *Timesheet) error {
	return CreateTimesheet(data)
}

func GetAllTimesheetService() ([]Timesheet, error) {
	return GetAllTimesheet()
}

func GetTimesheetByIDService(id uint) (*Timesheet, error) {
	return GetTimesheetByID(id)
}

func UpdateTimesheetService(data *Timesheet) error {
	return UpdateTimesheet(data)
}

func DeleteTimesheetService(id uint) error {
	return DeleteTimesheet(id)
}
