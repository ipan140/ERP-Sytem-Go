package recruitment

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateJobApplicant godoc
// @Summary Create a new JobApplicant
// @Description Create a new JobApplicant in the system
// @Tags hr-recruitment
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/recruitment [post]
// @Security BearerAuth
func CreateJobApplicantHandler(c echo.Context) error {
	var data JobApplicant
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateJobApplicantService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllJobApplicant godoc
// @Summary Get all JobApplicant
// @Description Retrieve a list of all JobApplicant
// @Tags hr-recruitment
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment [get]
// @Security BearerAuth
func GetAllJobApplicantHandler(c echo.Context) error {
	data, err := GetAllJobApplicantService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetJobApplicantByID godoc
// @Summary Get a JobApplicant by ID
// @Description Retrieve a specific JobApplicant by its ID
// @Tags hr-recruitment
// @Produce json
// @Param id path int true "JobApplicant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment/{id} [get]
// @Security BearerAuth
func GetJobApplicantByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobApplicantByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateJobApplicant godoc
// @Summary Update a JobApplicant
// @Description Update an existing JobApplicant
// @Tags hr-recruitment
// @Accept json
// @Produce json
// @Param id path int true "JobApplicant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment/{id} [put]
// @Security BearerAuth
func UpdateJobApplicantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobApplicantByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateJobApplicantService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteJobApplicant godoc
// @Summary Delete a JobApplicant
// @Description Delete a JobApplicant by ID
// @Tags hr-recruitment
// @Produce json
// @Param id path int true "JobApplicant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment/{id} [delete]
// @Security BearerAuth
func DeleteJobApplicantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJobApplicantService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
