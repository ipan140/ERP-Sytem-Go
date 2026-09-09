package notifications

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/notifications", middleware.Auth())
	api.GET("/logs", GetNotificationLogsHandler)
	api.POST("/send", SendNotificationHandler)
}
