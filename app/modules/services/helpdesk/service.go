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

func CreateHelpdeskSLAService(data *HelpdeskSLA) error        { return CreateHelpdeskSLA(data) }
func GetAllHelpdeskSLAService() ([]HelpdeskSLA, error)        { return GetAllHelpdeskSLA() }
func GetHelpdeskSLAByIDService(id uint) (*HelpdeskSLA, error) { return GetHelpdeskSLAByID(id) }
func UpdateHelpdeskSLAService(data *HelpdeskSLA) error        { return UpdateHelpdeskSLA(data) }
func DeleteHelpdeskSLAService(id uint) error                  { return DeleteHelpdeskSLA(id) }

func CreateHelpdeskCannedResponseService(data *HelpdeskCannedResponse) error {
	return CreateHelpdeskCannedResponse(data)
}
func GetAllHelpdeskCannedResponseService() ([]HelpdeskCannedResponse, error) {
	return GetAllHelpdeskCannedResponse()
}
func GetHelpdeskCannedResponseByIDService(id uint) (*HelpdeskCannedResponse, error) {
	return GetHelpdeskCannedResponseByID(id)
}
func UpdateHelpdeskCannedResponseService(data *HelpdeskCannedResponse) error {
	return UpdateHelpdeskCannedResponse(data)
}
func DeleteHelpdeskCannedResponseService(id uint) error { return DeleteHelpdeskCannedResponse(id) }
