package core

import (
	"ERP-System/config"
	"time"
)

type AuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    string    `gorm:"type:varchar(50);index" json:"user_id"`
	Action    string    `gorm:"type:varchar(50)" json:"action"`
	Module    string    `gorm:"type:varchar(100);index" json:"module"`
	RecordID  string    `gorm:"type:varchar(50);index" json:"record_id"`
	OldData   string    `gorm:"type:text" json:"old_data"`
	NewData   string    `gorm:"type:text" json:"new_data"`
	CreatedAt time.Time `json:"created_at"`
}

func (AuditLog) TableName() string {
	return "core.audit_logs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &AuditLog{})
}
