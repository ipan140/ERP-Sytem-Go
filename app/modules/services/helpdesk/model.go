package helpdesk

import (
	"ERP-System/config"
	"time"
)

type Ticket struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"` // Issue Title
	CustomerID       uint      `json:"customer_id"`                            // Partner ID
	Priority         string    `gorm:"type:varchar(20);default:'low'" json:"priority"` // low, medium, high
	AssigneeID       *uint     `json:"assignee_id"`                                    // Employee ID
	State            string    `gorm:"type:varchar(20);default:'new'" json:"state"`    // new, in_progress, solved, closed
	IssueDescription string    `gorm:"type:text" json:"issue_description"`
	CreatedAt        time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Ticket{})
}
