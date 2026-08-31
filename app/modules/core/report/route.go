package report

import (
	"ERP-System/common/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/report", middleware.Auth(), middleware.GlobalAutoRBAC())
	
	// Endpoint dinamis untuk export (Menerima data JSON dari Frontend)
	api.POST("/excel", GenerateDynamicExcelHandler)
	api.POST("/pdf", GenerateDynamicPDFHandler)
}


