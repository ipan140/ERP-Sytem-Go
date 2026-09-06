package marketing_automation

import "ERP-System/config"

type WorkflowActivity struct {
	ID               uint   `gorm:"primaryKey" json:"id"`
	CampaignID       uint   `json:"campaign_id"`
	ActivityName     string `gorm:"type:varchar(255)" json:"activity_name"`
	ActionType       string `gorm:"type:varchar(50)" json:"action_type"` // Email, WhatsApp, SMS, CRM_Task, Webhook
	Channel          string `gorm:"type:varchar(50);default:'Email'" json:"channel"` // Email, WhatsApp, SMS, CRM
	TargetTemplateID uint   `json:"target_template_id"`
	ActionPayload    string `gorm:"type:text" json:"action_payload"` // Isi pesan / instruksi task CRM
	DelayHours       int    `json:"delay_hours"`
	Condition        string `gorm:"type:varchar(50)" json:"condition"` // Opened, Clicked, Always, Not_Replied
}

type JourneyLog struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CampaignID   uint   `json:"campaign_id"`
	ActivityID   uint   `json:"activity_id"`
	LeadName     string `gorm:"type:varchar(255)" json:"lead_name"`
	LeadContact  string `gorm:"type:varchar(100)" json:"lead_contact"`
	Channel      string `gorm:"type:varchar(50)" json:"channel"`
	ActionName   string `gorm:"type:varchar(255)" json:"action_name"`
	Status       string `gorm:"type:varchar(50);default:'Delivered'" json:"status"` // Delivered, Pending, Failed
	ExecutedAt   string `gorm:"type:varchar(50)" json:"executed_at"`
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

func (JourneyLog) TableName() string {
	return "marketing.journey_logs"
}

func (AutomationCampaign) TableName() string {
	return "marketing.automation_campaigns"
}

func init() { 
	config.ModelsToMigrate = append(config.ModelsToMigrate, &WorkflowActivity{}, &AutomationCampaign{}, &JourneyLog{}) 
}
