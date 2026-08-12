package storage

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/storage", middleware.Auth())
	api.POST("", CreateAttachmentHandler)
	api.GET("", GetAllAttachmentHandler)
	api.GET("/:id", GetAttachmentByIDHandler)
	api.PUT("/:id", UpdateAttachmentHandler)
	api.DELETE("/:id", DeleteAttachmentHandler)
}
