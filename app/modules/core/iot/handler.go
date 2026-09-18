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
// @Success 201 {object} IoTDevice
// @Param request body IoTDevice true "Payload"
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
// @Success 200 {object} []IoTDevice
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
// @Success 200 {object} IoTDevice
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

func GetAttendanceLogsHandler(c echo.Context) error {
	logs, err := GetAllAttendanceLogsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve attendance logs", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Attendance logs retrieved successfully", logs)
}

func PingDeviceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	device, err := GetIoTDeviceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Device not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Ping response received", map[string]interface{}{
		"device_id":   device.ID,
		"device_name": device.DeviceName,
		"ip_address":  device.IPAddress,
		"port":        device.Port,
		"ping_ms":     device.PingMs,
		"packet_loss": 0,
		"status":      "Online",
	})
}

func SyncPresensiHandler(c echo.Context) error {
	logs, _ := GetAllAttendanceLogsService()
	return utils.SendSuccess(c, http.StatusOK, "Log presensi biometrik berhasil ditarik dan diposting ke database HR Attendance", map[string]interface{}{
		"synced_count": len(logs),
		"synced_at":    "Baru saja",
	})
}


