package subscriptions

import (
	"ERP-System/config"
)

func CreateSubscription(data *Subscription) error {
	return config.DB.Create(data).Error
}

func GetAllSubscription() ([]Subscription, error) {
	var list []Subscription
	err := config.DB.Find(&list).Error
	return list, err
}

func GetSubscriptionByID(id uint) (*Subscription, error) {
	var data Subscription
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateSubscription(data *Subscription) error {
	return config.DB.Save(data).Error
}

func DeleteSubscription(id uint) error {
	return config.DB.Delete(&Subscription{}, id).Error
}
