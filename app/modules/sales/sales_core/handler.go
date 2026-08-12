package sales_core

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateSaleOrder godoc
// @Summary Create a new SaleOrder
// @Description Create a new SaleOrder in the system
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/sales_core [post]
// @Security BearerAuth
func CreateSaleOrderHandler(c echo.Context) error {
	var data SaleOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSaleOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSaleOrder godoc
// @Summary Get all SaleOrder
// @Description Retrieve a list of all SaleOrder
// @Tags sales-sales_core
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core [get]
// @Security BearerAuth
func GetAllSaleOrderHandler(c echo.Context) error {
	data, err := GetAllSaleOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSaleOrderByID godoc
// @Summary Get a SaleOrder by ID
// @Description Retrieve a specific SaleOrder by its ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [get]
// @Security BearerAuth
func GetSaleOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSaleOrder godoc
// @Summary Update a SaleOrder
// @Description Update an existing SaleOrder
// @Tags sales-sales_core
// @Accept json
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [put]
// @Security BearerAuth
func UpdateSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSaleOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSaleOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSaleOrder godoc
// @Summary Delete a SaleOrder
// @Description Delete a SaleOrder by ID
// @Tags sales-sales_core
// @Produce json
// @Param id path int true "SaleOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/sales_core/{id} [delete]
// @Security BearerAuth
func DeleteSaleOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSaleOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
