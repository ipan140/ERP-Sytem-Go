package ecommerce

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/ecommerce", middleware.Auth())
	api.POST("", CreateCartHandler)
	api.GET("", GetAllCartHandler)
	api.GET("/:id", GetCartByIDHandler)
	api.PUT("/:id", UpdateCartHandler)
	api.DELETE("/:id", DeleteCartHandler)
}
