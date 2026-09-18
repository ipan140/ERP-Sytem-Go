package whatsapp

import (
	"ERP-System/config"
	"time"
)

// WhatsappConfig menyimpan kredensial API WhatsApp Gateway
type WhatsappConfig struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Provider   string    `gorm:"type:varchar(50);default:'WABA_CLOUD'" json:"provider"` // WABA_CLOUD, FONNTE, TWILIO
	ApiToken   string    `gorm:"type:text" json:"api_token"`
	PhoneID    string    `gorm:"type:varchar(100)" json:"phone_id"`
	WebhookURL string    `gorm:"type:varchar(255)" json:"webhook_url"`
	IsActive   bool      `gorm:"default:true" json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (WhatsappConfig) TableName() string {
	return "setting.whatsapp_configs"
}

// WaTemplate menyimpan master template pesan resmi (HSM)
type WaTemplate struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(100);default:''" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Category  string    `gorm:"type:varchar(100);default:'General'" json:"category"`
	Content   string    `gorm:"type:text;not null" json:"content"`
	Status    string    `gorm:"type:varchar(50);default:'APPROVED'" json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (WaTemplate) TableName() string {
	return "setting.wa_templates"
}

// WaLog mencatat riwayat semua pesan WhatsApp keluar/masuk
type WaLog struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Phone        string    `gorm:"type:varchar(50);not null" json:"phone"`
	Recipient    string    `gorm:"type:varchar(255)" json:"recipient"`
	TemplateName string    `gorm:"type:varchar(255)" json:"template_name"`
	Message      string    `gorm:"type:text" json:"message"`
	Status       string    `gorm:"type:varchar(50);default:'SENT'" json:"status"` // SENT, DELIVERED, READ, FAILED
	ErrorMessage string    `gorm:"type:text" json:"error_message"`
	SentAt       time.Time `json:"sent_at"`
	CreatedAt    time.Time `json:"created_at"`
}

func (WaLog) TableName() string {
	return "setting.wa_logs"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &WhatsappConfig{}, &WaTemplate{}, &WaLog{})
}
