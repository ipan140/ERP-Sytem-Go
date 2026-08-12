package events

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateEvent godoc
// @Summary Create a new Event
// @Description Create a new Event in the system
// @Tags marketing-events
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/marketing/events [post]
// @Security BearerAuth
func CreateEventHandler(c echo.Context) error {
	var data Event
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateEventService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllEvent godoc
// @Summary Get all Event
// @Description Retrieve a list of all Event
// @Tags marketing-events
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events [get]
// @Security BearerAuth
func GetAllEventHandler(c echo.Context) error {
	data, err := GetAllEventService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetEventByID godoc
// @Summary Get a Event by ID
// @Description Retrieve a specific Event by its ID
// @Tags marketing-events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/{id} [get]
// @Security BearerAuth
func GetEventByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEventByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateEvent godoc
// @Summary Update a Event
// @Description Update an existing Event
// @Tags marketing-events
// @Accept json
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/{id} [put]
// @Security BearerAuth
func UpdateEventHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEventByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateEventService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteEvent godoc
// @Summary Delete a Event
// @Description Delete a Event by ID
// @Tags marketing-events
// @Produce json
// @Param id path int true "Event ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/{id} [delete]
// @Security BearerAuth
func DeleteEventHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEventService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
