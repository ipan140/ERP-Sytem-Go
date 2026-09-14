package attendances

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAttendance godoc
// @Summary Create a new Attendance
// @Description Create a new Attendance in the system
// @Tags hr-attendances
// @Accept json
// @Produce json
// @Success 201 {object} Attendance
// @Param request body Attendance true "Payload"
// @Router /api/hr/attendances [post]
// @Security BearerAuth
func CreateAttendanceHandler(c echo.Context) error {
	var data Attendance
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAttendanceService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAttendance godoc
// @Summary Get all Attendance
// @Description Retrieve a list of all Attendance
// @Tags hr-attendances
// @Produce json
// @Success 200 {object} []Attendance
// @Router /api/hr/attendances [get]
// @Security BearerAuth
func GetAllAttendanceHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllAttendanceService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	employeeID := c.QueryParam("employee_id")
	date := c.QueryParam("date")

	data, total, err := GetPaginatedAttendanceService(offset, limit, search, employeeID, date)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetAttendanceByID godoc
// @Summary Get a Attendance by ID
// @Description Retrieve a specific Attendance by its ID
// @Tags hr-attendances
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} Attendance
// @Router /api/hr/attendances/{id} [get]
// @Security BearerAuth
func GetAttendanceByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAttendanceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAttendance godoc
// @Summary Update a Attendance
// @Description Update an existing Attendance
// @Tags hr-attendances
// @Accept json
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/attendances/{id} [put]
// @Security BearerAuth
func UpdateAttendanceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAttendanceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAttendanceService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAttendance godoc
// @Summary Delete a Attendance
// @Description Delete a Attendance by ID
// @Tags hr-attendances
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/attendances/{id} [delete]
// @Security BearerAuth
func DeleteAttendanceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAttendanceService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create Overtime
// @Description Create a new Overtime
// @Tags hr-attendances
// @Accept json
// @Produce json
// @Success 201 {object} Overtime
// @Param request body Overtime true "Payload"
// @Router /api/hr/attendances/overtime [post]
// @Security BearerAuth
func CreateOvertimeHandler(c echo.Context) error {
	var data Overtime
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateOvertimeService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Overtime
// @Description Retrieve a list of all Overtime
// @Tags hr-attendances
// @Produce json
// @Success 200 {object} Overtime
// @Router /api/hr/attendances/overtime [get]
// @Security BearerAuth
func GetAllOvertimeHandler(c echo.Context) error {
	data, err := GetAllOvertimeService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetOvertimeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Overtime
// @Description Update an existing Overtime
// @Tags hr-attendances
// @Accept json
// @Produce json
// @Param id path int true "Overtime ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/attendances/overtime/{id} [put]
// @Security BearerAuth
func UpdateOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateOvertimeService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Overtime
// @Description Delete Overtime by ID
// @Tags hr-attendances
// @Produce json
// @Param id path int true "Overtime ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/attendances/overtime/{id} [delete]
// @Security BearerAuth
func DeleteOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteOvertimeService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


