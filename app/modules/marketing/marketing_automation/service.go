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
