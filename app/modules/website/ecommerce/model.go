package ecommerce

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/supply_chain/inventory"
	"ERP-System/config"
	"time"
)

// 2. Portal Pelanggan B2B/B2C
type PortalUser struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `json:"customer_id"` // Nyambung ke tabel Partner
	Customer *base.Partner `gorm:"foreignKey:CustomerID" json:"customer,omitempty"` // Cross-module relation
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	LastLogin  time.Time `json:"last_login"`
}

// 8. Abandoned Cart Recovery (Penyelamat Keranjang)
type ShoppingCart struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CustomerID   *uint     `json:"customer_id"`
	Customer *base.Partner `gorm:"foreignKey:CustomerID" json:"customer,omitempty"` // Cross-module relation
	SessionID    string    `gorm:"type:varchar(255)" json:"session_id"`
	IsAbandoned  bool      `gorm:"default:false" json:"is_abandoned"`
	RecoverySent bool      `gorm:"default:false" json:"recovery_sent"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `json:"cart_id"`
	Cart *Cart `gorm:"foreignKey:CartID"` // Auto-added relation
	ProductID uint    `json:"product_id"`
	Product *inventory.Product `gorm:"foreignKey:ProductID" json:"product,omitempty"` // Cross-module relation
	Quantity  float64 `gorm:"type:numeric(15,2)" json:"quantity"`
}


func (PortalUser) TableName() string {
	return "website_portal.portal_users"
}

func (ShoppingCart) TableName() string {
	return "website_portal.shopping_carts"
}

func (CartItem) TableName() string {
	return "website_portal.cart_items"
}

func (Cart) TableName() string {
	return "website_portal.carts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PortalUser{}, &ShoppingCart{}, &CartItem{})
}

type Cart struct {
	ID uint `gorm:"primaryKey"`
}
