package maintenance

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetMaintenanceSummaryHandler godoc
// @Summary Get Maintenance executive KPI summary
// @Tags supply_chain-maintenance
// @Produce json
// @Success 200 {object} MaintenanceSummary
// @Router /api/supply_chain/maintenance/summary [get]
// @Security BearerAuth
func GetMaintenanceSummaryHandler(c echo.Context) error {
	summary, err := GetMaintenanceSummaryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve maintenance summary", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Maintenance summary retrieved successfully", summary)
}

// GetAllMaintenanceRequestHandler godoc
// @Summary Get paginated Maintenance Requests
// @Tags supply_chain-maintenance
// @Produce json
// @Success 200 {object} []MaintenanceRequest
// @Router /api/supply_chain/maintenance [get]
// @Security BearerAuth
func GetAllMaintenanceRequestHandler(c echo.Context) error {
	page, limit, _, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")
	reqType := c.QueryParam("type")

	list, total, err := GetPaginatedMaintenanceRequestsService(page, limit, search, state, reqType)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve maintenance requests", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Maintenance requests retrieved successfully", list, meta)
}

// CreateMaintenanceRequestHandler godoc
// @Summary Create a new Maintenance Request
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Success 201 {object} MaintenanceRequest
// @Router /api/supply_chain/maintenance [post]
// @Security BearerAuth
func CreateMaintenanceRequestHandler(c echo.Context) error {
	var req CreateMaintenanceRequestDto
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if req.Name == "" || req.EquipmentID == 0 {
		return utils.SendError(c, http.StatusBadRequest, "Subject name and equipment are required", "")
	}

	created, err := CreateMaintenanceRequestWithSequenceService(&req)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create maintenance request", err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "Maintenance request created successfully", created)
}

// GetMaintenanceRequestByIDHandler godoc
// @Summary Get a MaintenanceRequest by ID
// @Tags supply_chain-maintenance
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} MaintenanceRequest
// @Router /api/supply_chain/maintenance/{id} [get]
// @Security BearerAuth
func GetMaintenanceRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMaintenanceRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Maintenance request not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Maintenance request retrieved successfully", data)
}

// UpdateMaintenanceStateHandler godoc
// @Summary Update maintenance request workflow state (todo -> progress -> done / cancel)
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} MaintenanceRequest
// @Router /api/supply_chain/maintenance/{id}/state [put]
// @Security BearerAuth
func UpdateMaintenanceStateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req UpdateMaintenanceStateRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if req.State != "todo" && req.State != "progress" && req.State != "done" && req.State != "cancel" {
		return utils.SendError(c, http.StatusBadRequest, "Invalid state. Must be todo, progress, done, or cancel", "")
	}

	updated, err := UpdateMaintenanceStateService(uint(id), req.State, req.Duration, req.Notes)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update maintenance state", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Maintenance state updated successfully", updated)
}

// UpdateMaintenanceRequestHandler godoc
// @Summary Update a MaintenanceRequest
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} MaintenanceRequest
// @Router /api/supply_chain/maintenance/{id} [put]
// @Security BearerAuth
func UpdateMaintenanceRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMaintenanceRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Maintenance request not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateMaintenanceRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Maintenance request updated successfully", data)
}

// DeleteMaintenanceRequestHandler godoc
// @Summary Delete a MaintenanceRequest
// @Tags supply_chain-maintenance
// @Produce json
// @Param id path int true "MaintenanceRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance/{id} [delete]
// @Security BearerAuth
func DeleteMaintenanceRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMaintenanceRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete maintenance request", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Maintenance request deleted successfully", nil)
}

// GetMaintenanceEquipmentsHandler godoc
// @Summary Get paginated Maintenance Equipments
// @Tags supply_chain-maintenance
// @Produce json
// @Success 200 {object} []MaintenanceEquipment
// @Router /api/supply_chain/maintenance/equipment [get]
// @Security BearerAuth
func GetMaintenanceEquipmentsHandler(c echo.Context) error {
	page, limit, _, search := utils.GetPaginationQuery(c)
	list, total, err := GetPaginatedEquipmentsService(page, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve equipments", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Equipments retrieved successfully", list, meta)
}

// CreateMaintenanceEquipmentHandler godoc
// @Summary Create a Maintenance Equipment
// @Tags supply_chain-maintenance
// @Accept json
// @Produce json
// @Success 201 {object} MaintenanceEquipment
// @Router /api/supply_chain/maintenance/equipment [post]
// @Security BearerAuth
func CreateMaintenanceEquipmentHandler(c echo.Context) error {
	var dto CreateMaintenanceEquipmentDto
	if err := c.Bind(&dto); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}

	if dto.Name == "" {
		return utils.SendError(c, http.StatusBadRequest, "Equipment name is required", "")
	}

	created, err := CreateMaintenanceEquipmentService(&dto)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create equipment", err.Error())
	}

	return utils.SendSuccess(c, http.StatusCreated, "Equipment created successfully", created)
}

// DeleteMaintenanceEquipmentHandler godoc
// @Summary Delete Maintenance Equipment
// @Tags supply_chain-maintenance
// @Produce json
// @Param id path int true "Equipment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/maintenance/equipment/{id} [delete]
// @Security BearerAuth
func DeleteMaintenanceEquipmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMaintenanceEquipmentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete equipment", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Equipment deleted successfully", nil)
}


