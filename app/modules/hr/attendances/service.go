package attendances

func CreateAttendanceService(data *Attendance) error {
	return CreateAttendance(data)
}

func GetAllAttendanceService() ([]Attendance, error) {
	return GetAllAttendance()
}

func GetAttendanceByIDService(id uint) (*Attendance, error) {
	return GetAttendanceByID(id)
}

func UpdateAttendanceService(data *Attendance) error {
	return UpdateAttendance(data)
}

func DeleteAttendanceService(id uint) error {
	return DeleteAttendance(id)
}

func CreateOvertimeService(data *Overtime) error { return CreateOvertime(data) }
func GetAllOvertimeService() ([]Overtime, error) { return GetAllOvertime() }
func GetOvertimeByIDService(id uint) (*Overtime, error) { return GetOvertimeByID(id) }
func UpdateOvertimeService(data *Overtime) error { return UpdateOvertime(data) }
func DeleteOvertimeService(id uint) error { return DeleteOvertime(id) }
