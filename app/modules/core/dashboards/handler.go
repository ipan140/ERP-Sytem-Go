package dashboards

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateDashboard godoc
// @Summary Create a new Dashboard
// @Description Create a new Dashboard in the system
// @Tags dashboards
// @Accept json
// @Produce json
// @Success 201 {object} Dashboard
// @Param request body Dashboard true "Payload"
// @Router /api/dashboards [post]
// @Security BearerAuth
func CreateDashboardHandler(c echo.Context) error {
	var data Dashboard
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateDashboardService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllDashboard godoc
// @Summary Get all Dashboard
// @Description Retrieve a list of all Dashboard
// @Tags dashboards
// @Produce json
// @Success 200 {object} []Dashboard
// @Router /api/dashboards [get]
// @Security BearerAuth
func GetAllDashboardHandler(c echo.Context) error {
	data, err := GetAllDashboardService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetDashboardByID godoc
// @Summary Get a Dashboard by ID
// @Description Retrieve a specific Dashboard by its ID
// @Tags dashboards
// @Produce json
// @Param id path int true "Dashboard ID"
// @Success 200 {object} Dashboard
// @Router /api/dashboards/{id} [get]
// @Security BearerAuth
func GetDashboardByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDashboardByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateDashboard godoc
// @Summary Update a Dashboard
// @Description Update an existing Dashboard
// @Tags dashboards
// @Accept json
// @Produce json
// @Param id path int true "Dashboard ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/dashboards/{id} [put]
// @Security BearerAuth
func UpdateDashboardHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDashboardByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateDashboardService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteDashboard godoc
// @Summary Delete a Dashboard
// @Description Delete a Dashboard by ID
// @Tags dashboards
// @Produce json
// @Param id path int true "Dashboard ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/dashboards/{id} [delete]
// @Security BearerAuth
func DeleteDashboardHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteDashboardService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


