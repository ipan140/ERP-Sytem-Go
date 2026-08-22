package artificial_intelligence

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAIPrompt godoc
// @Summary Create a new AIPrompt
// @Description Create a new AIPrompt in the system
// @Tags artificial_intelligence
// @Accept json
// @Produce json
// @Success 201 {object} AIPrompt
// @Param request body AIPrompt true "Payload"
// @Router /api/artificial_intelligence [post]
// @Security BearerAuth
func CreateAIPromptHandler(c echo.Context) error {
	var data AIPrompt
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAIPromptService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAIPrompt godoc
// @Summary Get all AIPrompt
// @Description Retrieve a list of all AIPrompt
// @Tags artificial_intelligence
// @Produce json
// @Success 200 {object} []AIPrompt
// @Router /api/artificial_intelligence [get]
// @Security BearerAuth
func GetAllAIPromptHandler(c echo.Context) error {
	data, err := GetAllAIPromptService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAIPromptByID godoc
// @Summary Get a AIPrompt by ID
// @Description Retrieve a specific AIPrompt by its ID
// @Tags artificial_intelligence
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} AIPrompt
// @Router /api/artificial_intelligence/{id} [get]
// @Security BearerAuth
func GetAIPromptByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAIPromptByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAIPrompt godoc
// @Summary Update a AIPrompt
// @Description Update an existing AIPrompt
// @Tags artificial_intelligence
// @Accept json
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/artificial_intelligence/{id} [put]
// @Security BearerAuth
func UpdateAIPromptHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAIPromptByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAIPromptService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAIPrompt godoc
// @Summary Delete a AIPrompt
// @Description Delete a AIPrompt by ID
// @Tags artificial_intelligence
// @Produce json
// @Param id path int true "AIPrompt ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/artificial_intelligence/{id} [delete]
// @Security BearerAuth
func DeleteAIPromptHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAIPromptService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


