package field_service

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateFieldServiceTask godoc
// @Summary Create a new FieldServiceTask
// @Description Create a new FieldServiceTask in the system
// @Tags services-field_service
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/services/field_service [post]
// @Security BearerAuth
func CreateFieldServiceTaskHandler(c echo.Context) error {
	var data FieldServiceTask
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateFieldServiceTaskService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllFieldServiceTask godoc
// @Summary Get all FieldServiceTask
// @Description Retrieve a list of all FieldServiceTask
// @Tags services-field_service
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service [get]
// @Security BearerAuth
func GetAllFieldServiceTaskHandler(c echo.Context) error {
	data, err := GetAllFieldServiceTaskService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetFieldServiceTaskByID godoc
// @Summary Get a FieldServiceTask by ID
// @Description Retrieve a specific FieldServiceTask by its ID
// @Tags services-field_service
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service/{id} [get]
// @Security BearerAuth
func GetFieldServiceTaskByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFieldServiceTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateFieldServiceTask godoc
// @Summary Update a FieldServiceTask
// @Description Update an existing FieldServiceTask
// @Tags services-field_service
// @Accept json
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service/{id} [put]
// @Security BearerAuth
func UpdateFieldServiceTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFieldServiceTaskByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateFieldServiceTaskService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteFieldServiceTask godoc
// @Summary Delete a FieldServiceTask
// @Description Delete a FieldServiceTask by ID
// @Tags services-field_service
// @Produce json
// @Param id path int true "FieldServiceTask ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/services/field_service/{id} [delete]
// @Security BearerAuth
func DeleteFieldServiceTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFieldServiceTaskService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
