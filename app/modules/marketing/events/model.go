package events

import (
	"time"

	"ERP-System/app/modules/core/base"
	"ERP-System/config"
)

// 5. Event Ticketing & Barcode
type EventTicket struct {
	ID            uint          `gorm:"primaryKey" json:"id"`
	EventID       uint          `json:"event_id"`
	Event         *Event        `gorm:"foreignKey:EventID"` // Auto-added relation
	CustomerID    *uint         `json:"customer_id"`
	Customer      *base.Partner `gorm:"foreignKey:CustomerID" json:"customer,omitempty"` // Cross-module relation
	AttendeeName  string        `gorm:"type:varchar(255)" json:"attendee_name"`
	AttendeeEmail string        `gorm:"type:varchar(255)" json:"attendee_email"`
	Barcode       string        `gorm:"type:varchar(100);unique" json:"barcode"`         // Di-scan saat acara
	IsScanned     bool          `gorm:"default:false" json:"is_scanned"`
}

func (EventTicket) TableName() string {
	return "marketing.event_tickets"
}

func (Event) TableName() string {
	return "marketing.events"
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &Event{}, &EventTicket{}) }

type Event struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EventName   string    `gorm:"type:varchar(255);default:'General Event'" json:"event_name"`
	Location    string    `gorm:"type:varchar(255)" json:"location"`
	MaxCapacity int       `gorm:"default:100" json:"max_capacity"`
	EventDate   time.Time `json:"event_date"`
	Status      string    `gorm:"type:varchar(50);default:'Open'" json:"status"` // Open, Ongoing, Completed
}
