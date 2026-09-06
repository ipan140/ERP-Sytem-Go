package sms_marketing

import (
	"ERP-System/config"
	"time"
)

type SmsCampaign struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	Name           string     `gorm:"type:varchar(255);not null" json:"name"`
	Channel        string     `gorm:"type:varchar(50);default:'SMS'" json:"channel"` // SMS, WhatsApp, All
	Content        string     `gorm:"type:text" json:"content"`
	TargetAudience string     `gorm:"type:varchar(255)" json:"target_audience"`
	Status         string     `gorm:"type:varchar(50);default:'Draft'" json:"status"` // Draft, In-Queue, Sent
	SentCount      int        `gorm:"default:0" json:"sent_count"`
	DeliveredCount int        `gorm:"default:0" json:"delivered_count"`
	ReadCount      int        `gorm:"default:0" json:"read_count"` // WA Blue Ticks
	TemplateID     *uint      `json:"template_id"`
	Template       *WaTemplate `gorm:"foreignKey:TemplateID" json:"template,omitempty"`
	ScheduledAt    *time.Time `json:"scheduled_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

// WhatsApp Template Model (Meta Cloud API / WABA Standard)
type WaTemplate struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. promo_gajian_q3
	Category    string    `gorm:"type:varchar(50);default:'MARKETING'" json:"category"` // MARKETING, UTILITY, AUTHENTICATION
	Language    string    `gorm:"type:varchar(10);default:'id'" json:"language"`
	HeaderType  string    `gorm:"type:varchar(20);default:'TEXT'" json:"header_type"` // NONE, TEXT, IMAGE, DOCUMENT
	HeaderText  string    `gorm:"type:varchar(255)" json:"header_text"`
	BodyText    string    `gorm:"type:text;not null" json:"body_text"`
	FooterText  string    `gorm:"type:varchar(255)" json:"footer_text"`
	ButtonType  string    `gorm:"type:varchar(50);default:'NONE'" json:"button_type"` // NONE, QUICK_REPLY, CALL_TO_ACTION
	ButtonLabel string    `gorm:"type:varchar(100)" json:"button_label"`
	ButtonURL   string    `gorm:"type:varchar(255)" json:"button_url"`
	Status      string    `gorm:"type:varchar(50);default:'APPROVED'" json:"status"` // PENDING, APPROVED, REJECTED
	CreatedAt   time.Time `json:"created_at"`
}

// WhatsApp Meta / WABA Configuration
type WaConfig struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	PhoneNumberID  string    `gorm:"type:varchar(100)" json:"phone_number_id"`
	WabaID         string    `gorm:"type:varchar(100)" json:"waba_id"`
	AccessToken    string    `gorm:"type:text" json:"access_token"`
	ApiVersion     string    `gorm:"type:varchar(20);default:'v19.0'" json:"api_version"`
	WebhookToken   string    `gorm:"type:varchar(100)" json:"webhook_token"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SmsCampaign{}, &WaTemplate{}, &WaConfig{})
}


// ---- Auto-Generated TableName methods ----
func (SmsCampaign) TableName() string {
	return "marketing.sms_campaigns"
}

func (WaTemplate) TableName() string {
	return "marketing.wa_templates"
}

func (WaConfig) TableName() string {
	return "marketing.wa_configs"
}