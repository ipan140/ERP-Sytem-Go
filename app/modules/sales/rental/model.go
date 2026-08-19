package rental

import (
	"ERP-System/config"
	"time"
)

type RentalOrder struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // RO/2026/001
	PartnerID  uint      `json:"partner_id"`
	PickupDate time.Time `json:"pickup_date"`
	ReturnDate time.Time `json:"return_date"`
	Total      float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, reserved, pickedup, returned
	CreatedAt  time.Time `json:"created_at"`
}

type RentalOrderLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Qty       float64 `gorm:"type:numeric(15,2);default:1" json:"qty"`
	PriceUnit float64 `gorm:"type:numeric(15,2);default:0" json:"price_unit"`
	SubTotal  float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &RentalOrder{}, &RentalOrderLine{})
}
