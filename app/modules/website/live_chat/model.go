package live_chat

import (
	"ERP-System/app/modules/services/helpdesk"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
)

// 3. Helpdesk Live Chat
type ChatSession struct {
	ID         uint                `gorm:"primaryKey" json:"id"`
	VisitorIP  string              `gorm:"type:varchar(50)" json:"visitor_ip"`
	OperatorID *uint               `json:"operator_id"`                                     // CS yang membalas
	Operator   *employees.Employee `gorm:"foreignKey:OperatorID" json:"operator,omitempty"` // Cross-module relation
	TicketID   *uint               `json:"ticket_id"`                                       // Nyambung ke Helpdesk jika masalah tidak selesai
	Ticket *helpdesk.Ticket `gorm:"foreignKey:TicketID" json:"ticket,omitempty"` // Odoo relation mapped
}

func (ChatSession) TableName() string {
	return "website_portal.chat_sessions"
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &ChatSession{}) }
