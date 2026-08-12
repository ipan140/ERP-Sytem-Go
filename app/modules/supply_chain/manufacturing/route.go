package manufacturing

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/manufacturing", middleware.Auth())
	api.POST("", CreateManufacturingOrderHandler)
	api.GET("", GetAllManufacturingOrderHandler)
	api.GET("/:id", GetManufacturingOrderByIDHandler)
	api.PUT("/:id", UpdateManufacturingOrderHandler)
	api.DELETE("/:id", DeleteManufacturingOrderHandler)
}
