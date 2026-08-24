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
	CreatedAt           time.Time `json:"created_at"`
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
