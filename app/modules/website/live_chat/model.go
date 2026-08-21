package live_chat

import "ERP-System/config"

// 3. Helpdesk Live Chat
type ChatSession struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	VisitorIP  string `gorm:"type:varchar(50)" json:"visitor_ip"`
	OperatorID *uint  `json:"operator_id"` // CS yang membalas
	TicketID   *uint  `json:"ticket_id"`   // Nyambung ke Helpdesk jika masalah tidak selesai
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &ChatSession{}) }
