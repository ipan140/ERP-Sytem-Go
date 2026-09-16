package marketing_automation

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAutomationCampaign(data *AutomationCampaign) error {
	return config.DB.Create(data).Error
}

func GetAllAutomationCampaign() ([]AutomationCampaign, error) {
	var list []AutomationCampaign
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedAutomationCampaigns(offset, limit int, search string) ([]AutomationCampaign, int64, error) {
	var list []AutomationCampaign
	var total int64
	query := config.DB.Model(&AutomationCampaign{}).Preload(clause.Associations)
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR trigger_type ILIKE ?", s, s)
	}
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := query.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetAutomationCampaignByID(id uint) (*AutomationCampaign, error) {
	var data AutomationCampaign
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAutomationCampaign(data *AutomationCampaign) error {
	return config.DB.Save(data).Error
}

func DeleteAutomationCampaign(id uint) error {
	return config.DB.Delete(&AutomationCampaign{}, id).Error
}

func CreateWorkflowActivity(data *WorkflowActivity) error { return config.DB.Create(data).Error }
func GetAllWorkflowActivity() ([]WorkflowActivity, error) {
	var list []WorkflowActivity
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetWorkflowActivityByID(id uint) (*WorkflowActivity, error) {
	var data WorkflowActivity
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateWorkflowActivity(data *WorkflowActivity) error { return config.DB.Save(data).Error }
func DeleteWorkflowActivity(id uint) error                { return config.DB.Delete(&WorkflowActivity{}, id).Error }

// JourneyLog Repository
func CreateJourneyLog(data *JourneyLog) error {
	return config.DB.Create(data).Error
}

func GetJourneyLogsByCampaign(campaignID uint) ([]JourneyLog, error) {
	var list []JourneyLog
	err := config.DB.Where("campaign_id = ?", campaignID).Order("id desc").Limit(50).Find(&list).Error
	return list, err
}
