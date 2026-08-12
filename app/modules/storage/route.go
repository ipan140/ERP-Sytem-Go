package storage

import (
	"ERP-System/common/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Endpoint storage membutuhkan autentikasi
	api := e.Group("/api/storage", middleware.Auth())

	api.POST("/upload", UploadFileHandler)
}
