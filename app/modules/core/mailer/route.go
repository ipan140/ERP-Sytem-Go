package mailer

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/mailer", middleware.Auth())
	api.POST("", CreateEmailLogHandler)
	api.GET("", GetAllEmailLogHandler)
	api.GET("/:id", GetEmailLogByIDHandler)
	api.PUT("/:id", UpdateEmailLogHandler)
	api.DELETE("/:id", DeleteEmailLogHandler)
}
