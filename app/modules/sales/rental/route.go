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

	api.POST("/rentalorderline", CreateRentalOrderLineHandler)
	api.GET("/rentalorderline", GetAllRentalOrderLineHandler)
	api.GET("/rentalorderline/:id", GetRentalOrderLineByIDHandler)
	api.PUT("/rentalorderline/:id", UpdateRentalOrderLineHandler)
	api.DELETE("/rentalorderline/:id", DeleteRentalOrderLineHandler)
}
