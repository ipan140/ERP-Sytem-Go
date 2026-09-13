package dashboards

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/dashboards", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateDashboardHandler)
	api.GET("", GetAllDashboardHandler)
	api.GET("/system-health", GetSystemHealthHandler)
	api.GET("/audit-logs", GetAuditLogsHandler)
	api.GET("/:id", GetDashboardByIDHandler)
	api.PUT("/:id", UpdateDashboardHandler)
	api.DELETE("/:id", DeleteDashboardHandler)

	// Direct routes for enterprise core settings
	sys := e.Group("/api/sys")
	sys.GET("/health", GetSystemHealthHandler)
	sys.GET("/audit", GetAuditLogsHandler, middleware.Auth(), middleware.GlobalAutoRBAC())

	coreSys := e.Group("/api/core")
	coreSys.GET("/system-health", GetSystemHealthHandler)
	coreSys.GET("/audit", GetAuditLogsHandler, middleware.Auth(), middleware.GlobalAutoRBAC())
}



