package referrals

import (
	"ERP-System/config"
)

func CreateReferralReward(data *ReferralReward) error {
	return config.DB.Create(data).Error
}

func GetAllReferralReward() ([]ReferralReward, error) {
	var list []ReferralReward
	err := config.DB.Find(&list).Error
	return list, err
}

func GetReferralRewardByID(id uint) (*ReferralReward, error) {
	var data ReferralReward
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateReferralReward(data *ReferralReward) error {
	return config.DB.Save(data).Error
}

func DeleteReferralReward(id uint) error {
	return config.DB.Delete(&ReferralReward{}, id).Error
}
