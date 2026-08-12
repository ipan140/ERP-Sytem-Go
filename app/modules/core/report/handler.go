package report

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateReport godoc
// @Summary Create a new Report
// @Description Create a new Report in the system
// @Tags core-report
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/core/report [post]
// @Security BearerAuth
func CreateReportHandler(c echo.Context) error {
	var data Report
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateReportService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllReport godoc
// @Summary Get all Report
// @Description Retrieve a list of all Report
// @Tags core-report
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/core/report [get]
// @Security BearerAuth
func GetAllReportHandler(c echo.Context) error {
	data, err := GetAllReportService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetReportByID godoc
// @Summary Get a Report by ID
// @Description Retrieve a specific Report by its ID
// @Tags core-report
// @Produce json
// @Param id path int true "Report ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/report/{id} [get]
// @Security BearerAuth
func GetReportByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReportByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateReport godoc
// @Summary Update a Report
// @Description Update an existing Report
// @Tags core-report
// @Accept json
// @Produce json
// @Param id path int true "Report ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/report/{id} [put]
// @Security BearerAuth
func UpdateReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReportByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateReportService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteReport godoc
// @Summary Delete a Report
// @Description Delete a Report by ID
// @Tags core-report
// @Produce json
// @Param id path int true "Report ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/report/{id} [delete]
// @Security BearerAuth
func DeleteReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteReportService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
