package ecommerce

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
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
