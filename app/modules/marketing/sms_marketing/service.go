package sms_marketing

func CreateSmsCampaignService(data *SmsCampaign) error {
	return CreateSmsCampaign(data)
}

func GetAllSmsCampaignService() ([]SmsCampaign, error) {
	return GetAllSmsCampaign()
}

func GetSmsCampaignByIDService(id uint) (*SmsCampaign, error) {
	return GetSmsCampaignByID(id)
}

func UpdateSmsCampaignService(data *SmsCampaign) error {
	return UpdateSmsCampaign(data)
}

func DeleteSmsCampaignService(id uint) error {
	return DeleteSmsCampaign(id)
}
