package dashboards

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/dashboards", middleware.Auth(), middleware.RequireRoles(constants.RoleDirector, constants.RoleEmployee))
	api.POST("", CreateDashboardHandler)
	api.GET("", GetAllDashboardHandler)
	api.GET("/:id", GetDashboardByIDHandler)
	api.PUT("/:id", UpdateDashboardHandler)
	api.DELETE("/:id", DeleteDashboardHandler)
}
