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
