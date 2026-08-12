package events

import (
	"ERP-System/config"
	"time"
)

type Event struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Location  string    `gorm:"type:varchar(255)" json:"location"`
	State     string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, announced, running, done, cancelled
	CreatedAt time.Time `json:"created_at"`
}

type EventRegistration struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	EventID      uint      `json:"event_id"`
	AttendeeName string    `gorm:"type:varchar(255)" json:"attendee_name"`
	Email        string    `gorm:"type:varchar(255)" json:"email"`
	State        string    `gorm:"type:varchar(20);default:'open'" json:"state"` // open, done, cancel
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Event{}, &EventRegistration{})
}
