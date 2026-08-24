package sms_marketing

import (
	"ERP-System/config"
	"time"
)

type SmsCampaign struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SmsCampaign{})
}


// ---- Auto-Generated TableName methods ----
func (SmsCampaign) TableName() string {
	return "marketing.sms_campaigns"
}