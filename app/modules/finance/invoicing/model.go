package invoicing

import (

	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type PaymentTerm struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Days int    `json:"days"`
}

type Invoice struct {
	ID             uint          `gorm:"primaryKey" json:"id"`
	Name           string        `gorm:"type:varchar(100);not null" json:"name"`
	PartnerID      uint          `json:"partner_id"`
	Partner        *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"`
	PaymentTermID  *uint         `json:"payment_term_id"`
	PaymentTerm    *PaymentTerm  `gorm:"foreignKey:PaymentTermID" json:"payment_term,omitempty"`
	IncotermID     *uint         `json:"incoterm_id"`

	InvoiceDate    time.Time     `json:"invoice_date"`
	DueDate        time.Time     `json:"due_date"`
	State          string        `gorm:"type:varchar(20);default:'draft'" json:"state"`
	FollowUpLevel  int           `gorm:"default:0" json:"follow_up_level"`
	AmountUntaxed  float64       `gorm:"type:numeric(15,2);default:0" json:"amount_untaxed"`
	AmountTax      float64       `gorm:"type:numeric(15,2);default:0" json:"amount_tax"`
	AmountTotal    float64       `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	ResidualAmount float64       `gorm:"type:numeric(15,2);default:0" json:"residual_amount"`
	CreatedAt      time.Time     `json:"created_at"`
}

type InvoiceLine struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	InvoiceID   uint     `json:"invoice_id"`
	Invoice     *Invoice `gorm:"foreignKey:InvoiceID" json:"invoice,omitempty"`
	Description string   `gorm:"type:varchar(255)" json:"description"`
	Quantity    float64  `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	UnitPrice   float64  `gorm:"type:numeric(15,2);default:0" json:"unit_price"`
	TaxID       *uint    `json:"tax_id"`
	Tax         *Tax     `gorm:"foreignKey:TaxID" json:"tax,omitempty"`
	SubTotal    float64  `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

type Tax struct {
	ID   uint    `gorm:"primaryKey" json:"id"`
	Name string  `gorm:"type:varchar(100)" json:"name"`
	Rate float64 `gorm:"type:numeric(5,2)" json:"rate"`
}

type PaymentTermLine struct {
	ID            uint         `gorm:"primaryKey" json:"id"`
	PaymentTermID uint         `json:"payment_term_id"`
	PaymentTerm   *PaymentTerm `gorm:"foreignKey:PaymentTermID" json:"payment_term,omitempty"`
	ValueType     string       `gorm:"type:varchar(20);default:'percent'" json:"value_type"`
	ValueAmount   float64      `gorm:"type:numeric(15,2)" json:"value_amount"`
	Days          int          `json:"days"`
}

type TaxRepartitionLine struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	TaxID           uint    `json:"tax_id"`
	Tax             *Tax    `gorm:"foreignKey:TaxID" json:"tax,omitempty"`
	RepartitionType string  `gorm:"type:varchar(20);default:'tax'" json:"repartition_type"`
	FactorPercent   float64 `gorm:"type:numeric(5,2);default:100" json:"factor_percent"`
	AccountID       *uint   `json:"account_id"`

}

func (PaymentTerm) TableName() string {
	return "finance.payment_terms"
}
func (Invoice) TableName() string {
	return "finance.invoices"
}
func (InvoiceLine) TableName() string {
	return "finance.invoice_lines"
}
func (Tax) TableName() string {
	return "finance.taxes"
}
func (PaymentTermLine) TableName() string {
	return "finance.payment_term_lines"
}
func (TaxRepartitionLine) TableName() string {
	return "finance.tax_repartition_lines"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PaymentTerm{}, &Invoice{}, &InvoiceLine{}, &Tax{}, &PaymentTermLine{}, &TaxRepartitionLine{})
}
