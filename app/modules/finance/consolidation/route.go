package consolidation

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/consolidation", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("", GetAllConsolidationReportsHandler)
	api.POST("", CreateConsolidationReportHandler)
	api.GET("/:id", GetConsolidationReportByIDHandler)
	api.PUT("/:id", UpdateConsolidationReportHandler)
	api.DELETE("/:id", DeleteConsolidationReportHandler)
	api.POST("/generate", GenerateConsolidationHandler)
}



