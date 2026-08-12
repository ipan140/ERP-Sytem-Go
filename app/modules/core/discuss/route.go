package discuss

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/discuss", middleware.Auth())
	api.POST("", CreateChannelHandler)
	api.GET("", GetAllChannelHandler)
	api.GET("/:id", GetChannelByIDHandler)
	api.PUT("/:id", UpdateChannelHandler)
	api.DELETE("/:id", DeleteChannelHandler)
}
