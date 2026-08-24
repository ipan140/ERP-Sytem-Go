package whatsapp

import (
	"ERP-System/config"
	"time"
)

type WaTemplate struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}


func (WaTemplate) TableName() string {
	return "setting.wa_templates"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &WaTemplate{})
}
