package consolidation

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/consolidation", middleware.Auth())
	api.POST("", CreateConsolidationEntryHandler)
	api.GET("", GetAllConsolidationEntryHandler)
	api.GET("/:id", GetConsolidationEntryByIDHandler)
	api.PUT("/:id", UpdateConsolidationEntryHandler)
	api.DELETE("/:id", DeleteConsolidationEntryHandler)
}
