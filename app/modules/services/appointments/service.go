package appointments

func CreateAppointmentService(data *Appointment) error {
	return CreateAppointment(data)
}

func GetAllAppointmentService() ([]Appointment, error) {
	return GetAllAppointment()
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
