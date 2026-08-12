package live_chat

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/live_chat", middleware.Auth())
	api.POST("", CreateChatSessionHandler)
	api.GET("", GetAllChatSessionHandler)
	api.GET("/:id", GetChatSessionByIDHandler)
	api.PUT("/:id", UpdateChatSessionHandler)
	api.DELETE("/:id", DeleteChatSessionHandler)
}
