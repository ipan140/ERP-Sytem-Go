package inventory

import (
	"ERP-System/config"
	"time"
)

type Product struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	SKU       string    `gorm:"type:varchar(100);unique" json:"sku"`
	Price     float64   `gorm:"type:numeric(15,2);default:0" json:"price"` // Sale price
	Cost      float64   `gorm:"type:numeric(15,2);default:0" json:"cost"`  // Purchase cost
	StockQty  float64   `gorm:"type:numeric(15,2);default:0" json:"stock_qty"`
	CreatedAt time.Time `json:"created_at"`
}

type StockMove struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ProductID    uint      `json:"product_id"`
	FromLocation string    `gorm:"type:varchar(255)" json:"from_location"` // E.g. Vendor, Warehouse A
	ToLocation   string    `gorm:"type:varchar(255)" json:"to_location"`   // E.g. Warehouse A, Customer
	Quantity     float64   `gorm:"type:numeric(15,2);not null" json:"quantity"`
	State        string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done, cancelled
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Product{}, &StockMove{})
}
