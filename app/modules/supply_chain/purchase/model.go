package purchase

import (
	"ERP-System/config"
	"time"
)

type PurchaseRequisition struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"` // TE/2026/001
	State     string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, in_progress, open, done, cancel
	DateEnd   time.Time `json:"date_end"` // Deadline for vendor bids
	CreatedAt time.Time `json:"created_at"`
}

type ProductSupplierInfo struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	PartnerID uint    `json:"partner_id"` // Vendor
	ProductID uint    `json:"product_id"`
	MinQty    float64 `gorm:"type:numeric(15,2);default:0" json:"min_qty"`
	Price     float64 `gorm:"type:numeric(15,2);not null" json:"price"`
	Delay     int     `gorm:"default:1" json:"delay"` // Delivery lead time in days
}

type PurchaseOrder struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(100);not null" json:"name"` // PO/2026/001
	PartnerID     uint      `json:"partner_id"`                             // Vendor
	RequisitionID *uint     `json:"requisition_id"`                         // Link to Blanket Order
	State         string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, sent, purchase, done, cancel
	AmountUntaxed float64   `gorm:"type:numeric(15,2);default:0" json:"amount_untaxed"`
	AmountTax     float64   `gorm:"type:numeric(15,2);default:0" json:"amount_tax"`
	AmountTotal   float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	DateOrder     time.Time `json:"date_order"`
	CreatedAt     time.Time `json:"created_at"`
}

type PurchaseOrderLine struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	OrderID     uint    `json:"order_id"`
	ProductID   uint    `json:"product_id"`
	Name        string  `gorm:"type:varchar(255)" json:"name"` // Description
	Quantity    float64 `gorm:"type:numeric(15,2);not null;default:1" json:"quantity"`
	QtyReceived float64 `gorm:"type:numeric(15,2);default:0" json:"qty_received"`
	QtyInvoiced float64 `gorm:"type:numeric(15,2);default:0" json:"qty_invoiced"`
	PriceUnit   float64 `gorm:"type:numeric(15,2);not null;default:0" json:"price_unit"`
	TaxesID     *uint   `json:"taxes_id"`
	PriceSubtotal float64 `gorm:"type:numeric(15,2);default:0" json:"price_subtotal"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PurchaseRequisition{}, &ProductSupplierInfo{}, &PurchaseOrder{}, &PurchaseOrderLine{})
}
