package subscriptions

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateSubscription(data *Subscription) error {
	return config.DB.Create(data).Error
}

func GetAllSubscription() ([]Subscription, error) {
	var list []Subscription
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedSubscriptions(offset int, limit int, search string, state string) ([]Subscription, int64, error) {
	var list []Subscription
	var total int64

	query := config.DB.Model(&Subscription{}).Preload(clause.Associations)

	if state != "" && state != "All" && state != "all" {
		query = query.Where("state = ?", state)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ?", s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetSubscriptionByID(id uint) (*Subscription, error) {
	var data Subscription
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateSubscription(data *Subscription) error {
	return config.DB.Save(data).Error
}

func DeleteSubscription(id uint) error {
	return config.DB.Delete(&Subscription{}, id).Error
}

func CreateSubscriptionPlan(data *SubscriptionPlan) error { return config.DB.Create(data).Error }
func GetAllSubscriptionPlan() ([]SubscriptionPlan, error) {
	var list []SubscriptionPlan
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSubscriptionPlanByID(id uint) (*SubscriptionPlan, error) {
	var data SubscriptionPlan
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSubscriptionPlan(data *SubscriptionPlan) error { return config.DB.Save(data).Error }
func DeleteSubscriptionPlan(id uint) error                { return config.DB.Delete(&SubscriptionPlan{}, id).Error }
