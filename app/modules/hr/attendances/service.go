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
