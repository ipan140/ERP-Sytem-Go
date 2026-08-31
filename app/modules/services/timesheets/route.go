package timesheets

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/timesheets", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateTimesheetHandler)
	api.GET("", GetAllTimesheetHandler)
	api.GET("/:id", GetTimesheetByIDHandler)
	api.PUT("/:id", UpdateTimesheetHandler)
	api.DELETE("/:id", DeleteTimesheetHandler)
}



