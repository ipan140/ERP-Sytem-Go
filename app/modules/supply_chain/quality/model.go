package quality

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/supply_chain/manufacturing"
	"ERP-System/config"
	"time"
)

type QualityPoint struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	Name      string             `gorm:"type:varchar(255);not null" json:"name"` // e.g. Uji Dimensi & Presisi
	ProductID uint               `json:"product_id"`
	Product   *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	TestType  string             `gorm:"type:varchar(50);default:'passfail'" json:"test_type"` // passfail, measure
	Norm      float64            `gorm:"type:numeric(15,2);default:0" json:"norm"`             // Nilai target standar
	Tolerance float64            `gorm:"type:numeric(15,2);default:0" json:"tolerance"`        // Batas toleransi +/-
	CreatedAt time.Time          `json:"created_at"`
}

type QualityCheck struct {
	ID           uint                          `gorm:"primaryKey" json:"id"`
	Name         string                        `gorm:"type:varchar(100);not null" json:"name"` // QC/2026/0001
	PointID      *uint                         `json:"point_id"`
	Point        *QualityPoint                 `gorm:"foreignKey:PointID" json:"point,omitempty"` // Odoo relation mapped
	ProductID    uint                          `json:"product_id"`
	Product      *inventory.Product            `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	PickingID    *uint                         `json:"picking_id"`                                    // Jika QC saat barang datang
	Picking      *inventory.StockPicking       `gorm:"foreignKey:PickingID" json:"picking,omitempty"`  // Odoo relation mapped
	ProductionID *uint                         `json:"production_id"`                                 // Jika QC saat diproduksi
	Production   *manufacturing.MrpProduction  `gorm:"foreignKey:ProductionID" json:"production,omitempty"` // Odoo relation mapped
	Result       string                        `gorm:"type:varchar(50);default:'pending'" json:"result"` // pending, pass, fail
	MeasureValue float64                       `gorm:"type:numeric(15,2);default:0" json:"measure_value"`
	Notes        string                        `gorm:"type:text" json:"notes"`
	InspectorID  *uint                         `json:"inspector_id"`
	InspectedAt  *time.Time                    `json:"inspected_at"`
	CreatedAt    time.Time                     `json:"created_at"`
}

// Enterprise Quality DTOs (Fase 4)
type QualitySummary struct {
	TotalChecksCount   int64 `json:"total_checks_count"`
	PendingChecksCount int64 `json:"pending_checks_count"`
	PassedChecksCount  int64 `json:"passed_checks_count"`
	FailedChecksCount  int64 `json:"failed_checks_count"`
	TotalPointsCount   int64 `json:"total_points_count"`
}

type CreateQualityCheckRequest struct {
	PointID      *uint   `json:"point_id"`
	ProductID    uint    `json:"product_id" validate:"required"`
	PickingID    *uint   `json:"picking_id"`
	ProductionID *uint   `json:"production_id"`
	MeasureValue float64 `json:"measure_value"`
	Notes        string  `json:"notes"`
}

type ProcessQCRequest struct {
	Result       string  `json:"result" validate:"required"` // pass, fail
	MeasureValue float64 `json:"measure_value"`
	Notes        string  `json:"notes"`
}

type CreateQualityPointRequest struct {
	Name      string  `json:"name" validate:"required"`
	ProductID uint    `json:"product_id" validate:"required"`
	TestType  string  `json:"test_type"`
	Norm      float64 `json:"norm"`
	Tolerance float64 `json:"tolerance"`
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
