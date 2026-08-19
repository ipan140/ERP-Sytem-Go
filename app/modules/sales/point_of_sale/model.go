package point_of_sale

import (
	"ERP-System/config"
	"time"
)

type PosConfig struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Main Shop
	HasCredit bool      `gorm:"default:false" json:"has_credit"`
	CreatedAt time.Time `json:"created_at"`
}

type PosSession struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	ConfigID  uint      `json:"config_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	State     string    `gorm:"type:varchar(50);default:'opened'" json:"state"` // opened, closed
	CreatedAt time.Time `json:"created_at"`
}

type PosOrder struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // Order Ref
	SessionID uint      `json:"session_id"`
	PartnerID *uint     `json:"partner_id"`
	Total     float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State     string    `gorm:"type:varchar(50);default:'paid'" json:"state"` // draft, paid, done, invoiced
	CreatedAt time.Time `json:"created_at"`
}

type PosOrderLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Qty       float64 `gorm:"type:numeric(15,2);default:1" json:"qty"`
	PriceUnit float64 `gorm:"type:numeric(15,2);default:0" json:"price_unit"`
	SubTotal  float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

type PosPayment struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	Method    string  `gorm:"type:varchar(50);not null" json:"method"` // Cash, Bank, QRIS
	Amount    float64 `gorm:"type:numeric(15,2);not null" json:"amount"`
}

type LoyaltyProgram struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Buy 2 Get 1, Points
	Type      string `gorm:"type:varchar(50);not null" json:"type"`  // promotion, coupon, loyalty
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PosConfig{}, &PosSession{}, &PosOrder{}, &PosOrderLine{}, &PosPayment{}, &LoyaltyProgram{})
}
