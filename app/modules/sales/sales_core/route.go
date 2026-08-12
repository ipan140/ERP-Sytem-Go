package sales_core

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/sales_core", middleware.Auth())
	api.POST("", CreateSaleOrderHandler)
	api.GET("", GetAllSaleOrderHandler)
	api.GET("/:id", GetSaleOrderByIDHandler)
	api.PUT("/:id", UpdateSaleOrderHandler)
	api.DELETE("/:id", DeleteSaleOrderHandler)
}
