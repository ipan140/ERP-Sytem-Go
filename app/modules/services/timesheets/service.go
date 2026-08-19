package timesheets

import "ERP-System/app/modules/services/project"

func CreateTimesheetService(data *Timesheet) error {
	if data.IsBillable { project.ProcessBillableTimesheet(data.ID) }
	
	return CreateTimesheet(data)
}

func GetAllTimesheetService() ([]Timesheet, error) {
	return GetAllTimesheet()
}

func GetTimesheetByIDService(id uint) (*Timesheet, error) {
	return GetTimesheetByID(id)
}

func UpdateTimesheetService(data *Timesheet) error {
	if data.IsBillable { project.ProcessBillableTimesheet(data.ID) }
	
	return UpdateTimesheet(data)
}

func DeleteTimesheetService(id uint) error {
	return DeleteTimesheet(id)
}


