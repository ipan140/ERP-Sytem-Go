package planning

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateShift godoc
// @Summary Create a new Shift
// @Description Create a new Shift in the system
// @Tags services-planning
// @Accept json
// @Produce json
// @Success 201 {object} Shift
// @Param request body Shift true "Payload"
// @Router /api/services/planning [post]
// @Security BearerAuth
func CreateShiftHandler(c echo.Context) error {
	var data Shift
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateShiftService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllShift godoc
// @Summary Get all Shift
// @Description Retrieve a list of all Shift
// @Tags services-planning
// @Produce json
// @Success 200 {object} []Shift
// @Router /api/services/planning [get]
// @Security BearerAuth
func GetAllShiftHandler(c echo.Context) error {
	data, err := GetAllShiftService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetShiftByID godoc
// @Summary Get a Shift by ID
// @Description Retrieve a specific Shift by its ID
// @Tags services-planning
// @Produce json
// @Param id path int true "Shift ID"
// @Success 200 {object} Shift
// @Router /api/services/planning/{id} [get]
// @Security BearerAuth
func GetShiftByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetShiftByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateShift godoc
// @Summary Update a Shift
// @Description Update an existing Shift
// @Tags services-planning
// @Accept json
// @Produce json
// @Param id path int true "Shift ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/planning/{id} [put]
// @Security BearerAuth
func UpdateShiftHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetShiftByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateShiftService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteShift godoc
// @Summary Delete a Shift
// @Description Delete a Shift by ID
// @Tags services-planning
// @Produce json
// @Param id path int true "Shift ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/planning/{id} [delete]
// @Security BearerAuth
func DeleteShiftHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteShiftService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


