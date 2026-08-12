package website_builder

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/website_builder", middleware.Auth())
	api.POST("", CreatePageHandler)
	api.GET("", GetAllPageHandler)
	api.GET("/:id", GetPageByIDHandler)
	api.PUT("/:id", UpdatePageHandler)
	api.DELETE("/:id", DeletePageHandler)
}
