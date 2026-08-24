package mass_mailing

import "ERP-System/config"

// 4. Mass Mailing & Tracking
type MailingCampaign struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Name         string `gorm:"type:varchar(255)" json:"name"`
	SentCount    int    `gorm:"default:0" json:"sent_count"`
	OpenedCount  int    `gorm:"default:0" json:"opened_count"`
	ClickedCount int    `gorm:"default:0" json:"clicked_count"`
}

// 10. UTM Link Tracker
type UtmTracker struct {
	ID               uint    `gorm:"primaryKey" json:"id"`
	CampaignID       uint    `json:"campaign_id"`
	Campaign *MailingCampaign `gorm:"foreignKey:CampaignID" json:"campaign,omitempty"` // Odoo relation mapped
	UtmSource        string  `gorm:"type:varchar(50)" json:"utm_source"`          // misal: facebook
	UtmMedium        string  `gorm:"type:varchar(50)" json:"utm_medium"`          // misal: cpc
	GeneratedRevenue float64 `gorm:"type:numeric(15,2)" json:"generated_revenue"` // Untung dari link ini
}


func (MailingCampaign) TableName() string {
	return "marketing.mailing_campaigns"
}

func (UtmTracker) TableName() string {
	return "marketing.utm_trackers"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &MailingCampaign{}, &UtmTracker{})
}
