package surveys

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateSurvey godoc
// @Summary Create a new Survey
// @Description Create a new Survey in the system
// @Tags marketing-surveys
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/marketing/surveys [post]
// @Security BearerAuth
func CreateSurveyHandler(c echo.Context) error {
	var data Survey
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSurveyService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSurvey godoc
// @Summary Get all Survey
// @Description Retrieve a list of all Survey
// @Tags marketing-surveys
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/surveys [get]
// @Security BearerAuth
func GetAllSurveyHandler(c echo.Context) error {
	data, err := GetAllSurveyService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSurveyByID godoc
// @Summary Get a Survey by ID
// @Description Retrieve a specific Survey by its ID
// @Tags marketing-surveys
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/surveys/{id} [get]
// @Security BearerAuth
func GetSurveyByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSurveyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSurvey godoc
// @Summary Update a Survey
// @Description Update an existing Survey
// @Tags marketing-surveys
// @Accept json
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/surveys/{id} [put]
// @Security BearerAuth
func UpdateSurveyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSurveyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSurveyService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSurvey godoc
// @Summary Delete a Survey
// @Description Delete a Survey by ID
// @Tags marketing-surveys
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/surveys/{id} [delete]
// @Security BearerAuth
func DeleteSurveyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSurveyService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
