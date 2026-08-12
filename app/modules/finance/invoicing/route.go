package invoicing

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/invoicing", middleware.Auth())
	api.POST("", CreateInvoiceHandler)
	api.GET("", GetAllInvoiceHandler)
	api.GET("/:id", GetInvoiceByIDHandler)
	api.PUT("/:id", UpdateInvoiceHandler)
	api.DELETE("/:id", DeleteInvoiceHandler)
}
