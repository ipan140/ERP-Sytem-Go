package attendances

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/attendances", middleware.Auth())
	api.POST("", CreateAttendanceHandler)
	api.GET("", GetAllAttendanceHandler)
	api.GET("/:id", GetAttendanceByIDHandler)
	api.PUT("/:id", UpdateAttendanceHandler)
	api.DELETE("/:id", DeleteAttendanceHandler)
}
