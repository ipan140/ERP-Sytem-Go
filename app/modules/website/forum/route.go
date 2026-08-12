package forum

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/forum", middleware.Auth())
	api.POST("", CreateForumPostHandler)
	api.GET("", GetAllForumPostHandler)
	api.GET("/:id", GetForumPostByIDHandler)
	api.PUT("/:id", UpdateForumPostHandler)
	api.DELETE("/:id", DeleteForumPostHandler)
}
