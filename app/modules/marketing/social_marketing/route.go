package social_marketing

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/social_marketing", middleware.Auth())
	api.POST("", CreateSocialPostHandler)
	api.GET("", GetAllSocialPostHandler)
	api.GET("/:id", GetSocialPostByIDHandler)
	api.PUT("/:id", UpdateSocialPostHandler)
	api.DELETE("/:id", DeleteSocialPostHandler)
}
