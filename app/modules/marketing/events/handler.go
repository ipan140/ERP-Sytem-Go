package events

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

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

// @Summary Create EventTicket
// @Description Create a new EventTicket
// @Tags marketing-events
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/marketing/events/eventticket [post]
// @Security BearerAuth
func CreateEventTicketHandler(c echo.Context) error {
	var data EventTicket
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateEventTicketService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all EventTicket
// @Description Retrieve a list of all EventTicket
// @Tags marketing-events
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/eventticket [get]
// @Security BearerAuth
func GetAllEventTicketHandler(c echo.Context) error {
	data, err := GetAllEventTicketService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetEventTicketByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEventTicketByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update EventTicket
// @Description Update an existing EventTicket
// @Tags marketing-events
// @Accept json
// @Produce json
// @Param id path int true "EventTicket ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/eventticket/{id} [put]
// @Security BearerAuth
func UpdateEventTicketHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEventTicketByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateEventTicketService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete EventTicket
// @Description Delete EventTicket by ID
// @Tags marketing-events
// @Produce json
// @Param id path int true "EventTicket ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/events/eventticket/{id} [delete]
// @Security BearerAuth
func DeleteEventTicketHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEventTicketService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}
