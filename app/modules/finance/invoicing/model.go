package invoicing

import (
	"ERP-System/config"
	"time"
)

type Invoice struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. INV/2026/001
	PartnerID    uint      `json:"partner_id"`                             // Customer or Vendor ID
	InvoiceDate  time.Time `json:"invoice_date"`
	DueDate      time.Time `json:"due_date"`
	State        string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, posted, paid
	AmountTotal  float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	CreatedAt    time.Time `json:"created_at"`
}

type InvoiceLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	InvoiceID   uint    `json:"invoice_id"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Quantity    float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	UnitPrice   float64 `gorm:"type:numeric(15,2);default:0" json:"unit_price"`
	SubTotal    float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

type Tax struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `gorm:"type:varchar(100)" json:"name"`
	Rate   float64 `gorm:"type:numeric(5,2)" json:"rate"` // e.g. 11.00 for 11%
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Invoice{}, &InvoiceLine{}, &Tax{})
}
