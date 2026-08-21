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

func CreatePortalUserService(data *PortalUser) error        { return CreatePortalUser(data) }
func GetAllPortalUserService() ([]PortalUser, error)        { return GetAllPortalUser() }
func GetPortalUserByIDService(id uint) (*PortalUser, error) { return GetPortalUserByID(id) }
func UpdatePortalUserService(data *PortalUser) error        { return UpdatePortalUser(data) }
func DeletePortalUserService(id uint) error                 { return DeletePortalUser(id) }

func CreateShoppingCartService(data *ShoppingCart) error        { return CreateShoppingCart(data) }
func GetAllShoppingCartService() ([]ShoppingCart, error)        { return GetAllShoppingCart() }
func GetShoppingCartByIDService(id uint) (*ShoppingCart, error) { return GetShoppingCartByID(id) }
func UpdateShoppingCartService(data *ShoppingCart) error        { return UpdateShoppingCart(data) }
func DeleteShoppingCartService(id uint) error                   { return DeleteShoppingCart(id) }

func CreateCartItemService(data *CartItem) error        { return CreateCartItem(data) }
func GetAllCartItemService() ([]CartItem, error)        { return GetAllCartItem() }
func GetCartItemByIDService(id uint) (*CartItem, error) { return GetCartItemByID(id) }
func UpdateCartItemService(data *CartItem) error        { return UpdateCartItem(data) }
func DeleteCartItemService(id uint) error               { return DeleteCartItem(id) }
