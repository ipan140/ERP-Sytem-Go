package sales_core

import (
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/app/modules/core/base"
	"ERP-System/config"
	"time"
)

type Pricelist struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. VIP Customers
	Currency  string    `gorm:"type:varchar(10);default:'IDR'" json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type PricelistItem struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	PricelistID uint    `json:"pricelist_id"`
	Pricelist *Pricelist `gorm:"foreignKey:PricelistID"` // Auto-added relation
	ProductID   uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	FixedPrice  float64 `gorm:"type:numeric(15,2);default:0" json:"fixed_price"`
	MinQuantity float64 `gorm:"type:numeric(15,2);default:1" json:"min_quantity"`
}

type QuotationTemplate struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Basic Installation Package
}

type DeliveryMethod struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"type:varchar(100);not null" json:"name"` // e.g. JNE Regular
	FixedCost float64 `gorm:"type:numeric(15,2);default:0" json:"fixed_cost"`
}

type SaleOrder struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Name                string    `gorm:"type:varchar(100);not null" json:"name"` // SO/2026/001
	PartnerID           uint      `json:"partner_id"`
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` // Cross-module relation
	DateOrder           time.Time `json:"date_order"`
	State               string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, sent, sale, done, cancel
	AmountUntaxed       float64   `gorm:"type:numeric(15,2);default:0" json:"amount_untaxed"`
	AmountTax           float64   `gorm:"type:numeric(15,2);default:0" json:"amount_tax"`
	AmountTotal         float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	PricelistID         *uint     `json:"pricelist_id"`
	Pricelist *Pricelist `gorm:"foreignKey:PricelistID"` // Auto-added relation
	QuotationTemplateID *uint     `json:"quotation_template_id"`
	QuotationTemplate *QuotationTemplate `gorm:"foreignKey:QuotationTemplateID"` // Auto-added relation
	DeliveryMethodID    *uint     `json:"delivery_method_id"`
	DeliveryMethod *DeliveryMethod `gorm:"foreignKey:DeliveryMethodID"` // Auto-added relation
	InvoicingPolicy     string    `gorm:"type:varchar(50);default:'ordered'" json:"invoicing_policy"` // ordered, delivered
	IsPaymentLinkSent   bool      `gorm:"default:false" json:"is_payment_link_sent"`
	IsSigned            bool      `gorm:"default:false" json:"is_signed"`
	
	// Enterprise Quotation & Multi-Tier Pricing Fields
	CustomerName        string          `gorm:"type:varchar(255)" json:"customer_name"`
	CustomerEmail       string          `gorm:"type:varchar(100)" json:"customer_email"`
	PricelistName       string          `gorm:"type:varchar(100);default:'Standard Retail'" json:"pricelist_name"` // Standard, Grosir B2B, VIP Distributor
	MaxDiscount         float64         `gorm:"type:numeric(5,2);default:0" json:"max_discount"`
	NeedsApproval       bool            `gorm:"default:false" json:"needs_approval"`
	ApprovalStatus      string          `gorm:"type:varchar(50);default:'None'" json:"approval_status"` // None, Pending_ASM, Pending_Director, Approved, Rejected
	ApprovalTier        string          `gorm:"type:varchar(50);default:'Auto'" json:"approval_tier"`   // Auto, ASM (Area Sales Manager), Director (National Sales Director)
	ApprovedBy          string          `gorm:"type:varchar(100)" json:"approved_by"`

	// FASE 2: Multi-Branch & Multi-Company Partitioning
	BranchID            uint            `gorm:"default:1" json:"branch_id"`
	BranchName          string          `gorm:"type:varchar(100);default:'Head Office Jakarta'" json:"branch_name"` // Jakarta, Surabaya, Medan, Bandung, Bali
	CompanyID           uint            `gorm:"default:1" json:"company_id"`
	CompanyName         string          `gorm:"type:varchar(100);default:'PT Solusi Enterprise Indonesia'" json:"company_name"`

	// FASE 1: Payment Terms, Pajak Dinamis & Credit Limit Controls
	PaymentTerm         string          `gorm:"type:varchar(50);default:'Net 30'" json:"payment_term"` // COD, Net 14, Net 30, Net 60, DP 30%
	DueDate             time.Time       `json:"due_date"`                                              // Tanggal jatuh tempo
	TaxRate             float64         `gorm:"type:numeric(5,2);default:11.0" json:"tax_rate"`         // 0 (Non-PPN), 11 (PPN 11%), 12 (PPN 12%)
	TaxType             string          `gorm:"type:varchar(50);default:'PPN 11%'" json:"tax_type"`    // PPN 11%, PPN 12%, Non-PPN, PPh 23
	CreditStatus        string          `gorm:"type:varchar(50);default:'OK'" json:"credit_status"`    // OK, Warning, Exceeded, Hold
	CreditLimit         float64         `gorm:"type:numeric(15,2);default:0" json:"credit_limit"`      // Batas kredit pelanggan
	CurrentReceivable   float64         `gorm:"type:numeric(15,2);default:0" json:"current_receivable"`// Piutang saat ini
	IsCreditBypassed    bool            `gorm:"default:false" json:"is_credit_bypassed"`               // Izin bypass dari manajer finance
	BypassedBy          string          `gorm:"type:varchar(100)" json:"bypassed_by"`
	Nsfp                string          `gorm:"type:varchar(50)" json:"nsfp"`                          // Nomor Seri Faktur Pajak E-Faktur DJP

	// FASE 4: Salesperson Commission & KPI Fields
	SalespersonName     string          `gorm:"type:varchar(100);default:'Sales Team'" json:"salesperson_name"`
	CommissionRate      float64         `gorm:"type:numeric(5,2);default:3.0" json:"commission_rate"` // e.g. 3% standard commission
	CommissionAmount    float64         `gorm:"type:numeric(15,2);default:0" json:"commission_amount"`
	CommissionStatus    string          `gorm:"type:varchar(50);default:'Unpaid'" json:"commission_status"` // Unpaid, Paid
	Notes               string          `gorm:"type:text" json:"notes"`
	OrderLines          []SaleOrderLine `gorm:"foreignKey:OrderID" json:"order_lines"`
	CreatedAt           time.Time       `json:"created_at"`
}

type SaleOrderLine struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	OrderID      uint    `json:"order_id"`
	Order *SaleOrder `gorm:"foreignKey:OrderID" json:"order,omitempty"` // Odoo relation mapped
	ProductID    uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Description  string  `gorm:"type:varchar(255)" json:"description"`
	Quantity     float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	DeliveredQty float64 `gorm:"type:numeric(15,2);default:0" json:"delivered_qty"`
	InvoicedQty  float64 `gorm:"type:numeric(15,2);default:0" json:"invoiced_qty"`
	UnitPrice    float64 `gorm:"type:numeric(15,2);default:0" json:"unit_price"`
	Discount     float64 `gorm:"type:numeric(5,2);default:0" json:"discount"` // Percentage
	TaxRate      float64 `gorm:"type:numeric(5,2);default:11.0" json:"tax_rate"`
	SubTotal     float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
	IsOptional   bool    `gorm:"default:false" json:"is_optional"` // Upselling item
}


func (Pricelist) TableName() string {
	return "sales.pricelists"
}

func (PricelistItem) TableName() string {
	return "sales.pricelist_items"
}

func (QuotationTemplate) TableName() string {
	return "sales.quotation_templates"
}

func (DeliveryMethod) TableName() string {
	return "sales.delivery_methods"
}

func (SaleOrder) TableName() string {
	return "sales.sale_orders"
}

func (SaleOrderLine) TableName() string {
	return "sales.sale_order_lines"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Pricelist{}, &PricelistItem{}, &QuotationTemplate{}, &DeliveryMethod{}, &SaleOrder{}, &SaleOrderLine{})
}
