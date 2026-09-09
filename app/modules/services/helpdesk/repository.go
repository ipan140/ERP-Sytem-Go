package helpdesk

import (
	"strconv"
	"time"

	"ERP-System/config"
	"gorm.io/gorm/clause"
)

// ComputeSlaDeadline calculates deadline based on priority or SLA configuration
func ComputeSlaDeadline(priority string) time.Time {
	var sla HelpdeskSLA
	targetHours := 24.0 // default low: 24h
	switch priority {
	case "urgent":
		targetHours = 2.0
	case "high":
		targetHours = 4.0
	case "medium":
		targetHours = 8.0
	}

	if err := config.DB.Where("priority_level = ?", priority).First(&sla).Error; err == nil && sla.TargetHours > 0 {
		targetHours = sla.TargetHours
	}

	return time.Now().Add(time.Duration(targetHours * float64(time.Hour)))
}

func CreateTicket(data *Ticket) error {
	if data.SlaDeadline == nil {
		dl := ComputeSlaDeadline(data.Priority)
		data.SlaDeadline = &dl
	}
	if data.SlaStatus == "" {
		data.SlaStatus = "ok"
	}
	if data.EscalationLevel == 0 {
		data.EscalationLevel = 1
	}
	return config.DB.Create(data).Error
}

func GetAllTicket() ([]Ticket, error) {
	var list []Ticket
	err := config.DB.Preload(clause.Associations).Order("id DESC").Find(&list).Error
	return list, err
}

func GetPaginatedTickets(offset int, limit int, search string, state string, priority string, assigneeID uint, customerID uint, companyID uint) ([]Ticket, int64, error) {
	var list []Ticket
	var total int64

	query := config.DB.Model(&Ticket{}).Preload(clause.Associations)

	if companyID > 0 {
		query = query.Where("company_id = ?", companyID)
	}
	if assigneeID > 0 {
		query = query.Where("assignee_id = ?", assigneeID)
	}
	if customerID > 0 {
		query = query.Where("customer_id = ?", customerID)
	}
	if state != "" && state != "all" {
		query = query.Where("state = ?", state)
	}
	if priority != "" && priority != "all" {
		query = query.Where("priority = ?", priority)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR issue_description ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetTicketByID(id uint) (*Ticket, error) {
	var data Ticket
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateTicket(data *Ticket) error {
	return config.DB.Save(data).Error
}

func DeleteTicket(id uint) error {
	return config.DB.Delete(&Ticket{}, id).Error
}

func CreateHelpdeskSLA(data *HelpdeskSLA) error { return config.DB.Create(data).Error }
func GetAllHelpdeskSLA() ([]HelpdeskSLA, error) {
	var list []HelpdeskSLA
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetHelpdeskSLAByID(id uint) (*HelpdeskSLA, error) {
	var data HelpdeskSLA
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateHelpdeskSLA(data *HelpdeskSLA) error { return config.DB.Save(data).Error }
func DeleteHelpdeskSLA(id uint) error           { return config.DB.Delete(&HelpdeskSLA{}, id).Error }

func CreateHelpdeskCannedResponse(data *HelpdeskCannedResponse) error {
	return config.DB.Create(data).Error
}
func GetAllHelpdeskCannedResponse() ([]HelpdeskCannedResponse, error) {
	var list []HelpdeskCannedResponse
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetHelpdeskCannedResponseByID(id uint) (*HelpdeskCannedResponse, error) {
	var data HelpdeskCannedResponse
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateHelpdeskCannedResponse(data *HelpdeskCannedResponse) error {
	return config.DB.Save(data).Error
}
func DeleteHelpdeskCannedResponse(id uint) error {
	return config.DB.Delete(&HelpdeskCannedResponse{}, id).Error
}

func EscalateTicket(id uint, level int, reason string) (*Ticket, error) {
	ticket, err := GetTicketByID(id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	ticket.EscalationLevel = level
	ticket.EscalatedAt = &now
	ticket.SlaStatus = "breached"
	if reason != "" {
		ticket.IssueDescription += "\n[ESKALASI TIER-" + strconv.Itoa(level) + "]: " + reason
	}

	err = config.DB.Save(ticket).Error
	return ticket, err
}

func ProcessAutoEscalationSLA() (int, int, error) {
	var activeTickets []Ticket
	err := config.DB.Where("state NOT IN ('solved', 'closed')").Find(&activeTickets).Error
	if err != nil {
		return 0, 0, err
	}

	now := time.Now()
	warningsCount := 0
	escalationsCount := 0

	for _, t := range activeTickets {
		if t.SlaDeadline == nil {
			dl := ComputeSlaDeadline(t.Priority)
			t.SlaDeadline = &dl
			config.DB.Model(&Ticket{}).Where("id = ?", t.ID).Update("sla_deadline", dl)
		}

		// Check if breached
		if now.After(*t.SlaDeadline) {
			if t.SlaStatus != "breached" || t.EscalationLevel < 2 {
				t.SlaStatus = "breached"
				t.EscalationLevel = 2
				t.EscalatedAt = &now
				config.DB.Model(&Ticket{}).Where("id = ?", t.ID).Updates(map[string]interface{}{
					"sla_status":       "breached",
					"escalation_level": 2,
					"escalated_at":     now,
				})
				escalationsCount++
			}
		} else {
			// Check if warning (< 20% remaining)
			totalDuration := t.SlaDeadline.Sub(t.CreatedAt)
			remaining := t.SlaDeadline.Sub(now)
			if totalDuration > 0 && float64(remaining)/float64(totalDuration) <= 0.20 {
				if t.SlaStatus != "warning" && t.SlaStatus != "breached" {
					config.DB.Model(&Ticket{}).Where("id = ?", t.ID).Update("sla_status", "warning")
					warningsCount++
				}
			}
		}
	}

	return warningsCount, escalationsCount, nil
}

