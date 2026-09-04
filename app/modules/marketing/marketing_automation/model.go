package marketing_automation

import "ERP-System/config"

type WorkflowActivity struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CampaignID   uint   `json:"campaign_id"`
	ActivityName string `gorm:"type:varchar(255)" json:"activity_name"`
	ActionType   string `gorm:"type:varchar(50)" json:"action_type"` // Email, SMS, Notification, Webhook
	DelayHours   int    `json:"delay_hours"`                         
	Condition    string `gorm:"type:varchar(50)" json:"condition"`   // Opened, Clicked, Always
}

type AutomationCampaign struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(255)" json:"name"`
	TriggerType string `gorm:"type:varchar(100)" json:"trigger_type"` // Lead Created, Invoice Paid, Form Submitted
	Status      string `gorm:"type:varchar(50);default:'Active'" json:"status"` // Active, Paused, Draft
	TargetModel string `gorm:"type:varchar(50);default:'Leads'" json:"target_model"`
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
