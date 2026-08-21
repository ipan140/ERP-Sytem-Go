package live_chat

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateChatSession godoc
// @Summary Create a new ChatSession
// @Description Create a new ChatSession in the system
// @Tags website-live_chat
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/website/live_chat [post]
// @Security BearerAuth
func CreateChatSessionHandler(c echo.Context) error {
	var data ChatSession
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateChatSessionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllChatSession godoc
// @Summary Get all ChatSession
// @Description Retrieve a list of all ChatSession
// @Tags website-live_chat
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/website/live_chat [get]
// @Security BearerAuth
func GetAllChatSessionHandler(c echo.Context) error {
	data, err := GetAllChatSessionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetChatSessionByID godoc
// @Summary Get a ChatSession by ID
// @Description Retrieve a specific ChatSession by its ID
// @Tags website-live_chat
// @Produce json
// @Param id path int true "ChatSession ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/live_chat/{id} [get]
// @Security BearerAuth
func GetChatSessionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetChatSessionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateChatSession godoc
// @Summary Update a ChatSession
// @Description Update an existing ChatSession
// @Tags website-live_chat
// @Accept json
// @Produce json
// @Param id path int true "ChatSession ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/live_chat/{id} [put]
// @Security BearerAuth
func UpdateChatSessionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetChatSessionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateChatSessionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteChatSession godoc
// @Summary Delete a ChatSession
// @Description Delete a ChatSession by ID
// @Tags website-live_chat
// @Produce json
// @Param id path int true "ChatSession ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/website/live_chat/{id} [delete]
// @Security BearerAuth
func DeleteChatSessionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteChatSessionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
