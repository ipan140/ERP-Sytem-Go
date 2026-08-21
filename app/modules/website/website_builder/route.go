package website_builder

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/website_builder", middleware.Auth())
	api.POST("", CreatePageHandler)
	api.GET("", GetAllPageHandler)
	api.GET("/:id", GetPageByIDHandler)
	api.PUT("/:id", UpdatePageHandler)
	api.DELETE("/:id", DeletePageHandler)
}
