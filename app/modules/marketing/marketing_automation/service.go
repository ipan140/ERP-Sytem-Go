package marketing_automation

func CreateAutomationCampaignService(data *AutomationCampaign) error {
	return CreateAutomationCampaign(data)
}

func GetAllAutomationCampaignService() ([]AutomationCampaign, error) {
	return GetAllAutomationCampaign()
}

func GetPaginatedAutomationCampaignService(offset, limit int, search string) ([]AutomationCampaign, int64, error) {
	return GetPaginatedAutomationCampaigns(offset, limit, search)
}

func GetAutomationCampaignByIDService(id uint) (*AutomationCampaign, error) {
	return GetAutomationCampaignByID(id)
}

func UpdateAutomationCampaignService(data *AutomationCampaign) error {
	return UpdateAutomationCampaign(data)
}

func DeleteAutomationCampaignService(id uint) error {
	return DeleteAutomationCampaign(id)
}

func CreateWorkflowActivityService(data *WorkflowActivity) error { return CreateWorkflowActivity(data) }
func GetAllWorkflowActivityService() ([]WorkflowActivity, error) { return GetAllWorkflowActivity() }
func GetWorkflowActivityByIDService(id uint) (*WorkflowActivity, error) {
	return GetWorkflowActivityByID(id)
}
func UpdateWorkflowActivityService(data *WorkflowActivity) error { return UpdateWorkflowActivity(data) }
func DeleteWorkflowActivityService(id uint) error                { return DeleteWorkflowActivity(id) }

func GetJourneyLogsService(campaignID uint) ([]JourneyLog, error) {
	return GetJourneyLogsByCampaign(campaignID)
}

// RunJourneySimulatorService mensimulasikan perjalanan prospek (Leads) menembus alur omnichannel
func RunJourneySimulatorService(campaignID uint) ([]JourneyLog, error) {
	campaign, err := GetAutomationCampaignByID(campaignID)
	if err != nil {
		return nil, err
	}

	activities, err := GetAllWorkflowActivity()
	if err != nil {
		return nil, err
	}

	var campaignActs []WorkflowActivity
	for _, act := range activities {
		if act.CampaignID == campaign.ID {
			campaignActs = append(campaignActs, act)
		}
	}

	// Buat simulasi log untuk 3 prospek nyata/contoh
	leads := []struct {
		Name    string
		Contact string
	}{
		{"Budi Santoso (PT Maju Jaya)", "+62 812-3456-7890"},
		{"Siti Rahma (CV Berkah Mandiri)", "+62 813-8822-1100"},
		{"Hendra Wijaya (PT Sinar Terang)", "+62 818-9990-2211"},
	}

	var createdLogs []JourneyLog
	nowStr := "Baru saja"

	for _, lead := range leads {
		for _, act := range campaignActs {
			channel := act.Channel
			if channel == "" {
				channel = act.ActionType
			}
			logEntry := JourneyLog{
				CampaignID:  campaign.ID,
				ActivityID:  act.ID,
				LeadName:    lead.Name,
				LeadContact: lead.Contact,
				Channel:     channel,
				ActionName:  act.ActivityName,
				Status:      "Delivered",
				ExecutedAt:  nowStr,
			}
			_ = CreateJourneyLog(&logEntry)
			createdLogs = append(createdLogs, logEntry)
		}
	}

	return GetJourneyLogsByCampaign(campaignID)
}
