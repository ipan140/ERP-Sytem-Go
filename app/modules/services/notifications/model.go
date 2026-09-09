package notifications

import (
	"ERP-System/config"
	"time"
)

type NotificationLog struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	CompanyID     *uint     `json:"company_id"`
	Channel       string    `gorm:"type:varchar(20);default:'whatsapp'" json:"channel"` // 'whatsapp', 'email'
	Recipient     string    `gorm:"type:varchar(100);not null" json:"recipient"`
	RecipientName string    `gorm:"type:varchar(100)" json:"recipient_name"`
	EntityType    string    `gorm:"type:varchar(50)" json:"entity_type"`
	EntityID      *uint     `json:"entity_id"`
	Subject       string    `gorm:"type:varchar(255)" json:"subject"`
	Message       string    `gorm:"type:text;not null" json:"message"`
	Status        string    `gorm:"type:varchar(20);default:'sent'" json:"status"` // 'sent', 'queued', 'failed'
	ErrorMessage  string    `gorm:"type:text" json:"error_message,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &NotificationLog{})
}

func (NotificationLog) TableName() string {
	return "services.notification_logs"
}
