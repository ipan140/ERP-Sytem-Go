package helpdesk

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

func CreateTicketService(data *Ticket) error {
	err := CreateTicket(data)
	if err == nil && rabbitmq.Channel != nil {
		body, _ := json.Marshal(data)
		_ = rabbitmq.PublishDelayedEvent(rabbitmq.Channel, "services_sla_escalation", body, 14400000) // 4 jam delay
		log.Println("🚨 Event RabbitMQ: Timer SLA 4 Jam dimulai untuk Tiket Baru!")
	}
	return err
}

func GetAllTicketService() ([]Ticket, error) {
	return GetAllTicket()
}

func GetPaginatedTicketService(offset int, limit int, search string, state string, priority string, assigneeID uint, customerID uint, companyID uint) ([]Ticket, int64, error) {
	return GetPaginatedTickets(offset, limit, search, state, priority, assigneeID, customerID, companyID)
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

func EscalateTicketService(id uint, level int, reason string) (*Ticket, error) {
	return EscalateTicket(id, level, reason)
}

func ProcessAutoEscalationSLAService() (int, int, error) {
	return ProcessAutoEscalationSLA()
}
