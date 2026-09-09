package manufacturing

import (
	"ERP-System/app/modules/supply_chain/inventory"
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
	ID        uint               `gorm:"primaryKey" json:"id"`
	ProductID uint               `json:"product_id"`
	Product   *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Code      string             `gorm:"type:varchar(100)" json:"code"`                 // Reference
	Type      string             `gorm:"type:varchar(50);default:'normal'" json:"type"` // normal, phantom
	Quantity  float64            `gorm:"type:numeric(15,2);default:1" json:"quantity"`  // Product qty produced
	BomLines  []MrpBomLine       `gorm:"foreignKey:BomID;-:migration" json:"bom_lines,omitempty"`
}

type MrpBomLine struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	BomID     uint               `json:"bom_id"`
	Bom       *MrpBom            `gorm:"foreignKey:BomID" json:"bom,omitempty"` // Odoo relation mapped
	ProductID uint               `json:"product_id"`
	Product   *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Quantity  float64            `gorm:"type:numeric(15,2);not null;default:1" json:"quantity"` // Component qty consumed
}

type MrpBomByproduct struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	BomID     uint               `json:"bom_id"`
	Bom       *MrpBom            `gorm:"foreignKey:BomID" json:"bom,omitempty"` // Odoo relation mapped
	ProductID uint               `json:"product_id"` // Scrap/byproduct
	Product   *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Quantity  float64            `gorm:"type:numeric(15,2);default:1" json:"quantity"`
}

type MrpProduction struct {
	ID           uint                      `gorm:"primaryKey" json:"id"`
	Name         string                    `gorm:"type:varchar(100);not null" json:"name"` // WH/MO/0001
	ProductID    uint                      `json:"product_id"`
	Product      *inventory.Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	ProductQty   float64                   `gorm:"type:numeric(15,2);not null;default:1" json:"product_qty"`
	BomID        *uint                     `json:"bom_id"`
	Bom          *MrpBom                   `gorm:"foreignKey:BomID" json:"bom,omitempty"` // Odoo relation mapped
	WarehouseID  *uint                     `json:"warehouse_id"`
	Warehouse    *inventory.StockWarehouse `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	State        string                    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, confirmed, progress, to_close, done, cancel
	DatePlanned  *time.Time                `json:"date_planned"`
	DateStart    *time.Time                `json:"date_start"`
	DateFinished *time.Time                `json:"date_finished"`
	Notes        string                    `gorm:"type:text" json:"notes"`
	CompanyID    *uint                     `json:"company_id"`
	CreatedAt    time.Time                 `json:"created_at"`
}

type MrpWorkorder struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"` // Operation name (e.g. Cutting)
	ProductionID uint           `json:"production_id"`
	Production   *MrpProduction `gorm:"foreignKey:ProductionID" json:"production,omitempty"` // Odoo relation mapped
	WorkcenterID uint           `json:"workcenter_id"`
	Workcenter   *MrpWorkcenter `gorm:"foreignKey:WorkcenterID" json:"workcenter,omitempty"` // Odoo relation mapped
	State        string         `gorm:"type:varchar(50);default:'pending'" json:"state"` // pending, ready, progress, done, cancel
	Duration     float64        `gorm:"type:numeric(15,2);default:0" json:"duration"`    // Actual minutes spent
	CreatedAt    time.Time      `json:"created_at"`
}

type MrpUnbuild struct {
	ID           uint               `gorm:"primaryKey" json:"id"`
	Name         string             `gorm:"type:varchar(100);not null" json:"name"` // UB/0001
	ProductID    uint               `json:"product_id"`
	Product      *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	ProductQty   float64            `gorm:"type:numeric(15,2);not null;default:1" json:"product_qty"`
	BomID        *uint              `json:"bom_id"`
	Bom          *MrpBom            `gorm:"foreignKey:BomID" json:"bom,omitempty"` // Odoo relation mapped
	ProductionID *uint              `json:"production_id"`                                 // Jika berasal dari MO
	Production   *MrpProduction     `gorm:"foreignKey:ProductionID" json:"production,omitempty"` // Odoo relation mapped
	State        string             `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt    time.Time          `json:"created_at"`
}

// Enterprise SCM MRP DTOs (Fase 3)
type MrpSummary struct {
	TotalMOCount     int64 `json:"total_mo_count"`
	MOInProgress     int64 `json:"mo_in_progress"`
	MODoneCount      int64 `json:"mo_done_count"`
	TotalWorkcenters int64 `json:"total_workcenters"`
	TotalBoms        int64 `json:"total_boms"`
}

type CreateMORequest struct {
	ProductID   uint    `json:"product_id" validate:"required"`
	BomID       *uint   `json:"bom_id"`
	ProductQty  float64 `json:"product_qty" validate:"required,gt=0"`
	WarehouseID *uint   `json:"warehouse_id"`
	DatePlanned *string `json:"date_planned"`
	Notes       string  `json:"notes"`
}

type CreateBomLineInput struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity" validate:"required,gt=0"`
}

type CreateBomRequest struct {
	ProductID uint                 `json:"product_id" validate:"required"`
	Code      string               `json:"code"`
	Type      string               `json:"type"`
	Quantity  float64              `json:"quantity" validate:"required,gt=0"`
	Lines     []CreateBomLineInput `json:"lines" validate:"required,dive"`
}

func (MrpWorkcenter) TableName() string {
	return "supply_chain.mrp_workcenters"
}

func (MrpBom) TableName() string {
	return "supply_chain.mrp_boms"
}

func (MrpBomLine) TableName() string {
	return "supply_chain.mrp_bom_lines"
}

func (MrpBomByproduct) TableName() string {
	return "supply_chain.mrp_bom_byproducts"
}

func (MrpProduction) TableName() string {
	return "supply_chain.mrp_productions"
}

func (MrpWorkorder) TableName() string {
	return "supply_chain.mrp_workorders"
}

func (MrpUnbuild) TableName() string {
	return "supply_chain.mrp_unbuilds"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MrpWorkcenter{}, &MrpBom{}, &MrpBomLine{}, &MrpBomByproduct{}, &MrpProduction{}, &MrpWorkorder{}, &MrpUnbuild{})
}
