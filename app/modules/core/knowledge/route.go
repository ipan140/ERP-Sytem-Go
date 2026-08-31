package knowledge

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/knowledge", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateArticleHandler)
	api.GET("", GetAllArticleHandler)
	api.GET("/:id", GetArticleByIDHandler)
	api.PUT("/:id", UpdateArticleHandler)
	api.DELETE("/:id", DeleteArticleHandler)
}



