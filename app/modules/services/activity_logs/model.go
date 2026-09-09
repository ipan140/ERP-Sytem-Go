package activity_logs

import (
	"ERP-System/config"
	"time"
)

type ActivityLog struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CompanyID  *uint     `json:"company_id"`
	EntityType string    `gorm:"type:varchar(50);not null" json:"entity_type"`
	EntityID   uint      `gorm:"not null" json:"entity_id"`
	Action     string    `gorm:"type:varchar(50);not null" json:"action"`
	UserID     *uint     `json:"user_id"`
	UserName   string    `gorm:"type:varchar(100)" json:"user_name"`
	OldValue   string    `gorm:"type:jsonb" json:"old_value,omitempty"`
	NewValue   string    `gorm:"type:jsonb" json:"new_value,omitempty"`
	Notes      string    `gorm:"type:text" json:"notes,omitempty"`
	IPAddress  string    `gorm:"type:varchar(50)" json:"ip_address,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ActivityLog{})
}

func (ActivityLog) TableName() string {
	return "services.activity_logs"
}
