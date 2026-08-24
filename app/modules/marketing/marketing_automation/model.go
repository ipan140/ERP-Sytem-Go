package marketing_automation

import "ERP-System/config"

type WorkflowActivity struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	CampaignID uint   `json:"campaign_id"`
	ActionType string `gorm:"type:varchar(50)" json:"action_type"`
	DelayHours int    `json:"delay_hours"`                         
	Condition  string `gorm:"type:varchar(50)" json:"condition"`   
}

type AutomationCampaign struct {
	ID uint `gorm:"primaryKey"`
}

func (WorkflowActivity) TableName() string {
	return "marketing.workflow_activities"
}

func (AutomationCampaign) TableName() string {
	return "marketing.automation_campaigns"
}

func init() { 
	config.ModelsToMigrate = append(config.ModelsToMigrate, &WorkflowActivity{}, &AutomationCampaign{}) 
}
