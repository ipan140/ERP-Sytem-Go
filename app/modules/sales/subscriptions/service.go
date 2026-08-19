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

func CreateSubscriptionPlanService(data *SubscriptionPlan) error { return CreateSubscriptionPlan(data) }
func GetAllSubscriptionPlanService() ([]SubscriptionPlan, error) { return GetAllSubscriptionPlan() }
func GetSubscriptionPlanByIDService(id uint) (*SubscriptionPlan, error) { return GetSubscriptionPlanByID(id) }
func UpdateSubscriptionPlanService(data *SubscriptionPlan) error { return UpdateSubscriptionPlan(data) }
func DeleteSubscriptionPlanService(id uint) error { return DeleteSubscriptionPlan(id) }
