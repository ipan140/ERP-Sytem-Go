package purchase

import (
	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"time"
)

type PurchaseRequisition struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`        // TE/2026/001
	State     string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, in_progress, open, done, cancel
	DateEnd   time.Time `json:"date_end"`                                      // Deadline for vendor bids
	CreatedAt time.Time `json:"created_at"`
}

type ProductSupplierInfo struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	PartnerID uint    `json:"partner_id"` // Vendor
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` // Cross-module relation
	ProductID uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	MinQty    float64 `gorm:"type:numeric(15,2);default:0" json:"min_qty"`
	Price     float64 `gorm:"type:numeric(15,2);not null" json:"price"`
	Delay     int     `gorm:"default:1" json:"delay"` // Delivery lead time in days
}

type PurchaseOrder struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Name          string    `gorm:"type:varchar(100);not null" json:"name"`        // PO/2026/001
	PartnerID     uint      `json:"partner_id"`                                    // Vendor
	Partner *base.Partner `gorm:"foreignKey:PartnerID" json:"partner,omitempty"` // Cross-module relation
	RequisitionID *uint     `json:"requisition_id"`                                // Link to Blanket Order
	Requisition *PurchaseRequisition `gorm:"foreignKey:RequisitionID" json:"requisition,omitempty"` // Odoo relation mapped
	State         string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, sent, to_approve, purchase, done, cancel
	AmountUntaxed float64   `gorm:"type:numeric(15,2);default:0" json:"amount_untaxed"`
	AmountTax     float64   `gorm:"type:numeric(15,2);default:0" json:"amount_tax"`
	AmountTotal   float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	DateOrder     time.Time `json:"date_order"`
	Notes         string    `gorm:"type:text" json:"notes"`
	CompanyID     uint      `json:"company_id"`
	ApprovedBy    *uint     `json:"approved_by"`
	ApprovedAt    *time.Time `json:"approved_at"`
	CreatedAt     time.Time `json:"created_at"`
	OrderLines    []PurchaseOrderLine `gorm:"foreignKey:OrderID;-:migration" json:"order_lines,omitempty"`
}

type PurchaseOrderLine struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	OrderID       uint    `json:"order_id"`
	Order *PurchaseOrder  `gorm:"foreignKey:OrderID;-:migration" json:"order,omitempty"`
	ProductID     uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Name          string  `gorm:"type:varchar(255)" json:"name"` // Description
	Quantity      float64 `gorm:"type:numeric(15,2);not null;default:1" json:"quantity"`
	QtyReceived   float64 `gorm:"type:numeric(15,2);default:0" json:"qty_received"`
	QtyInvoiced   float64 `gorm:"type:numeric(15,2);default:0" json:"qty_invoiced"`
	PriceUnit     float64 `gorm:"type:numeric(15,2);not null;default:0" json:"price_unit"`
	TaxesID       *uint   `json:"taxes_id"`
	Taxes *invoicing.Tax `gorm:"foreignKey:TaxesID" json:"taxes,omitempty"` // Odoo relation mapped
	PriceSubtotal float64 `gorm:"type:numeric(15,2);default:0" json:"price_subtotal"`
}

// Enterprise SCM Procurement DTOs (Fase 2)
type PurchaseSummary struct {
	TotalSpentMonthly float64 `json:"total_spent_monthly"`
	ToApproveCount    int64   `json:"to_approve_count"`
	ToReceiveCount    int64   `json:"to_receive_count"`
	ActiveVendorCount int64   `json:"active_vendor_count"`
	TotalPOCount      int64   `json:"total_po_count"`
}

type CreatePOLineInput struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Name      string  `json:"name"`
	Quantity  float64 `json:"quantity" validate:"required,gt=0"`
	PriceUnit float64 `json:"price_unit" validate:"required,gte=0"`
}

type CreatePORequest struct {
	PartnerID uint                `json:"partner_id" validate:"required"`
	DateOrder *time.Time          `json:"date_order"`
	Notes     string              `json:"notes"`
	Lines     []CreatePOLineInput `json:"lines" validate:"required,dive"`
}

type ReceiveItemInput struct {
	LineID      uint    `json:"line_id" validate:"required"`
	QtyReceived float64 `json:"qty_received" validate:"required,gt=0"`
}

type ReceiveGoodsRequest struct {
	WarehouseID uint               `json:"warehouse_id"`
	Items       []ReceiveItemInput `json:"items" validate:"required,dive"`
	Notes       string             `json:"notes"`
}


func (PurchaseRequisition) TableName() string {
	return "supply_chain.purchase_requisitions"
}

func (ProductSupplierInfo) TableName() string {
	return "supply_chain.product_supplier_infos"
}

func (PurchaseOrder) TableName() string {
	return "supply_chain.purchase_orders"
}

func (PurchaseOrderLine) TableName() string {
	return "supply_chain.purchase_order_lines"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PurchaseRequisition{}, &ProductSupplierInfo{}, &PurchaseOrder{}, &PurchaseOrderLine{})
}
