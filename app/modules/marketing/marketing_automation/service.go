package marketing_automation

func CreateAutomationCampaignService(data *AutomationCampaign) error {
	return CreateAutomationCampaign(data)
}

func GetAllAutomationCampaignService() ([]AutomationCampaign, error) {
	return GetAllAutomationCampaign()
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
