package manufacturing

func CreateManufacturingOrderService(data *ManufacturingOrder) error {
	return CreateManufacturingOrder(data)
}

func GetAllManufacturingOrderService() ([]ManufacturingOrder, error) {
	return GetAllManufacturingOrder()
}

func GetManufacturingOrderByIDService(id uint) (*ManufacturingOrder, error) {
	return GetManufacturingOrderByID(id)
}

func UpdateManufacturingOrderService(data *ManufacturingOrder) error {
	return UpdateManufacturingOrder(data)
}

func DeleteManufacturingOrderService(id uint) error {
	return DeleteManufacturingOrder(id)
}
