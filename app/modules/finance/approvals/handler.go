package approvals

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateApprovalRequest godoc
// @Summary Create a new ApprovalRequest
// @Description Create a new ApprovalRequest in the system
// @Tags finance-approvals
// @Accept json
// @Produce json
// @Success 201 {object} ApprovalRequest
// @Param request body ApprovalRequest true "Payload"
// @Router /api/finance/approvals [post]
// @Security BearerAuth
func CreateApprovalRequestHandler(c echo.Context) error {
	var data ApprovalRequest
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateApprovalRequestService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllApprovalRequest godoc
// @Summary Get all ApprovalRequest
// @Description Retrieve a list of all ApprovalRequest
// @Tags finance-approvals
// @Produce json
// @Success 200 {object} []ApprovalRequest
// @Router /api/finance/approvals [get]
// @Security BearerAuth
func GetAllApprovalRequestHandler(c echo.Context) error {
	data, err := GetAllApprovalRequestService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetApprovalRequestByID godoc
// @Summary Get a ApprovalRequest by ID
// @Description Retrieve a specific ApprovalRequest by its ID
// @Tags finance-approvals
// @Produce json
// @Param id path int true "ApprovalRequest ID"
// @Success 200 {object} ApprovalRequest
// @Router /api/finance/approvals/{id} [get]
// @Security BearerAuth
func GetApprovalRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetApprovalRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateApprovalRequest godoc
// @Summary Update a ApprovalRequest
// @Description Update an existing ApprovalRequest
// @Tags finance-approvals
// @Accept json
// @Produce json
// @Param id path int true "ApprovalRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/approvals/{id} [put]
// @Security BearerAuth
func UpdateApprovalRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetApprovalRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateApprovalRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteApprovalRequest godoc
// @Summary Delete a ApprovalRequest
// @Description Delete a ApprovalRequest by ID
// @Tags finance-approvals
// @Produce json
// @Param id path int true "ApprovalRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/approvals/{id} [delete]
// @Security BearerAuth
func DeleteApprovalRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteApprovalRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


