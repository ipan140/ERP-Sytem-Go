package knowledge

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/knowledge", middleware.Auth())
	api.POST("", CreateArticleHandler)
	api.GET("", GetAllArticleHandler)
	api.GET("/:id", GetArticleByIDHandler)
	api.PUT("/:id", UpdateArticleHandler)
	api.DELETE("/:id", DeleteArticleHandler)
}
