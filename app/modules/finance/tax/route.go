package tax

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Public export route
	e.GET("/api/finance/tax/export-ebupot", ExportEBupotCsvHandler)

	api := e.Group("/api/finance/tax", middleware.Auth(), middleware.GlobalAutoRBAC())
	// Transaction summary
	api.GET("", GetAllTaxReportsHandler)
	api.POST("", CreateTaxReportHandler)
	api.PUT("/:id", UpdateTaxReportHandler)
	api.DELETE("/:id", DeleteTaxReportHandler)

	// Master Tax Configs
	api.GET("/configs", GetAllTaxMasterConfigsHandler)
	api.POST("/configs", CreateTaxMasterConfigHandler)
	api.PUT("/configs/:id", UpdateTaxMasterConfigHandler)
	api.DELETE("/configs/:id", DeleteTaxMasterConfigHandler)
}
