package referrals

func CreateReferralService(data *Referral) error {
	return CreateReferral(data)
}

func GetAllReferralService() ([]Referral, error) {
	return GetAllReferral()
}

func GetReferralByIDService(id uint) (*Referral, error) {
	return GetReferralByID(id)
}

func UpdateReferralService(data *Referral) error {
	return UpdateReferral(data)
}

func DeleteReferralService(id uint) error {
	return DeleteReferral(id)
}
