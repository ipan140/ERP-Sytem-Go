package discuss

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	for _, path := range []string{"/api/discuss", "/api/core/discuss"} {
		api := e.Group(path, middleware.Auth(), middleware.GlobalAutoRBAC())
		api.GET("/channels", GetAllChannelHandler)
		api.GET("/messages", GetMessagesHandler)
		api.POST("/messages", SendMessageHandler)
		api.POST("", CreateChannelHandler)
		api.GET("", GetAllChannelHandler)
		api.GET("/:id", GetChannelByIDHandler)
		api.PUT("/:id", UpdateChannelHandler)
		api.DELETE("/:id", DeleteChannelHandler)
	}
}



