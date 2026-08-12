package dashboards

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/dashboards", middleware.Auth())
	api.POST("", CreateDashboardHandler)
	api.GET("", GetAllDashboardHandler)
	api.GET("/:id", GetDashboardByIDHandler)
	api.PUT("/:id", UpdateDashboardHandler)
	api.DELETE("/:id", DeleteDashboardHandler)
}
