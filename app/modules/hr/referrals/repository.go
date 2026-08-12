package referrals

import (
	"ERP-System/config"
)

func CreateReferral(data *Referral) error {
	return config.DB.Create(data).Error
}

func GetAllReferral() ([]Referral, error) {
	var list []Referral
	err := config.DB.Find(&list).Error
	return list, err
}

func GetReferralByID(id uint) (*Referral, error) {
	var data Referral
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateReferral(data *Referral) error {
	return config.DB.Save(data).Error
}

func DeleteReferral(id uint) error {
	return config.DB.Delete(&Referral{}, id).Error
}
