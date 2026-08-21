package live_chat

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/live_chat", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager, constants.RoleCustomer, constants.RoleGuest))
	api.POST("", CreateChatSessionHandler)
	api.GET("", GetAllChatSessionHandler)
	api.GET("/:id", GetChatSessionByIDHandler)
	api.PUT("/:id", UpdateChatSessionHandler)
	api.DELETE("/:id", DeleteChatSessionHandler)
}
