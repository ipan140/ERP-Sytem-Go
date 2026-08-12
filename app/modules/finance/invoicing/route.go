package invoicing

import (
	"ERP-System/common/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/invoicing", middleware.Auth())
	api.POST("", CreateInvoiceHandler)
	api.GET("", GetAllInvoiceHandler)
	api.GET("/:id", GetInvoiceByIDHandler)
	api.PUT("/:id", UpdateInvoiceHandler)
	api.DELETE("/:id", DeleteInvoiceHandler)
	api.POST("/:id/post", PostInvoiceHandler)
	api.POST("/:id/refund", RefundInvoiceHandler)
	api.POST("/dunning", TriggerDunningHandler)
}
