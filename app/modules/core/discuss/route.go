package discuss

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/discuss", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateChannelHandler)
	api.GET("", GetAllChannelHandler)
	api.GET("/:id", GetChannelByIDHandler)
	api.PUT("/:id", UpdateChannelHandler)
	api.DELETE("/:id", DeleteChannelHandler)
}



