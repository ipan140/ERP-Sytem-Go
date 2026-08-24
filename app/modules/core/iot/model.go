package iot

import (
	"ERP-System/config"
	"time"
)

type IoTDevice struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &IoTDevice{})
}


// ---- Auto-Generated TableName methods ----
func (IoTDevice) TableName() string {
	return "setting.iot_devices"
}