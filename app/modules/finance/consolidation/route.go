package consolidation

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/consolidation", middleware.Auth())
	api.POST("/generate", GenerateConsolidationHandler)
}
