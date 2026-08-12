package purchase

import (
	"ERP-System/config"
	"time"
)

type PurchaseOrder struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"` // PO/2026/001
	PartnerID   uint      `json:"partner_id"`                             // Vendor ID
	DateOrder   time.Time `json:"date_order"`
	State       string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, purchase, done, cancel
	AmountTotal float64   `gorm:"type:numeric(15,2);default:0" json:"amount_total"`
	CreatedAt   time.Time `json:"created_at"`
}

type PurchaseOrderLine struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  float64 `gorm:"type:numeric(15,2);default:1" json:"quantity"`
	PriceUnit float64 `gorm:"type:numeric(15,2);default:0" json:"price_unit"`
	SubTotal  float64 `gorm:"type:numeric(15,2);default:0" json:"sub_total"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PurchaseOrder{}, &PurchaseOrderLine{})
}
