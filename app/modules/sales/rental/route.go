package rental

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/rental", middleware.Auth())
	api.POST("", CreateRentalOrderHandler)
	api.GET("", GetAllRentalOrderHandler)
	api.GET("/:id", GetRentalOrderByIDHandler)
	api.PUT("/:id", UpdateRentalOrderHandler)
	api.DELETE("/:id", DeleteRentalOrderHandler)
}
