package mailer

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/mailer", middleware.Auth(), middleware.GlobalAutoRBAC())
	// SMTP Outgoing Server Config
	api.GET("/config", GetSmtpConfigHandler)
	api.POST("/config", SaveSmtpConfigHandler)
	api.POST("/test", TestSendEmailHandler)

	// Email Logs Outbox
	api.POST("", CreateEmailLogHandler)
	api.GET("", GetAllEmailLogHandler)
	api.GET("/:id", GetEmailLogByIDHandler)
	api.PUT("/:id", UpdateEmailLogHandler)
	api.DELETE("/:id", DeleteEmailLogHandler)
}



