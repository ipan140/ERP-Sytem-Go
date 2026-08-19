package quality

import (
	"ERP-System/config"
	"time"
)

type QualityPoint struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Check Temperature
	ProductID uint      `json:"product_id"`
	TestType  string    `gorm:"type:varchar(50);default:'passfail'" json:"test_type"` // passfail, measure
	Tolerance float64   `gorm:"type:numeric(15,2);default:0" json:"tolerance"` // For measure tests
	CreatedAt time.Time `json:"created_at"`
}

type QualityCheck struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // QC/001
	PointID      uint      `json:"point_id"`
	ProductID    uint      `json:"product_id"`
	PickingID    *uint     `json:"picking_id"`    // Jika QC saat barang datang
	ProductionID *uint     `json:"production_id"` // Jika QC saat diproduksi
	Result       string    `gorm:"type:varchar(50);default:'pending'" json:"result"` // pending, pass, fail
	MeasureValue float64   `gorm:"type:numeric(15,2);default:0" json:"measure_value"`
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &QualityPoint{}, &QualityCheck{})
}
