package sms_marketing

import (
	"ERP-System/config"
	"time"
)

type SmsCampaign struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Channel        string    `gorm:"type:varchar(50);default:'SMS'" json:"channel"` // SMS, WhatsApp, All
	Content        string    `gorm:"type:text" json:"content"`
	TargetAudience string    `gorm:"type:varchar(255)" json:"target_audience"`
	Status         string    `gorm:"type:varchar(50);default:'Draft'" json:"status"` // Draft, In-Queue, Sent
	SentCount      int       `gorm:"default:0" json:"sent_count"`
	DeliveredCount int       `gorm:"default:0" json:"delivered_count"`
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SmsCampaign{})
}


// ---- Auto-Generated TableName methods ----
func (SmsCampaign) TableName() string {
	return "marketing.sms_campaigns"
}