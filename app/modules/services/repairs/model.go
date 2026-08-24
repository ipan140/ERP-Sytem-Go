package repairs

import (
	"ERP-System/config"
	"time"
)

type RepairOrder struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &RepairOrder{})
}


// ---- Auto-Generated TableName methods ----
func (RepairOrder) TableName() string {
	return "services.repair_orders"
}