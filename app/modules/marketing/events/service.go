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

func CreateEventTicketService(data *EventTicket) error        { return CreateEventTicket(data) }
func GetAllEventTicketService() ([]EventTicket, error)        { return GetAllEventTicket() }
func GetEventTicketByIDService(id uint) (*EventTicket, error) { return GetEventTicketByID(id) }
func UpdateEventTicketService(data *EventTicket) error        { return UpdateEventTicket(data) }
func DeleteEventTicketService(id uint) error                  { return DeleteEventTicket(id) }
