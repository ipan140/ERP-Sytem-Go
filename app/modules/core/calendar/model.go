package calendar

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"gorm.io/gorm"
)

type CalendarEvent struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Description string     `gorm:"type:text" json:"description"`
	Location    string     `gorm:"type:varchar(255)" json:"location"`
	MeetingURL  string     `gorm:"type:varchar(255)" json:"meeting_url"`
	JoinToken   string     `gorm:"type:varchar(100);uniqueIndex" json:"join_token"`
	EventType   string     `gorm:"type:varchar(100)" json:"event_type"`
	Start       time.Time  `gorm:"not null" json:"start"`
	End         time.Time  `json:"end"`
	AllDay      bool       `gorm:"default:false" json:"allDay"`
	Color       string     `gorm:"type:varchar(20)" json:"color"`

	ResModel    string     `gorm:"type:varchar(100);index" json:"res_model"`
	ResID       uint       `gorm:"index" json:"res_id"`

	Visibility  string     `gorm:"type:varchar(20);default:'private'" json:"visibility"`

	UserID      uint       `json:"user_id"`
	User        *auth.User `gorm:"foreignKey:UserID" json:"user,omitempty"`

	Attendees   []*auth.User `gorm:"many2many:calendar_event_attendees;" json:"attendees,omitempty"`

	AttendeeIDs []uint     `gorm:"-" json:"attendee_ids,omitempty"`

	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (e *CalendarEvent) BeforeCreate(tx *gorm.DB) (err error) {
	if e.JoinToken == "" {
		bytes := make([]byte, 16)
		rand.Read(bytes)
		e.JoinToken = hex.EncodeToString(bytes)
	}
	return
}

type CalendarCategory struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Module      string    `gorm:"type:varchar(50);index;not null" json:"module"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Color       string    `gorm:"type:varchar(50);not null" json:"color"`
	Icon        string    `gorm:"type:varchar(50)" json:"icon"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (CalendarCategory) TableName() string {
	return "setting.calendar_categories"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &CalendarEvent{}, &CalendarCategory{})
}
