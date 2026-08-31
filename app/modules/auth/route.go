package auth

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/auth")

	// Endpoint publik (tidak perlu middleware auth)
	api.POST("/login", LoginHandler)
	api.POST("/register", RegisterHandler)
}

