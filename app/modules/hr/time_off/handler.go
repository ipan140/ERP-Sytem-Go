package time_off

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateLeaveRequest godoc
// @Summary Create a new LeaveRequest
// @Description Create a new LeaveRequest in the system
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Success 201 {object} LeaveRequest
// @Param request body LeaveRequest true "Payload"
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
// @Success 200 {object} []LeaveRequest
// @Router /api/hr/time_off [get]
// @Security BearerAuth
func GetAllLeaveRequestHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllLeaveRequestService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	employeeID := c.QueryParam("employee_id")
	status := c.QueryParam("status")

	data, total, err := GetPaginatedLeaveRequestService(offset, limit, search, employeeID, status)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetLeaveRequestByID godoc
// @Summary Get a LeaveRequest by ID
// @Description Retrieve a specific LeaveRequest by its ID
// @Tags hr-time_off
// @Produce json
// @Param id path int true "LeaveRequest ID"
// @Success 200 {object} LeaveRequest
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

// @Summary Create LeaveType
// @Description Create a new LeaveType
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Success 201 {object} LeaveType
// @Param request body LeaveType true "Payload"
// @Router /api/hr/time_off/leavetype [post]
// @Security BearerAuth
func CreateLeaveTypeHandler(c echo.Context) error {
	var data LeaveType
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateLeaveTypeService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all LeaveType
// @Description Retrieve a list of all LeaveType
// @Tags hr-time_off
// @Produce json
// @Success 200 {object} LeaveType
// @Router /api/hr/time_off/leavetype [get]
// @Security BearerAuth
func GetAllLeaveTypeHandler(c echo.Context) error {
	data, err := GetAllLeaveTypeService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLeaveTypeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveTypeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update LeaveType
// @Description Update an existing LeaveType
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Param id path int true "LeaveType ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/leavetype/{id} [put]
// @Security BearerAuth
func UpdateLeaveTypeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveTypeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateLeaveTypeService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete LeaveType
// @Description Delete LeaveType by ID
// @Tags hr-time_off
// @Produce json
// @Param id path int true "LeaveType ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/leavetype/{id} [delete]
// @Security BearerAuth
func DeleteLeaveTypeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeaveTypeService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create LeaveAllocation
// @Description Create a new LeaveAllocation
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Success 201 {object} LeaveAllocation
// @Param request body LeaveAllocation true "Payload"
// @Router /api/hr/time_off/leaveallocation [post]
// @Security BearerAuth
func CreateLeaveAllocationHandler(c echo.Context) error {
	var data LeaveAllocation
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateLeaveAllocationService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all LeaveAllocation
// @Description Retrieve a list of all LeaveAllocation
// @Tags hr-time_off
// @Produce json
// @Success 200 {object} LeaveAllocation
// @Router /api/hr/time_off/leaveallocation [get]
// @Security BearerAuth
func GetAllLeaveAllocationHandler(c echo.Context) error {
	data, err := GetAllLeaveAllocationService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetLeaveAllocationByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveAllocationByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update LeaveAllocation
// @Description Update an existing LeaveAllocation
// @Tags hr-time_off
// @Accept json
// @Produce json
// @Param id path int true "LeaveAllocation ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/leaveallocation/{id} [put]
// @Security BearerAuth
func UpdateLeaveAllocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetLeaveAllocationByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateLeaveAllocationService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete LeaveAllocation
// @Description Delete LeaveAllocation by ID
// @Tags hr-time_off
// @Produce json
// @Param id path int true "LeaveAllocation ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/time_off/leaveallocation/{id} [delete]
// @Security BearerAuth
func DeleteLeaveAllocationHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteLeaveAllocationService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


