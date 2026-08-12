package appointments

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/appointments", middleware.Auth())
	api.POST("", CreateAppointmentHandler)
	api.GET("", GetAllAppointmentHandler)
	api.GET("/:id", GetAppointmentByIDHandler)
	api.PUT("/:id", UpdateAppointmentHandler)
	api.DELETE("/:id", DeleteAppointmentHandler)
}
