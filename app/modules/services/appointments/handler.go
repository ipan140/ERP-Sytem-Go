package appointments

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAppointment godoc
// @Summary Create a new Appointment
// @Description Create a new Appointment in the system
// @Tags services-appointments
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/appointments [post]
// @Security BearerAuth
func CreateAppointmentHandler(c echo.Context) error {
	var data Appointment
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAppointmentService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAppointment godoc
// @Summary Get all Appointment
// @Description Retrieve a list of all Appointment
// @Tags services-appointments
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/appointments [get]
// @Security BearerAuth
func GetAllAppointmentHandler(c echo.Context) error {
	data, err := GetAllAppointmentService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAppointmentByID godoc
// @Summary Get a Appointment by ID
// @Description Retrieve a specific Appointment by its ID
// @Tags services-appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/appointments/{id} [get]
// @Security BearerAuth
func GetAppointmentByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAppointmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAppointment godoc
// @Summary Update a Appointment
// @Description Update an existing Appointment
// @Tags services-appointments
// @Accept json
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/appointments/{id} [put]
// @Security BearerAuth
func UpdateAppointmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAppointmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAppointmentService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAppointment godoc
// @Summary Delete a Appointment
// @Description Delete a Appointment by ID
// @Tags services-appointments
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/appointments/{id} [delete]
// @Security BearerAuth
func DeleteAppointmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAppointmentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
