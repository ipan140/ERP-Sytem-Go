package ecommerce

import (
	"ERP-System/config"
	"time"
)

type Cart struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SessionID  string    `gorm:"type:varchar(255)" json:"session_id"` // For anonymous users
	CustomerID *uint     `json:"customer_id"`                         // Optional (if logged in)
	State      string    `gorm:"type:varchar(20);default:'active'" json:"state"` // active, abandoned, checkout, paid
	CreatedAt  time.Time `json:"created_at"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"` // Inventory Product
	Quantity  float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	PriceUnit float64 `gorm:"type:numeric(15,2);default:0" json:"price_unit"`
	SubTotal  float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Cart{}, &CartItem{})
}
