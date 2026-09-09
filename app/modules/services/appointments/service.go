package appointments

func CreateAppointmentService(data *Appointment) error {
	return CreateAppointment(data)
}

func GetAllAppointmentService() ([]Appointment, error) {
	return GetAllAppointment()
}

func GetPaginatedAppointmentService(offset int, limit int, search string, state string, employeeID uint, partnerID uint, companyID uint) ([]Appointment, int64, error) {
	return GetPaginatedAppointments(offset, limit, search, state, employeeID, partnerID, companyID)
}

func GetAppointmentByIDService(id uint) (*Appointment, error) {
	return GetAppointmentByID(id)
}

func UpdateAppointmentService(data *Appointment) error {
	return UpdateAppointment(data)
}

func DeleteAppointmentService(id uint) error {
	return DeleteAppointment(id)
}
