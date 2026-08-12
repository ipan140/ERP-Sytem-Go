package sales_core

func CreateSaleOrderService(data *SaleOrder) error {
	return CreateSaleOrder(data)
}

func GetAllSaleOrderService() ([]SaleOrder, error) {
	return GetAllSaleOrder()
}

func GetSaleOrderByIDService(id uint) (*SaleOrder, error) {
	return GetSaleOrderByID(id)
}

func UpdateSaleOrderService(data *SaleOrder) error {
	return UpdateSaleOrder(data)
}

func DeleteSaleOrderService(id uint) error {
	return DeleteSaleOrder(id)
}
