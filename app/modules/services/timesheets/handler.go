package timesheets

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateTimesheet godoc
// @Summary Create a new Timesheet
// @Description Create a new Timesheet in the system
// @Tags services-timesheets
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/timesheets [post]
// @Security BearerAuth
func CreateTimesheetHandler(c echo.Context) error {
	var data Timesheet
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateTimesheetService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllTimesheet godoc
// @Summary Get all Timesheet
// @Description Retrieve a list of all Timesheet
// @Tags services-timesheets
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/timesheets [get]
// @Security BearerAuth
func GetAllTimesheetHandler(c echo.Context) error {
	data, err := GetAllTimesheetService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetTimesheetByID godoc
// @Summary Get a Timesheet by ID
// @Description Retrieve a specific Timesheet by its ID
// @Tags services-timesheets
// @Produce json
// @Param id path int true "Timesheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/timesheets/{id} [get]
// @Security BearerAuth
func GetTimesheetByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTimesheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateTimesheet godoc
// @Summary Update a Timesheet
// @Description Update an existing Timesheet
// @Tags services-timesheets
// @Accept json
// @Produce json
// @Param id path int true "Timesheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/timesheets/{id} [put]
// @Security BearerAuth
func UpdateTimesheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTimesheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateTimesheetService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteTimesheet godoc
// @Summary Delete a Timesheet
// @Description Delete a Timesheet by ID
// @Tags services-timesheets
// @Produce json
// @Param id path int true "Timesheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/timesheets/{id} [delete]
// @Security BearerAuth
func DeleteTimesheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTimesheetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
