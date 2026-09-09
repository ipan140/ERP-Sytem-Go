package field_service

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/field_service", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateFieldServiceTaskHandler)
	api.GET("", GetAllFieldServiceTaskHandler)
	api.GET("/:id", GetFieldServiceTaskByIDHandler)
	api.PUT("/:id", UpdateFieldServiceTaskHandler)
	api.DELETE("/:id", DeleteFieldServiceTaskHandler)

	// Fase 2: e-BAST Validation Gate Route
	api.POST("/:id/validate-bast", ValidateBastHandler)

	// Fase 4: GPS Geotagging Routes
	api.POST("/:id/check-in", GPSCheckInHandler)
	api.POST("/:id/check-out", GPSCheckOutHandler)
}




