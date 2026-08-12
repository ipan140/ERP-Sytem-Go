package rental

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateRentalOrder godoc
// @Summary Create a new RentalOrder
// @Description Create a new RentalOrder in the system
// @Tags sales-rental
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/rental [post]
// @Security BearerAuth
func CreateRentalOrderHandler(c echo.Context) error {
	var data RentalOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateRentalOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllRentalOrder godoc
// @Summary Get all RentalOrder
// @Description Retrieve a list of all RentalOrder
// @Tags sales-rental
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental [get]
// @Security BearerAuth
func GetAllRentalOrderHandler(c echo.Context) error {
	data, err := GetAllRentalOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetRentalOrderByID godoc
// @Summary Get a RentalOrder by ID
// @Description Retrieve a specific RentalOrder by its ID
// @Tags sales-rental
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/{id} [get]
// @Security BearerAuth
func GetRentalOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateRentalOrder godoc
// @Summary Update a RentalOrder
// @Description Update an existing RentalOrder
// @Tags sales-rental
// @Accept json
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/{id} [put]
// @Security BearerAuth
func UpdateRentalOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRentalOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateRentalOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteRentalOrder godoc
// @Summary Delete a RentalOrder
// @Description Delete a RentalOrder by ID
// @Tags sales-rental
// @Produce json
// @Param id path int true "RentalOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/rental/{id} [delete]
// @Security BearerAuth
func DeleteRentalOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteRentalOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
