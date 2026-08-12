package time_off

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/time_off", middleware.Auth())
	api.POST("", CreateLeaveRequestHandler)
	api.GET("", GetAllLeaveRequestHandler)
	api.GET("/:id", GetLeaveRequestByIDHandler)
	api.PUT("/:id", UpdateLeaveRequestHandler)
	api.DELETE("/:id", DeleteLeaveRequestHandler)
}
