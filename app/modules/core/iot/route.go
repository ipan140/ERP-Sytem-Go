package iot

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	for _, path := range []string{"/api/iot", "/api/core/iot"} {
		api := e.Group(path, middleware.Auth(), middleware.GlobalAutoRBAC())
		api.GET("/devices", GetAllIoTDeviceHandler)
		api.GET("/attendance-logs", GetAttendanceLogsHandler)
		api.POST("/ping/:id", PingDeviceHandler)
		api.POST("/sync-presensi", SyncPresensiHandler)
		api.POST("", CreateIoTDeviceHandler)
		api.GET("", GetAllIoTDeviceHandler)
		api.GET("/:id", GetIoTDeviceByIDHandler)
		api.PUT("/:id", UpdateIoTDeviceHandler)
		api.DELETE("/:id", DeleteIoTDeviceHandler)
	}
}



