package ecommerce

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateCart godoc
// @Summary Create a new Cart
// @Description Create a new Cart in the system
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/ecommerce [post]
// @Security BearerAuth
func CreateCartHandler(c echo.Context) error {
	var data Cart
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateCartService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllCart godoc
// @Summary Get all Cart
// @Description Retrieve a list of all Cart
// @Tags website-ecommerce
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce [get]
// @Security BearerAuth
func GetAllCartHandler(c echo.Context) error {
	data, err := GetAllCartService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetCartByID godoc
// @Summary Get a Cart by ID
// @Description Retrieve a specific Cart by its ID
// @Tags website-ecommerce
// @Produce json
// @Param id path int true "Cart ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/{id} [get]
// @Security BearerAuth
func GetCartByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCartByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateCart godoc
// @Summary Update a Cart
// @Description Update an existing Cart
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Param id path int true "Cart ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/{id} [put]
// @Security BearerAuth
func UpdateCartHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCartByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateCartService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteCart godoc
// @Summary Delete a Cart
// @Description Delete a Cart by ID
// @Tags website-ecommerce
// @Produce json
// @Param id path int true "Cart ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/{id} [delete]
// @Security BearerAuth
func DeleteCartHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCartService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create PortalUser
// @Description Create a new PortalUser
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/ecommerce/portaluser [post]
// @Security BearerAuth
func CreatePortalUserHandler(c echo.Context) error {
	var data PortalUser
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePortalUserService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all PortalUser
// @Description Retrieve a list of all PortalUser
// @Tags website-ecommerce
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/portaluser [get]
// @Security BearerAuth
func GetAllPortalUserHandler(c echo.Context) error {
	data, err := GetAllPortalUserService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetPortalUserByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPortalUserByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update PortalUser
// @Description Update an existing PortalUser
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Param id path int true "PortalUser ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/portaluser/{id} [put]
// @Security BearerAuth
func UpdatePortalUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPortalUserByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdatePortalUserService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete PortalUser
// @Description Delete PortalUser by ID
// @Tags website-ecommerce
// @Produce json
// @Param id path int true "PortalUser ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/portaluser/{id} [delete]
// @Security BearerAuth
func DeletePortalUserHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePortalUserService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create ShoppingCart
// @Description Create a new ShoppingCart
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/ecommerce/shoppingcart [post]
// @Security BearerAuth
func CreateShoppingCartHandler(c echo.Context) error {
	var data ShoppingCart
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateShoppingCartService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all ShoppingCart
// @Description Retrieve a list of all ShoppingCart
// @Tags website-ecommerce
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/shoppingcart [get]
// @Security BearerAuth
func GetAllShoppingCartHandler(c echo.Context) error {
	data, err := GetAllShoppingCartService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetShoppingCartByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetShoppingCartByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update ShoppingCart
// @Description Update an existing ShoppingCart
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Param id path int true "ShoppingCart ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/shoppingcart/{id} [put]
// @Security BearerAuth
func UpdateShoppingCartHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetShoppingCartByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateShoppingCartService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete ShoppingCart
// @Description Delete ShoppingCart by ID
// @Tags website-ecommerce
// @Produce json
// @Param id path int true "ShoppingCart ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/shoppingcart/{id} [delete]
// @Security BearerAuth
func DeleteShoppingCartHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteShoppingCartService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create CartItem
// @Description Create a new CartItem
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/ecommerce/cartitem [post]
// @Security BearerAuth
func CreateCartItemHandler(c echo.Context) error {
	var data CartItem
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateCartItemService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all CartItem
// @Description Retrieve a list of all CartItem
// @Tags website-ecommerce
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/cartitem [get]
// @Security BearerAuth
func GetAllCartItemHandler(c echo.Context) error {
	data, err := GetAllCartItemService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetCartItemByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCartItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update CartItem
// @Description Update an existing CartItem
// @Tags website-ecommerce
// @Accept json
// @Produce json
// @Param id path int true "CartItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/cartitem/{id} [put]
// @Security BearerAuth
func UpdateCartItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCartItemByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateCartItemService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete CartItem
// @Description Delete CartItem by ID
// @Tags website-ecommerce
// @Produce json
// @Param id path int true "CartItem ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/ecommerce/cartitem/{id} [delete]
// @Security BearerAuth
func DeleteCartItemHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCartItemService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}
