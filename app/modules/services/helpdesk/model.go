package helpdesk

import (
	"ERP-System/config"
	"time"
)

type Ticket struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`         // Issue Title
	CustomerID       uint      `json:"customer_id"`                                    // Partner ID
	Priority         string    `gorm:"type:varchar(20);default:'low'" json:"priority"` // low, medium, high
	AssigneeID       *uint     `json:"assignee_id"`                                    // Employee ID
	State            string    `gorm:"type:varchar(20);default:'new'" json:"state"`    // new, in_progress, solved, closed
	IssueDescription string    `gorm:"type:text" json:"issue_description"`
	CreatedAt        time.Time `json:"created_at"`
}

// 4. Helpdesk SLA (Robot Pengawas Waktu)
type HelpdeskSLA struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	Name          string  `gorm:"type:varchar(100);not null" json:"name"`                // e.g., Balas VIP < 2 Jam
	PriorityLevel string  `gorm:"type:varchar(20);default:'high'" json:"priority_level"` // Berlaku untuk prioritas apa
	TargetHours   float64 `gorm:"type:numeric(5,2)" json:"target_hours"`
}

// 10. Jawaban Instan & Otomatisasi Bantuan (Canned Responses)
type HelpdeskCannedResponse struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Keyword  string `gorm:"type:varchar(100);unique" json:"keyword"` // Kata Kunci, e.g., 'mati_lampu'
	Response string `gorm:"type:text" json:"response"`               // Template Jawaban Otomatis
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Ticket{}, &HelpdeskSLA{}, &HelpdeskCannedResponse{})
}


// ---- Auto-Generated TableName methods ----
func (Ticket) TableName() string {
	return "services.tickets"
}

func (HelpdeskSLA) TableName() string {
	return "services.helpdesk_s_l_as"
}

func (HelpdeskCannedResponse) TableName() string {
	return "services.helpdesk_canned_responses"
}