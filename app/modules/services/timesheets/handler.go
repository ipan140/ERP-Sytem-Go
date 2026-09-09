package timesheets

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateTimesheet godoc
// @Summary Create a new Timesheet
// @Description Create a new Timesheet in the system
// @Tags services-timesheets
// @Accept json
// @Produce json
// @Success 201 {object} Timesheet
// @Param request body Timesheet true "Payload"
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
// @Success 200 {object} []Timesheet
// @Router /api/services/timesheets [get]
// @Security BearerAuth
func GetAllTimesheetHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllTimesheetService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	projectID, _ := strconv.Atoi(c.QueryParam("project_id"))
	employeeID, _ := strconv.Atoi(c.QueryParam("employee_id"))
	billable := c.QueryParam("is_billable")
	status := c.QueryParam("status")
	companyID, _ := strconv.Atoi(c.QueryParam("company_id"))

	userRole, _ := c.Get("role").(string)
	if c.QueryParam("my_only") == "true" && (userRole == "staff" || userRole == "technician") {
		if currentEmpID, _ := strconv.Atoi(c.QueryParam("my_employee_id")); currentEmpID > 0 {
			employeeID = currentEmpID
		}
	}

	data, total, err := GetPaginatedTimesheetService(offset, limit, search, uint(projectID), uint(employeeID), billable, status, uint(companyID))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetTimesheetByID godoc
// @Summary Get a Timesheet by ID
// @Description Retrieve a specific Timesheet by its ID
// @Tags services-timesheets
// @Produce json
// @Param id path int true "Timesheet ID"
// @Success 200 {object} Timesheet
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

// Fase 2: Workflow Approval Handlers
func SubmitTimesheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := SubmitTimesheetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to submit timesheet", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Timesheet submitted for review", nil)
}

func ApproveTimesheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	approverID := uint(1)
	if uid, ok := c.Get("user_id").(uint); ok && uid > 0 {
		approverID = uid
	} else if uidFloat, ok := c.Get("user_id").(float64); ok && uidFloat > 0 {
		approverID = uint(uidFloat)
	}

	if err := ApproveTimesheetService(uint(id), approverID); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to approve timesheet", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Timesheet approved successfully", nil)
}

func RejectTimesheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Reason string `json:"reason"`
	}
	_ = c.Bind(&req)

	if err := RejectTimesheetService(uint(id), req.Reason); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to reject timesheet", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Timesheet rejected", nil)
}

func BulkApproveTimesheetsHandler(c echo.Context) error {
	var req struct {
		IDs []uint `json:"ids"`
	}
	if err := c.Bind(&req); err != nil || len(req.IDs) == 0 {
		return utils.SendError(c, http.StatusBadRequest, "Pilih setidaknya satu timesheet untuk disetujui", "")
	}

	approverID := uint(1)
	if uid, ok := c.Get("user_id").(uint); ok && uid > 0 {
		approverID = uid
	} else if uidFloat, ok := c.Get("user_id").(float64); ok && uidFloat > 0 {
		approverID = uint(uidFloat)
	}

	if err := BulkApproveTimesheetsService(req.IDs, approverID); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyetujui timesheet terpilih", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Seluruh timesheet terpilih berhasil disetujui", nil)
}


