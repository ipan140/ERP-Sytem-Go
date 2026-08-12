package subscriptions

func CreateSubscriptionService(data *Subscription) error {
	return CreateSubscription(data)
}

func GetAllSubscriptionService() ([]Subscription, error) {
	return GetAllSubscription()
}

func GetSubscriptionByIDService(id uint) (*Subscription, error) {
	return GetSubscriptionByID(id)
}

func UpdateSubscriptionService(data *Subscription) error {
	return UpdateSubscription(data)
}

func DeleteSubscriptionService(id uint) error {
	return DeleteSubscription(id)
}
