package lunch

func CreateLunchOrderService(data *LunchOrder) error {
	return CreateLunchOrder(data)
}

func GetAllLunchOrderService() ([]LunchOrder, error) {
	return GetAllLunchOrder()
}

func GetLunchOrderByIDService(id uint) (*LunchOrder, error) {
	return GetLunchOrderByID(id)
}

func UpdateLunchOrderService(data *LunchOrder) error {
	return UpdateLunchOrder(data)
}

func DeleteLunchOrderService(id uint) error {
	return DeleteLunchOrder(id)
}
