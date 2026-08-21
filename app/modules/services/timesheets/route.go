package timesheets

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/timesheets", middleware.Auth())
	api.POST("", CreateTimesheetHandler)
	api.GET("", GetAllTimesheetHandler)
	api.GET("/:id", GetTimesheetByIDHandler)
	api.PUT("/:id", UpdateTimesheetHandler)
	api.DELETE("/:id", DeleteTimesheetHandler)
}
