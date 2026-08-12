package mailer

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/mailer", middleware.Auth())
	api.POST("/test-send", TestSendMailHandler)
}
