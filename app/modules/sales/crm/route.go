package crm

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/crm", middleware.Auth())
	api.POST("", CreateLeadHandler)
	api.GET("", GetAllLeadHandler)
	api.GET("/:id", GetLeadByIDHandler)
	api.PUT("/:id", UpdateLeadHandler)
	api.DELETE("/:id", DeleteLeadHandler)
}
