package purchase

func CreatePurchaseOrderService(data *PurchaseOrder) error {
	return CreatePurchaseOrder(data)
}

func GetAllPurchaseOrderService() ([]PurchaseOrder, error) {
	return GetAllPurchaseOrder()
}

func GetPurchaseOrderByIDService(id uint) (*PurchaseOrder, error) {
	return GetPurchaseOrderByID(id)
}

func UpdatePurchaseOrderService(data *PurchaseOrder) error {
	return UpdatePurchaseOrder(data)
}

func DeletePurchaseOrderService(id uint) error {
	return DeletePurchaseOrder(id)
}
