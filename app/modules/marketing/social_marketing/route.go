package social_marketing

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/social_marketing", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager))
	api.POST("", CreateSocialPostHandler)
	api.GET("", GetAllSocialPostHandler)
	api.GET("/:id", GetSocialPostByIDHandler)
	api.PUT("/:id", UpdateSocialPostHandler)
	api.DELETE("/:id", DeleteSocialPostHandler)
}
