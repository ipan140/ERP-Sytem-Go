package iot

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/iot", middleware.Auth())
	api.POST("", CreateIoTDeviceHandler)
	api.GET("", GetAllIoTDeviceHandler)
	api.GET("/:id", GetIoTDeviceByIDHandler)
	api.PUT("/:id", UpdateIoTDeviceHandler)
	api.DELETE("/:id", DeleteIoTDeviceHandler)
}
