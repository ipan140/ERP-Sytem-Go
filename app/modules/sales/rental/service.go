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

func CreateRentalOrderLineService(data *RentalOrderLine) error { return CreateRentalOrderLine(data) }
func GetAllRentalOrderLineService() ([]RentalOrderLine, error) { return GetAllRentalOrderLine() }
func GetRentalOrderLineByIDService(id uint) (*RentalOrderLine, error) {
	return GetRentalOrderLineByID(id)
}
func UpdateRentalOrderLineService(data *RentalOrderLine) error { return UpdateRentalOrderLine(data) }
func DeleteRentalOrderLineService(id uint) error               { return DeleteRentalOrderLine(id) }
