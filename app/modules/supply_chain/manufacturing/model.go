package manufacturing

import (
	"ERP-System/config"
	"time"
)

type MrpWorkcenter struct {
	ID             uint    `gorm:"primaryKey" json:"id"`
	Name           string  `gorm:"type:varchar(255);not null" json:"name"` // Assembly Line 1
	Code           string  `gorm:"type:varchar(50)" json:"code"`
	TimeEfficiency float64 `gorm:"type:numeric(5,2);default:100" json:"time_efficiency"`
	Capacity       float64 `gorm:"type:numeric(15,2);default:1" json:"capacity"`
	CostsHour      float64 `gorm:"type:numeric(15,2);default:0" json:"costs_hour"`
}

type MrpBom struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ProductID uint    `json:"product_id"`
	Code      string  `gorm:"type:varchar(100)" json:"code"`                 // Reference
	Type      string  `gorm:"type:varchar(50);default:'normal'" json:"type"` // normal, phantom
	Quantity  float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`  // Product qty produced
}

type MrpBomLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	BomID     uint    `json:"bom_id"`
	ProductID uint    `json:"product_id"`
	Quantity  float64 `gorm:"type:numeric(15,2);not null;default:1" json:"quantity"` // Component qty consumed
}

type MrpBomByproduct struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	BomID     uint    `json:"bom_id"`
	ProductID uint    `json:"product_id"` // Scrap/byproduct
	Quantity  float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
}

type MrpProduction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"` // WH/MO/0001
	ProductID   uint      `json:"product_id"`
	ProductQty  float64   `gorm:"type:numeric(15,2);not null;default:1" json:"product_qty"`
	BomID       uint      `json:"bom_id"`
	State       string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, confirmed, progress, to_close, done, cancel
	DatePlanned time.Time `json:"date_planned"`
	CreatedAt   time.Time `json:"created_at"`
}

type MrpWorkorder struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"` // Operation name (e.g. Cutting)
	ProductionID uint      `json:"production_id"`
	WorkcenterID uint      `json:"workcenter_id"`
	State        string    `gorm:"type:varchar(50);default:'pending'" json:"state"` // pending, ready, progress, done, cancel
	Duration     float64   `gorm:"type:numeric(15,2);default:0" json:"duration"`    // Actual minutes spent
	CreatedAt    time.Time `json:"created_at"`
}

type MrpUnbuild struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // UB/0001
	ProductID    uint      `json:"product_id"`
	ProductQty   float64   `gorm:"type:numeric(15,2);not null;default:1" json:"product_qty"`
	BomID        *uint     `json:"bom_id"`
	ProductionID *uint     `json:"production_id"`                                 // Jika berasal dari MO
	State        string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MrpWorkcenter{}, &MrpBom{}, &MrpBomLine{}, &MrpBomByproduct{}, &MrpProduction{}, &MrpWorkorder{}, &MrpUnbuild{})
}
