package storage

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/storage", middleware.Auth(), middleware.GlobalAutoRBAC())
	// Config & Stats
	api.GET("/config", GetStorageConfigHandler)
	api.POST("/config", SaveStorageConfigHandler)
	api.GET("/stats", GetStorageStatsHandler)
	api.POST("/clean-temp", CleanTempStorageHandler)

	// Attachments
	api.POST("", CreateAttachmentHandler)
	api.GET("", GetAllAttachmentHandler)
	api.GET("/:id", GetAttachmentByIDHandler)
	api.PUT("/:id", UpdateAttachmentHandler)
	api.DELETE("/:id", DeleteAttachmentHandler)
}



