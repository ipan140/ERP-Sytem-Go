package activity_logs

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/activity-logs", middleware.Auth())
	api.GET("", GetActivityLogsHandler)
}
