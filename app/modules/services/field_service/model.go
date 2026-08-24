package field_service

import (
	"ERP-System/config"
	"time"
)

type FieldServiceTask struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &FieldServiceTask{})
}


// ---- Auto-Generated TableName methods ----
func (FieldServiceTask) TableName() string {
	return "services.field_service_tasks"
}