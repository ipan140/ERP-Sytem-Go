package time_off

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateLeaveRequest godoc
// @Summary Create a new LeaveRequest
// @Description Create a new LeaveRequest in the system
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/time_off [post]
// @Security BearerAuth
func CreateLeaveRequestHandler(c echo.Context) error {
	var data LeaveRequest
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateLeaveRequestService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllLeaveRequest godoc
// @Summary Get all LeaveRequest
// @Description Retrieve a list of all LeaveRequest
// @Tags hr-time_off
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off [get]
// @Security BearerAuth
func GetAllLeaveRequestHandler(c echo.Context) error {
	data, err := GetAllLeaveRequestService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetLeaveRequestByID godoc
// @Summary Get a LeaveRequest by ID
// @Description Retrieve a specific LeaveRequest by its ID
// @Tags hr-time_off
// @Produce json
// @Param id path int true "LeaveRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/{id} [get]
// @Security BearerAuth
func GetLeaveRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateLeaveRequest godoc
// @Summary Update a LeaveRequest
// @Description Update an existing LeaveRequest
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Param id path int true "LeaveRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/{id} [put]
// @Security BearerAuth
func UpdateLeaveRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateLeaveRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteLeaveRequest godoc
// @Summary Delete a LeaveRequest
// @Description Delete a LeaveRequest by ID
// @Tags hr-time_off
// @Produce json
// @Param id path int true "LeaveRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/{id} [delete]
// @Security BearerAuth
func DeleteLeaveRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeaveRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func CreateLeaveTypeHandler(c echo.Context) error {
	var data LeaveType
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateLeaveTypeService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllLeaveTypeHandler(c echo.Context) error {
	data, err := GetAllLeaveTypeService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLeaveTypeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveTypeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateLeaveTypeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveTypeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateLeaveTypeService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteLeaveTypeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeaveTypeService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateLeaveAllocationHandler(c echo.Context) error {
	var data LeaveAllocation
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateLeaveAllocationService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllLeaveAllocationHandler(c echo.Context) error {
	data, err := GetAllLeaveAllocationService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLeaveAllocationByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveAllocationByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateLeaveAllocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveAllocationByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateLeaveAllocationService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteLeaveAllocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeaveAllocationService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
