package point_of_sale

import (
	"ERP-System/app/modules/sales/sales_core"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"time"
)

type PosConfig struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` 
	HasCredit bool      `gorm:"default:false" json:"has_credit"`
	CreatedAt time.Time `json:"created_at"`
}

type PosSession struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ConfigID  uint      `json:"config_id"`
	Config *PosConfig `gorm:"foreignKey:ConfigID" json:"config,omitempty"` 
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	State     string    `gorm:"type:varchar(50);default:'opened'" json:"state"` 
	CreatedAt time.Time `json:"created_at"`
}

type PosOrder struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` 
	SessionID uint      `json:"session_id"`
	Session *PosSession `gorm:"foreignKey:SessionID" json:"session,omitempty"` 
	PartnerID *uint     `json:"partner_id"`
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` 
	Total     float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State     string    `gorm:"type:varchar(50);default:'paid'" json:"state"` 
	CreatedAt time.Time `json:"created_at"`
}

type PosOrderLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	Order *sales_core.SaleOrder `gorm:"foreignKey:OrderID" json:"order,omitempty"` 
	ProductID uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` 
	Qty       float64 `gorm:"type:numeric(15,2);default:1" json:"qty"`
	PriceUnit float64 `gorm:"type:numeric(15,2);default:0" json:"price_unit"`
	SubTotal  float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

type PosPayment struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	OrderID uint    `json:"order_id"`
	Order *sales_core.SaleOrder `gorm:"foreignKey:OrderID" json:"order,omitempty"` 
	Method  string  `gorm:"type:varchar(50);not null" json:"method"` 
	Amount  float64 `gorm:"type:numeric(15,2);not null" json:"amount"`
}

type LoyaltyProgram struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	Name      string             `gorm:"type:varchar(255);not null" json:"name"` 
	Type      string             `gorm:"type:varchar(50);not null" json:"type"`  
	RewardID  uint               `json:"reward_id"`
	Reward    *inventory.Product `gorm:"foreignKey:RewardID" json:"reward,omitempty"`
}

func (PosConfig) TableName() string {
	return "sales.pos_configs"
}

func (PosSession) TableName() string {
	return "sales.pos_sessions"
}

func (PosOrder) TableName() string {
	return "sales.pos_orders"
}

func (PosOrderLine) TableName() string {
	return "sales.pos_order_lines"
}

func (PosPayment) TableName() string {
	return "sales.pos_payments"
}

func (LoyaltyProgram) TableName() string {
	return "sales.loyalty_programs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PosConfig{}, &PosSession{}, &PosOrder{}, &PosOrderLine{}, &PosPayment{}, &LoyaltyProgram{})
}
