package surveys

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSurvey godoc
// @Summary Create a new Survey
// @Description Create a new customer feedback or NPS survey in the system
// @Tags marketing-surveys
// @Accept json
// @Produce json
// @Param request body Survey true "Survey Payload"
// @Success 201 {object} utils.SuccessResponse{data=Survey} "Survey created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload"
// @Failure 500 {object} utils.ErrorResponse "Failed to create data"
// @Router /api/marketing/surveys [post]
// @Security BearerAuth
func CreateSurveyHandler(c echo.Context) error {
	var data Survey
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if uid, ok := c.Get("user_id").(uint); ok && uid > 0 {
		data.UserID = &uid
	} else if data.UserID != nil && *data.UserID == 0 {
		data.UserID = nil
	}
	if err := CreateSurveyService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSurvey godoc
// @Summary Get all Survey
// @Description Retrieve a list of all surveys with calculated NPS scores
// @Tags marketing-surveys
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]Survey} "List of surveys"
// @Failure 500 {object} utils.ErrorResponse "Failed to retrieve data"
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
// @Description Retrieve a specific survey by its ID
// @Tags marketing-surveys
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} utils.SuccessResponse{data=Survey} "Survey found"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 404 {object} utils.ErrorResponse "Data not found"
// @Router /api/marketing/surveys/{id} [get]
// @Security BearerAuth
func GetSurveyByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	data, err := GetSurveyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSurvey godoc
// @Summary Update a Survey
// @Description Update an existing survey details
// @Tags marketing-surveys
// @Accept json
// @Produce json
// @Param id path int true "Survey ID"
// @Param request body Survey true "Updated Survey Payload"
// @Success 200 {object} utils.SuccessResponse{data=Survey} "Survey updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or ID"
// @Failure 404 {object} utils.ErrorResponse "Data not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to update data"
// @Router /api/marketing/surveys/{id} [put]
// @Security BearerAuth
func UpdateSurveyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
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
// @Description Delete a survey by ID
// @Tags marketing-surveys
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} utils.SuccessResponse{data=nil} "Data deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 500 {object} utils.ErrorResponse "Failed to delete data"
// @Router /api/marketing/surveys/{id} [delete]
// @Security BearerAuth
func DeleteSurveyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	if err := DeleteSurveyService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// SurveyResponsePayload request body for submitting survey response
type SurveyResponsePayload struct {
	Rating   int    `json:"rating" example:"9"` // 0 - 10
	Feedback string `json:"feedback,omitempty" example:"Pelayanan sangat memuaskan!"`
}

// SubmitSurveyResponseHandler godoc
// @Summary Submit survey response (NPS Rating)
// @Description Submit a customer NPS rating and update calculation (promoters, passives, detractors, nps_score)
// @Tags marketing-surveys
// @Accept json
// @Produce json
// @Param id path int true "Survey ID"
// @Param request body SurveyResponsePayload true "Survey Response Rating"
// @Success 200 {object} utils.SuccessResponse{data=Survey} "Survey response recorded"
// @Failure 400 {object} utils.ErrorResponse "Invalid rating (must be 0-10) or ID"
// @Failure 500 {object} utils.ErrorResponse "Failed to record survey response"
// @Router /api/public/surveys/{id}/respond [post]
func SubmitSurveyResponseHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid survey ID", "")
	}

	var payload SurveyResponsePayload
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	if payload.Rating < 0 || payload.Rating > 10 {
		return utils.SendError(c, http.StatusBadRequest, "Rating NPS harus antara 0 dan 10", "")
	}

	updatedSurvey, err := RecordSurveyResponseService(uint(id), payload.Rating)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mencatat respon survei", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Respon survei berhasil dicatat", updatedSurvey)
}

// PublicSurveyInfo response model for public survey info
type PublicSurveyInfo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	GformURL    string `json:"gform_url"`
	State       string `json:"state"`
}

// GetPublicSurveyHandler godoc
// @Summary Get public survey details
// @Description Retrieve public info of a survey for respondents without authentication
// @Tags marketing-surveys
// @Produce json
// @Param id path int true "Survey ID"
// @Success 200 {object} utils.SuccessResponse{data=PublicSurveyInfo} "Survey details found"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 404 {object} utils.ErrorResponse "Survey not found"
// @Router /api/public/surveys/{id} [get]
func GetPublicSurveyHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid survey ID", "")
	}

	survey, err := GetSurveyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Survei tidak ditemukan", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Survei ditemukan", map[string]interface{}{
		"id":          survey.ID,
		"title":       survey.Title,
		"description": survey.Description,
		"gform_url":   survey.GformURL,
		"state":       survey.State,
	})
}



