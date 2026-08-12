package repairs

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateRepairOrder godoc
// @Summary Create a new RepairOrder
// @Description Create a new RepairOrder in the system
// @Tags services-repairs
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/repairs [post]
// @Security BearerAuth
func CreateRepairOrderHandler(c echo.Context) error {
	var data RepairOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateRepairOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllRepairOrder godoc
// @Summary Get all RepairOrder
// @Description Retrieve a list of all RepairOrder
// @Tags services-repairs
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs [get]
// @Security BearerAuth
func GetAllRepairOrderHandler(c echo.Context) error {
	data, err := GetAllRepairOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetRepairOrderByID godoc
// @Summary Get a RepairOrder by ID
// @Description Retrieve a specific RepairOrder by its ID
// @Tags services-repairs
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs/{id} [get]
// @Security BearerAuth
func GetRepairOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRepairOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateRepairOrder godoc
// @Summary Update a RepairOrder
// @Description Update an existing RepairOrder
// @Tags services-repairs
// @Accept json
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs/{id} [put]
// @Security BearerAuth
func UpdateRepairOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetRepairOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateRepairOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteRepairOrder godoc
// @Summary Delete a RepairOrder
// @Description Delete a RepairOrder by ID
// @Tags services-repairs
// @Produce json
// @Param id path int true "RepairOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/repairs/{id} [delete]
// @Security BearerAuth
func DeleteRepairOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteRepairOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
