package dashboard

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/dashboard", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("/summary", GetScmDashboardSummaryHandler)
	api.GET("/calendar-events", GetScmCalendarEventsHandler)
}
