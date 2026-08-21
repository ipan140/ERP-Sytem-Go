package appraisals

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAppraisal godoc
// @Summary Create a new Appraisal
// @Description Create a new Appraisal in the system
// @Tags hr-appraisals
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/appraisals [post]
// @Security BearerAuth
func CreateAppraisalHandler(c echo.Context) error {
	var data Appraisal
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAppraisalService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAppraisal godoc
// @Summary Get all Appraisal
// @Description Retrieve a list of all Appraisal
// @Tags hr-appraisals
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/appraisals [get]
// @Security BearerAuth
func GetAllAppraisalHandler(c echo.Context) error {
	data, err := GetAllAppraisalService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAppraisalByID godoc
// @Summary Get a Appraisal by ID
// @Description Retrieve a specific Appraisal by its ID
// @Tags hr-appraisals
// @Produce json
// @Param id path int true "Appraisal ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/appraisals/{id} [get]
// @Security BearerAuth
func GetAppraisalByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAppraisalByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAppraisal godoc
// @Summary Update a Appraisal
// @Description Update an existing Appraisal
// @Tags hr-appraisals
// @Accept json
// @Produce json
// @Param id path int true "Appraisal ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/appraisals/{id} [put]
// @Security BearerAuth
func UpdateAppraisalHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAppraisalByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAppraisalService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAppraisal godoc
// @Summary Delete a Appraisal
// @Description Delete a Appraisal by ID
// @Tags hr-appraisals
// @Produce json
// @Param id path int true "Appraisal ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/appraisals/{id} [delete]
// @Security BearerAuth
func DeleteAppraisalHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAppraisalService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
