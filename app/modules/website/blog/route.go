package blog

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/blog", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager, constants.RoleCustomer, constants.RoleGuest))
	api.POST("", CreateBlogPostHandler)
	api.GET("", GetAllBlogPostHandler)
	api.GET("/:id", GetBlogPostByIDHandler)
	api.PUT("/:id", UpdateBlogPostHandler)
	api.DELETE("/:id", DeleteBlogPostHandler)
}
