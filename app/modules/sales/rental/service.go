package rental

func CreateRentalOrderService(data *RentalOrder) error {
	return CreateRentalOrder(data)
}

func GetAllRentalOrderService() ([]RentalOrder, error) {
	return GetAllRentalOrder()
}

func GetRentalOrderByIDService(id uint) (*RentalOrder, error) {
	return GetRentalOrderByID(id)
}

func UpdateRentalOrderService(data *RentalOrder) error {
	return UpdateRentalOrder(data)
}

func DeleteRentalOrderService(id uint) error {
	return DeleteRentalOrder(id)
}
