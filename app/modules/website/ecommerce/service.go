package ecommerce

func CreateCartService(data *Cart) error {
	return CreateCart(data)
}

func GetAllCartService() ([]Cart, error) {
	return GetAllCart()
}

func GetCartByIDService(id uint) (*Cart, error) {
	return GetCartByID(id)
}

func UpdateCartService(data *Cart) error {
	return UpdateCart(data)
}

func DeleteCartService(id uint) error {
	return DeleteCart(id)
}
