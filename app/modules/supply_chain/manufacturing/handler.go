package manufacturing

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateManufacturingOrder godoc
// @Summary Create a new ManufacturingOrder
// @Description Create a new ManufacturingOrder in the system
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing [post]
// @Security BearerAuth
func CreateManufacturingOrderHandler(c echo.Context) error {
	var data ManufacturingOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateManufacturingOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllManufacturingOrder godoc
// @Summary Get all ManufacturingOrder
// @Description Retrieve a list of all ManufacturingOrder
// @Tags supply_chain-manufacturing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing [get]
// @Security BearerAuth
func GetAllManufacturingOrderHandler(c echo.Context) error {
	data, err := GetAllManufacturingOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetManufacturingOrderByID godoc
// @Summary Get a ManufacturingOrder by ID
// @Description Retrieve a specific ManufacturingOrder by its ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "ManufacturingOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/{id} [get]
// @Security BearerAuth
func GetManufacturingOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetManufacturingOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateManufacturingOrder godoc
// @Summary Update a ManufacturingOrder
// @Description Update an existing ManufacturingOrder
// @Tags supply_chain-manufacturing
// @Accept json
// @Produce json
// @Param id path int true "ManufacturingOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/{id} [put]
// @Security BearerAuth
func UpdateManufacturingOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetManufacturingOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateManufacturingOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteManufacturingOrder godoc
// @Summary Delete a ManufacturingOrder
// @Description Delete a ManufacturingOrder by ID
// @Tags supply_chain-manufacturing
// @Produce json
// @Param id path int true "ManufacturingOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/manufacturing/{id} [delete]
// @Security BearerAuth
func DeleteManufacturingOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteManufacturingOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
