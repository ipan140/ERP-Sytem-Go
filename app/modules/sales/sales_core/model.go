package sales_core

import (
	"ERP-System/config"
	"time"
)

type SaleOrder struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. SO/2026/001
	PartnerID   uint      `json:"partner_id"`                             // Customer ID
	DateOrder   time.Time `json:"date_order"`
	State       string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, sent, sale, done, cancel
	AmountTotal float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	CreatedAt   time.Time `json:"created_at"`
}

type SaleOrderLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	OrderID     uint    `json:"order_id"`
	ProductID   uint    `json:"product_id"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Quantity    float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	UnitPrice   float64 `gorm:"type:numeric(15,2);default:0" json:"unit_price"`
	SubTotal    float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SaleOrder{}, &SaleOrderLine{})
}
