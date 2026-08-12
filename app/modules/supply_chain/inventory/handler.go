package inventory

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateProduct godoc
// @Summary Create a new Product
// @Description Create a new Product in the system
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/inventory [post]
// @Security BearerAuth
func CreateProductHandler(c echo.Context) error {
	var data Product
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateProductService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllProduct godoc
// @Summary Get all Product
// @Description Retrieve a list of all Product
// @Tags supply_chain-inventory
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory [get]
// @Security BearerAuth
func GetAllProductHandler(c echo.Context) error {
	data, err := GetAllProductService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetProductByID godoc
// @Summary Get a Product by ID
// @Description Retrieve a specific Product by its ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [get]
// @Security BearerAuth
func GetProductByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateProduct godoc
// @Summary Update a Product
// @Description Update an existing Product
// @Tags supply_chain-inventory
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [put]
// @Security BearerAuth
func UpdateProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetProductByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateProductService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteProduct godoc
// @Summary Delete a Product
// @Description Delete a Product by ID
// @Tags supply_chain-inventory
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/inventory/{id} [delete]
// @Security BearerAuth
func DeleteProductHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteProductService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
