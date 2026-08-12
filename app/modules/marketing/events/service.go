package events

func CreateEventService(data *Event) error {
	return CreateEvent(data)
}

func GetAllEventService() ([]Event, error) {
	return GetAllEvent()
}

func GetEventByIDService(id uint) (*Event, error) {
	return GetEventByID(id)
}

func UpdateEventService(data *Event) error {
	return UpdateEvent(data)
}

func DeleteEventService(id uint) error {
	return DeleteEvent(id)
}
