package inventory

func CreateProductService(data *Product) error {
	return CreateProduct(data)
}

func GetAllProductService() ([]Product, error) {
	return GetAllProduct()
}

func GetProductByIDService(id uint) (*Product, error) {
	return GetProductByID(id)
}

func UpdateProductService(data *Product) error {
	return UpdateProduct(data)
}

func DeleteProductService(id uint) error {
	return DeleteProduct(id)
}
