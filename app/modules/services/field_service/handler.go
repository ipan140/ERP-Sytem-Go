package field_service

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateFieldServiceTask godoc
// @Summary Create a new FieldServiceTask
// @Description Create a new FieldServiceTask in the system
// @Tags services-field_service
// @Accept json
// @Produce json
// @Success 201 {object} FieldServiceTask
// @Param request body FieldServiceTask true "Payload"
// @Router /api/services/field_service [post]
// @Security BearerAuth
func CreateFieldServiceTaskHandler(c echo.Context) error {
	var data FieldServiceTask
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateFieldServiceTaskService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllFieldServiceTask godoc
// @Summary Get all FieldServiceTask
// @Description Retrieve a list of all FieldServiceTask
// @Tags services-field_service
// @Produce json
// @Success 200 {object} []FieldServiceTask
// @Router /api/services/field_service [get]
// @Security BearerAuth
func GetAllFieldServiceTaskHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllFieldServiceTaskService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")
	priority := c.QueryParam("priority")
	employeeID, _ := strconv.Atoi(c.QueryParam("employee_id"))
	companyID, _ := strconv.Atoi(c.QueryParam("company_id"))

	userRole, _ := c.Get("role").(string)
	if c.QueryParam("my_only") == "true" && (userRole == "staff" || userRole == "technician") {
		if currentEmpID, _ := strconv.Atoi(c.QueryParam("my_employee_id")); currentEmpID > 0 {
			employeeID = currentEmpID
		}
	}

	data, total, err := GetPaginatedFieldServiceTaskService(offset, limit, search, state, priority, uint(employeeID), uint(companyID))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetFieldServiceTaskByID godoc
// @Summary Get a FieldServiceTask by ID
// @Description Retrieve a specific FieldServiceTask by its ID
// @Tags services-field_service
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} FieldServiceTask
// @Router /api/services/field_service/{id} [get]
// @Security BearerAuth
func GetFieldServiceTaskByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFieldServiceTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateFieldServiceTask godoc
// @Summary Update a FieldServiceTask
// @Description Update an existing FieldServiceTask
// @Tags services-field_service
// @Accept json
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service/{id} [put]
// @Security BearerAuth
func UpdateFieldServiceTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFieldServiceTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateFieldServiceTaskService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteFieldServiceTask godoc
// @Summary Delete a FieldServiceTask
// @Description Delete a FieldServiceTask by ID
// @Tags services-field_service
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service/{id} [delete]
// @Security BearerAuth
func DeleteFieldServiceTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFieldServiceTaskService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// Fase 2: e-BAST Validation Gate Handler
func ValidateBastHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	validatorID := uint(1)
	if uid, ok := c.Get("user_id").(uint); ok && uid > 0 {
		validatorID = uid
	} else if uidFloat, ok := c.Get("user_id").(float64); ok && uidFloat > 0 {
		validatorID = uint(uidFloat)
	}

	if err := ValidateBastFieldServiceTaskService(uint(id), validatorID); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memvalidasi e-BAST pekerjaan", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Dokumen e-BAST telah diverifikasi dan tugas resmi diselesaikan", nil)
}

// GPSCheckInHandler godoc
// @Summary Record technician GPS check-in at client site
// @Tags services-field_service
// @Router /api/services/field_service/{id}/check-in [post]
func GPSCheckInHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var req struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Koordinat GPS tidak valid", err.Error())
	}

	task, err := RecordGPSCheckInService(uint(id), req.Lat, req.Lng)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mencatat Check-in GPS", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Berhasil Check-in GPS di lokasi klien", task)
}

// GPSCheckOutHandler godoc
// @Summary Record technician GPS check-out
// @Tags services-field_service
// @Router /api/services/field_service/{id}/check-out [post]
func GPSCheckOutHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", err.Error())
	}

	var req struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Koordinat GPS tidak valid", err.Error())
	}

	task, err := RecordGPSCheckOutService(uint(id), req.Lat, req.Lng)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mencatat Check-out GPS", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Berhasil Check-out GPS pekerjaan", task)
}



