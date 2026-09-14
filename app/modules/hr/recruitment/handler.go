package recruitment

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateApplicant godoc
// @Summary Create a new Applicant
// @Description Create a new Applicant in the system
// @Tags hr-recruitment
// @Accept json
// @Produce json
// @Success 201 {object} Applicant
// @Param request body Applicant true "Payload"
// @Router /api/hr/recruitment [post]
// @Security BearerAuth
func CreateApplicantHandler(c echo.Context) error {
	var data Applicant
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateApplicantService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllApplicant godoc
// @Summary Get all Applicant
// @Description Retrieve a list of all Applicant
// @Tags hr-recruitment
// @Produce json
// @Success 200 {object} []Applicant
// @Router /api/hr/recruitment [get]
// @Security BearerAuth
func GetAllApplicantHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllApplicantService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	jobPositionID := c.QueryParam("job_position_id")
	stageID := c.QueryParam("stage_id")
	state := c.QueryParam("state")

	data, total, err := GetPaginatedApplicantService(offset, limit, search, jobPositionID, stageID, state)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetApplicantByID godoc
// @Summary Get a Applicant by ID
// @Description Retrieve a specific Applicant by its ID
// @Tags hr-recruitment
// @Produce json
// @Param id path int true "Applicant ID"
// @Success 200 {object} Applicant
// @Router /api/hr/recruitment/{id} [get]
// @Security BearerAuth
func GetApplicantByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetApplicantByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateApplicant godoc
// @Summary Update a Applicant
// @Description Update an existing Applicant
// @Tags hr-recruitment
// @Accept json
// @Produce json
// @Param id path int true "Applicant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment/{id} [put]
// @Security BearerAuth
func UpdateApplicantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetApplicantByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateApplicantService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteApplicant godoc
// @Summary Delete a Applicant
// @Description Delete a Applicant by ID
// @Tags hr-recruitment
// @Produce json
// @Param id path int true "Applicant ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/recruitment/{id} [delete]
// @Security BearerAuth
func DeleteApplicantHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteApplicantService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


