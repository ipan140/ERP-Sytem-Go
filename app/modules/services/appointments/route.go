package appointments

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/appointments", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateAppointmentHandler)
	api.GET("", GetAllAppointmentHandler)
	api.GET("/:id", GetAppointmentByIDHandler)
	api.PUT("/:id", UpdateAppointmentHandler)
	api.DELETE("/:id", DeleteAppointmentHandler)
}



