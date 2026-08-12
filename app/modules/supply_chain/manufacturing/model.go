package manufacturing

import (
	"ERP-System/config"
	"time"
)

type ManufacturingOrder struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // MO/2026/001
	ProductID    uint      `json:"product_id"`                             // Finished good ID
	QtyToProduce float64   `gorm:"type:numeric(15,2);default:1" json:"qty_to_produce"`
	State        string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, confirmed, progress, done, cancel
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ManufacturingOrder{})
}
