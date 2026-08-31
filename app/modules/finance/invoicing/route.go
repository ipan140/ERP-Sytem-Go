package invoicing

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/invoicing", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateInvoiceHandler)
	api.GET("", GetAllInvoiceHandler)
	api.GET("/:id", GetInvoiceByIDHandler)
	api.PUT("/:id", UpdateInvoiceHandler)
	api.DELETE("/:id", DeleteInvoiceHandler)
	api.GET("/:id/export", ExportInvoicePDFHandler)
	api.GET("/:id/export/excel", ExportInvoiceExcelHTMLHandler)
	api.POST("/:id/post", PostInvoiceHandler)
	api.POST("/:id/refund", RefundInvoiceHandler)
	api.POST("/dunning", TriggerDunningHandler)
	api.POST("/paymenttermline", CreatePaymentTermLineHandler)
	api.GET("/paymenttermline", GetAllPaymentTermLineHandler)
	api.GET("/paymenttermline/:id", GetPaymentTermLineByIDHandler)
	api.PUT("/paymenttermline/:id", UpdatePaymentTermLineHandler)
	api.DELETE("/paymenttermline/:id", DeletePaymentTermLineHandler)

	api.POST("/taxrepartitionline", CreateTaxRepartitionLineHandler)
	api.GET("/taxrepartitionline", GetAllTaxRepartitionLineHandler)
	api.GET("/taxrepartitionline/:id", GetTaxRepartitionLineByIDHandler)
	api.PUT("/taxrepartitionline/:id", UpdateTaxRepartitionLineHandler)
	api.DELETE("/taxrepartitionline/:id", DeleteTaxRepartitionLineHandler)
}




