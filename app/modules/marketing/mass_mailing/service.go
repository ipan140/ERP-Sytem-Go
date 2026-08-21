package mass_mailing

func CreateMailingCampaignService(data *MailingCampaign) error {
	return CreateMailingCampaign(data)
}

func GetAllMailingCampaignService() ([]MailingCampaign, error) {
	return GetAllMailingCampaign()
}

func GetMailingCampaignByIDService(id uint) (*MailingCampaign, error) {
	return GetMailingCampaignByID(id)
}

func UpdateMailingCampaignService(data *MailingCampaign) error {
	return UpdateMailingCampaign(data)
}

func DeleteMailingCampaignService(id uint) error {
	return DeleteMailingCampaign(id)
}

func CreateUtmTrackerService(data *UtmTracker) error        { return CreateUtmTracker(data) }
func GetAllUtmTrackerService() ([]UtmTracker, error)        { return GetAllUtmTracker() }
func GetUtmTrackerByIDService(id uint) (*UtmTracker, error) { return GetUtmTrackerByID(id) }
func UpdateUtmTrackerService(data *UtmTracker) error        { return UpdateUtmTracker(data) }
func DeleteUtmTrackerService(id uint) error                 { return DeleteUtmTracker(id) }
