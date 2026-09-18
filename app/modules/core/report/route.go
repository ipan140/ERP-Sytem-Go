package report

import (
	"ERP-System/common/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	for _, path := range []string{"/api/core/report", "/api/report"} {
		api := e.Group(path, middleware.Auth(), middleware.GlobalAutoRBAC())
		api.GET("/templates", GetPrintTemplatesHandler)
		api.POST("/templates", CreatePrintTemplateHandler)
		api.GET("/exports", GetExportReportsHandler)
		api.POST("/excel", GenerateDynamicExcelHandler)
		api.POST("/pdf", GenerateDynamicPDFHandler)
	}
}


