package referrals

func CreateReferralRewardService(data *ReferralReward) error {
	return CreateReferralReward(data)
}

func GetAllReferralRewardService() ([]ReferralReward, error) {
	return GetAllReferralReward()
}

func GetReferralRewardByIDService(id uint) (*ReferralReward, error) {
	return GetReferralRewardByID(id)
}

func UpdateReferralRewardService(data *ReferralReward) error {
	return UpdateReferralReward(data)
}

func DeleteReferralRewardService(id uint) error {
	return DeleteReferralReward(id)
}
