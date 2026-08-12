package blog

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/blog", middleware.Auth())
	api.POST("", CreateBlogPostHandler)
	api.GET("", GetAllBlogPostHandler)
	api.GET("/:id", GetBlogPostByIDHandler)
	api.PUT("/:id", UpdateBlogPostHandler)
	api.DELETE("/:id", DeleteBlogPostHandler)
}
