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

func CreateReferralPointService(data *ReferralPoint) error { return CreateReferralPoint(data) }
func GetAllReferralPointService() ([]ReferralPoint, error) { return GetAllReferralPoint() }
func GetReferralPointByIDService(id uint) (*ReferralPoint, error) { return GetReferralPointByID(id) }
func UpdateReferralPointService(data *ReferralPoint) error { return UpdateReferralPoint(data) }
func DeleteReferralPointService(id uint) error { return DeleteReferralPoint(id) }
