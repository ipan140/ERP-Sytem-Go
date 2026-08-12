package inventory

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/inventory", middleware.Auth())
	api.POST("", CreateProductHandler)
	api.GET("", GetAllProductHandler)
	api.GET("/:id", GetProductByIDHandler)
	api.PUT("/:id", UpdateProductHandler)
	api.DELETE("/:id", DeleteProductHandler)
}
