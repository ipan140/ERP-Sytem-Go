package quality

import (
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"time"
)

type QualityPoint struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Check Temperature
	ProductID uint      `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	TestType  string    `gorm:"type:varchar(50);default:'passfail'" json:"test_type"` // passfail, measure
	Tolerance float64   `gorm:"type:numeric(15,2);default:0" json:"tolerance"`        // For measure tests
	CreatedAt time.Time `json:"created_at"`
}

type QualityCheck struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // QC/001
	PointID      uint      `json:"point_id"`
	Point *QualityPoint `gorm:"foreignKey:PointID" json:"point,omitempty"` // Odoo relation mapped
	ProductID    uint      `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	PickingID    *uint     `json:"picking_id"`                                       // Jika QC saat barang datang
	Picking *inventory.StockPicking `gorm:"foreignKey:PickingID" json:"picking,omitempty"` // Odoo relation mapped
	ProductionID *uint     `json:"production_id"`                                    // Jika QC saat diproduksi
	Production *manufacturing.MrpProduction `gorm:"foreignKey:ProductionID" json:"production,omitempty"` // Odoo relation mapped
	Result       string    `gorm:"type:varchar(50);default:'pending'" json:"result"` // pending, pass, fail
	MeasureValue float64   `gorm:"type:numeric(15,2);default:0" json:"measure_value"`
	CreatedAt    time.Time `json:"created_at"`
}


func (QualityPoint) TableName() string {
	return "supply_chain.quality_points"
}

func (QualityCheck) TableName() string {
	return "supply_chain.quality_checks"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &QualityPoint{}, &QualityCheck{})
}
