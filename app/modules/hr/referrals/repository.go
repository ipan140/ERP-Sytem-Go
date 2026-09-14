package referrals

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateReferralReward(data *ReferralReward) error {
	return config.DB.Create(data).Error
}

func GetAllReferralReward() ([]ReferralReward, error) {
	var list []ReferralReward
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedReferralRewards(offset, limit int, search string) ([]ReferralReward, int64, error) {
	var list []ReferralReward
	var total int64

	query := config.DB.Model(&ReferralReward{})

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("hrd.referral_rewards.name ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.referral_rewards.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetPaginatedReferralPoints(offset, limit int, search, employeeID string) ([]ReferralPoint, int64, error) {
	var list []ReferralPoint
	var total int64

	query := config.DB.Model(&ReferralPoint{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.referral_points.employee_id = ?", employeeID)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.referral_points.employee_id").
			Where("hrd.employees.name ILIKE ? OR hrd.referral_points.reason ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.referral_points.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetReferralRewardByID(id uint) (*ReferralReward, error) {
	var data ReferralReward
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateReferralReward(data *ReferralReward) error {
	return config.DB.Save(data).Error
}

func DeleteReferralReward(id uint) error {
	return config.DB.Delete(&ReferralReward{}, id).Error
}

func CreateReferralPoint(data *ReferralPoint) error { return config.DB.Create(data).Error }
func GetAllReferralPoint() ([]ReferralPoint, error) {
	var list []ReferralPoint
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetReferralPointByID(id uint) (*ReferralPoint, error) {
	var data ReferralPoint
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateReferralPoint(data *ReferralPoint) error { return config.DB.Save(data).Error }
func DeleteReferralPoint(id uint) error             { return config.DB.Delete(&ReferralPoint{}, id).Error }
