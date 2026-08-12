package fleet

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateFleetVehicle godoc
// @Summary Create a new FleetVehicle
// @Description Create a new FleetVehicle in the system
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet [post]
// @Security BearerAuth
func CreateFleetVehicleHandler(c echo.Context) error {
	var data FleetVehicle
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateFleetVehicleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllFleetVehicle godoc
// @Summary Get all FleetVehicle
// @Description Retrieve a list of all FleetVehicle
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet [get]
// @Security BearerAuth
func GetAllFleetVehicleHandler(c echo.Context) error {
	data, err := GetAllFleetVehicleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetFleetVehicleByID godoc
// @Summary Get a FleetVehicle by ID
// @Description Retrieve a specific FleetVehicle by its ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "FleetVehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [get]
// @Security BearerAuth
func GetFleetVehicleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFleetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateFleetVehicle godoc
// @Summary Update a FleetVehicle
// @Description Update an existing FleetVehicle
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "FleetVehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [put]
// @Security BearerAuth
func UpdateFleetVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFleetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateFleetVehicleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteFleetVehicle godoc
// @Summary Delete a FleetVehicle
// @Description Delete a FleetVehicle by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "FleetVehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [delete]
// @Security BearerAuth
func DeleteFleetVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFleetVehicleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
