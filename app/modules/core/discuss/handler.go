package discuss

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateChannel godoc
// @Summary Create a new Channel
// @Description Create a new Channel in the system
// @Tags discuss
// @Accept json
// @Produce json
// @Success 201 {object} Channel
// @Param request body Channel true "Payload"
// @Router /api/discuss [post]
// @Security BearerAuth
func CreateChannelHandler(c echo.Context) error {
	var data Channel
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateChannelService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllChannel godoc
// @Summary Get all Channel
// @Description Retrieve a list of all Channel
// @Tags discuss
// @Produce json
// @Success 200 {object} []Channel
// @Router /api/discuss [get]
// @Security BearerAuth
func GetAllChannelHandler(c echo.Context) error {
	data, err := GetAllChannelService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetChannelByID godoc
// @Summary Get a Channel by ID
// @Description Retrieve a specific Channel by its ID
// @Tags discuss
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} Channel
// @Router /api/discuss/{id} [get]
// @Security BearerAuth
func GetChannelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetChannelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateChannel godoc
// @Summary Update a Channel
// @Description Update an existing Channel
// @Tags discuss
// @Accept json
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/discuss/{id} [put]
// @Security BearerAuth
func UpdateChannelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetChannelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateChannelService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteChannel godoc
// @Summary Delete a Channel
// @Description Delete a Channel by ID
// @Tags discuss
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/discuss/{id} [delete]
// @Security BearerAuth
func DeleteChannelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteChannelService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


