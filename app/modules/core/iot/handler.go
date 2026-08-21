package iot

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateIoTDevice godoc
// @Summary Create a new IoTDevice
// @Description Create a new IoTDevice in the system
// @Tags iot
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/iot [post]
// @Security BearerAuth
func CreateIoTDeviceHandler(c echo.Context) error {
	var data IoTDevice
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateIoTDeviceService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllIoTDevice godoc
// @Summary Get all IoTDevice
// @Description Retrieve a list of all IoTDevice
// @Tags iot
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/iot [get]
// @Security BearerAuth
func GetAllIoTDeviceHandler(c echo.Context) error {
	data, err := GetAllIoTDeviceService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetIoTDeviceByID godoc
// @Summary Get a IoTDevice by ID
// @Description Retrieve a specific IoTDevice by its ID
// @Tags iot
// @Produce json
// @Param id path int true "IoTDevice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/iot/{id} [get]
// @Security BearerAuth
func GetIoTDeviceByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetIoTDeviceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateIoTDevice godoc
// @Summary Update a IoTDevice
// @Description Update an existing IoTDevice
// @Tags iot
// @Accept json
// @Produce json
// @Param id path int true "IoTDevice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/iot/{id} [put]
// @Security BearerAuth
func UpdateIoTDeviceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetIoTDeviceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateIoTDeviceService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteIoTDevice godoc
// @Summary Delete a IoTDevice
// @Description Delete a IoTDevice by ID
// @Tags iot
// @Produce json
// @Param id path int true "IoTDevice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/iot/{id} [delete]
// @Security BearerAuth
func DeleteIoTDeviceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteIoTDeviceService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
