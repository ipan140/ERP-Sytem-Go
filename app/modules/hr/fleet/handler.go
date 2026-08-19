package fleet

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateVehicle godoc
// @Summary Create a new Vehicle
// @Description Create a new Vehicle in the system
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/fleet [post]
// @Security BearerAuth
func CreateVehicleHandler(c echo.Context) error {
	var data Vehicle
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateVehicleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllVehicle godoc
// @Summary Get all Vehicle
// @Description Retrieve a list of all Vehicle
// @Tags hr-fleet
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet [get]
// @Security BearerAuth
func GetAllVehicleHandler(c echo.Context) error {
	data, err := GetAllVehicleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetVehicleByID godoc
// @Summary Get a Vehicle by ID
// @Description Retrieve a specific Vehicle by its ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [get]
// @Security BearerAuth
func GetVehicleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateVehicle godoc
// @Summary Update a Vehicle
// @Description Update an existing Vehicle
// @Tags hr-fleet
// @Accept json
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [put]
// @Security BearerAuth
func UpdateVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetVehicleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateVehicleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteVehicle godoc
// @Summary Delete a Vehicle
// @Description Delete a Vehicle by ID
// @Tags hr-fleet
// @Produce json
// @Param id path int true "Vehicle ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/fleet/{id} [delete]
// @Security BearerAuth
func DeleteVehicleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteVehicleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
