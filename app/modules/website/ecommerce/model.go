package ecommerce



import (
	"ERP-System/config"
	"time"
)

// 2. Portal Pelanggan B2B/B2C
type PortalUser struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CustomerID uint      `json:"customer_id"` // Nyambung ke tabel Partner
	Password   string    `gorm:"type:varchar(255)" json:"password"`
	LastLogin  time.Time `json:"last_login"`
}

// 8. Abandoned Cart Recovery (Penyelamat Keranjang)
type ShoppingCart struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CustomerID  *uint     `json:"customer_id"`
	SessionID   string    `gorm:"type:varchar(255)" json:"session_id"`
	IsAbandoned bool      `gorm:"default:false" json:"is_abandoned"`
	RecoverySent bool     `gorm:"default:false" json:"recovery_sent"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CartItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Quantity  float64 `gorm:"type:numeric(15,2)" json:"quantity"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &PortalUser{}, &ShoppingCart{}, &CartItem{})
}


type Cart struct { ID uint `gorm:"primaryKey"` }

