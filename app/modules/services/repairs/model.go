package repairs

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"time"
)

type RepairOrder struct {
	ID             uint                `gorm:"primaryKey" json:"id"`
	ProductID      uint                `json:"product_id"`
	Product        *inventory.Product  `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	PartnerID      *uint               `json:"partner_id"`
	Partner        *base.Partner       `gorm:"foreignKey:PartnerID" json:"partner,omitempty"`
	TechnicianID   *uint               `json:"technician_id"`
	Technician     *employees.Employee `gorm:"foreignKey:TechnicianID" json:"technician,omitempty"`
	Name           string              `gorm:"type:varchar(255);not null" json:"name"`
	SerialNumber   string              `gorm:"type:varchar(100)" json:"serial_number"`
	WarrantyStatus string              `gorm:"type:varchar(50);default:'under_warranty'" json:"warranty_status"`
	Diagnosis      string              `gorm:"type:text" json:"diagnosis"`
	PartsCost      float64             `gorm:"type:numeric(15,2);default:0" json:"parts_cost"`
	LaborCost      float64             `gorm:"type:numeric(15,2);default:0" json:"labor_cost"`
	TotalCost      float64             `gorm:"type:numeric(15,2);default:0" json:"total_cost"`
	State          string              `gorm:"type:varchar(50);default:'confirmed'" json:"state"`
	CreatedAt      time.Time           `json:"created_at"`
	UpdatedAt      time.Time           `json:"updated_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &RepairOrder{})
}

// ---- Auto-Generated TableName methods ----
func (RepairOrder) TableName() string {
	return "services.repair_orders"
}