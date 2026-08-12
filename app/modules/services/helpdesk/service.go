package helpdesk

func CreateTicketService(data *Ticket) error {
	return CreateTicket(data)
}

func GetAllTicketService() ([]Ticket, error) {
	return GetAllTicket()
}

func GetTicketByIDService(id uint) (*Ticket, error) {
	return GetTicketByID(id)
}

func UpdateTicketService(data *Ticket) error {
	return UpdateTicket(data)
}

func DeleteTicketService(id uint) error {
	return DeleteTicket(id)
}
