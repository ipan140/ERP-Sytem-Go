package helpdesk

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateTicket godoc
// @Summary Create a new Ticket
// @Description Create a new Ticket in the system
// @Tags services-helpdesk
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/helpdesk [post]
// @Security BearerAuth
func CreateTicketHandler(c echo.Context) error {
	var data Ticket
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateTicketService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllTicket godoc
// @Summary Get all Ticket
// @Description Retrieve a list of all Ticket
// @Tags services-helpdesk
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/helpdesk [get]
// @Security BearerAuth
func GetAllTicketHandler(c echo.Context) error {
	data, err := GetAllTicketService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetTicketByID godoc
// @Summary Get a Ticket by ID
// @Description Retrieve a specific Ticket by its ID
// @Tags services-helpdesk
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/helpdesk/{id} [get]
// @Security BearerAuth
func GetTicketByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTicketByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateTicket godoc
// @Summary Update a Ticket
// @Description Update an existing Ticket
// @Tags services-helpdesk
// @Accept json
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/helpdesk/{id} [put]
// @Security BearerAuth
func UpdateTicketHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetTicketByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateTicketService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteTicket godoc
// @Summary Delete a Ticket
// @Description Delete a Ticket by ID
// @Tags services-helpdesk
// @Produce json
// @Param id path int true "Ticket ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/helpdesk/{id} [delete]
// @Security BearerAuth
func DeleteTicketHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTicketService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

func CreateHelpdeskSLAHandler(c echo.Context) error { var data HelpdeskSLA; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateHelpdeskSLAService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllHelpdeskSLAHandler(c echo.Context) error { data, err := GetAllHelpdeskSLAService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetHelpdeskSLAByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetHelpdeskSLAByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateHelpdeskSLAHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetHelpdeskSLAByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateHelpdeskSLAService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteHelpdeskSLAHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteHelpdeskSLAService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

func CreateHelpdeskCannedResponseHandler(c echo.Context) error { var data HelpdeskCannedResponse; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateHelpdeskCannedResponseService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
func GetAllHelpdeskCannedResponseHandler(c echo.Context) error { data, err := GetAllHelpdeskCannedResponseService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetHelpdeskCannedResponseByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetHelpdeskCannedResponseByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func UpdateHelpdeskCannedResponseHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetHelpdeskCannedResponseByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateHelpdeskCannedResponseService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func DeleteHelpdeskCannedResponseHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteHelpdeskCannedResponseService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

