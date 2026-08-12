package invoicing

import (
	"ERP-System/config"
	"time"
)

type PaymentTerm struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Days int    `json:"days"`
}

type Invoice struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(100);not null" json:"name"`
	PartnerID      uint      `json:"partner_id"`
	PaymentTermID  *uint     `json:"payment_term_id"`
	InvoiceDate    time.Time `json:"invoice_date"`
	DueDate        time.Time `json:"due_date"`
	State          string    `gorm:"type:varchar(20);default:'draft'" json:"state"`
	FollowUpLevel  int       `gorm:"default:0" json:"follow_up_level"`
	AmountUntaxed  float64   `gorm:"type:numeric(15,2);default:0" json:"amount_untaxed"`
	AmountTax      float64   `gorm:"type:numeric(15,2);default:0" json:"amount_tax"`
	AmountTotal    float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	ResidualAmount float64   `gorm:"type:numeric(15,2);default:0" json:"residual_amount"`
	CreatedAt      time.Time `json:"created_at"`
}

type InvoiceLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	InvoiceID   uint    `json:"invoice_id"`
	Description string  `gorm:"type:varchar(255)" json:"description"`
	Quantity    float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	UnitPrice   float64 `gorm:"type:numeric(15,2);default:0" json:"unit_price"`
	TaxID       *uint   `json:"tax_id"`
	SubTotal    float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

type Tax struct {
	ID   uint    `gorm:"primaryKey" json:"id"`
	Name string  `gorm:"type:varchar(100)" json:"name"`
	Rate float64 `gorm:"type:numeric(5,2)" json:"rate"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PaymentTerm{}, &Invoice{}, &InvoiceLine{}, &Tax{})
}
