package maintenance

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateMaintenanceRequest godoc
// @Summary Create a new MaintenanceRequest
// @Description Create a new MaintenanceRequest in the system
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance [post]
// @Security BearerAuth
func CreateMaintenanceRequestHandler(c echo.Context) error {
	var data MaintenanceRequest
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateMaintenanceRequestService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllMaintenanceRequest godoc
// @Summary Get all MaintenanceRequest
// @Description Retrieve a list of all MaintenanceRequest
// @Tags supply_chain-maintenance
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance [get]
// @Security BearerAuth
func GetAllMaintenanceRequestHandler(c echo.Context) error {
	data, err := GetAllMaintenanceRequestService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetMaintenanceRequestByID godoc
// @Summary Get a MaintenanceRequest by ID
// @Description Retrieve a specific MaintenanceRequest by its ID
// @Tags supply_chain-maintenance
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance/{id} [get]
// @Security BearerAuth
func GetMaintenanceRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMaintenanceRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateMaintenanceRequest godoc
// @Summary Update a MaintenanceRequest
// @Description Update an existing MaintenanceRequest
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance/{id} [put]
// @Security BearerAuth
func UpdateMaintenanceRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMaintenanceRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateMaintenanceRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteMaintenanceRequest godoc
// @Summary Delete a MaintenanceRequest
// @Description Delete a MaintenanceRequest by ID
// @Tags supply_chain-maintenance
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance/{id} [delete]
// @Security BearerAuth
func DeleteMaintenanceRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMaintenanceRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
