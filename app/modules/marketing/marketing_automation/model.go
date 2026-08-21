package marketing_automation

import "ERP-System/config"

// 6. Marketing Automation (Robot Penjual)
type WorkflowActivity struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CampaignID uint   `json:"campaign_id"`
	ActionType string `gorm:"type:varchar(50)" json:"action_type"` // send_email, send_sms
	DelayHours int    `json:"delay_hours"`                         // Tunggu berapa jam sebelum action
	Condition  string `gorm:"type:varchar(50)" json:"condition"`   // if_opened, if_ignored
}

func init() { config.ModelsToMigrate = append(config.ModelsToMigrate, &WorkflowActivity{}) }

type AutomationCampaign struct {
	ID uint `gorm:"primaryKey"`
}
