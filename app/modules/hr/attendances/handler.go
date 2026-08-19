package attendances

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateAttendance godoc
// @Summary Create a new Attendance
// @Description Create a new Attendance in the system
// @Tags hr-attendances
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
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
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/attendances [get]
// @Security BearerAuth
func GetAllAttendanceHandler(c echo.Context) error {
	data, err := GetAllAttendanceService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAttendanceByID godoc
// @Summary Get a Attendance by ID
// @Description Retrieve a specific Attendance by its ID
// @Tags hr-attendances
// @Produce json
// @Param id path int true "Attendance ID"
// @Success 200 {object} map[string]interface{}
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

func CreateOvertimeHandler(c echo.Context) error {
	var data Overtime
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateOvertimeService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllOvertimeHandler(c echo.Context) error {
	data, err := GetAllOvertimeService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetOvertimeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateOvertimeService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteOvertimeService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
