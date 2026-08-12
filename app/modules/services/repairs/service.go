package repairs

func CreateRepairOrderService(data *RepairOrder) error {
	return CreateRepairOrder(data)
}

func GetAllRepairOrderService() ([]RepairOrder, error) {
	return GetAllRepairOrder()
}

func GetRepairOrderByIDService(id uint) (*RepairOrder, error) {
	return GetRepairOrderByID(id)
}

func UpdateRepairOrderService(data *RepairOrder) error {
	return UpdateRepairOrder(data)
}

func DeleteRepairOrderService(id uint) error {
	return DeleteRepairOrder(id)
}
