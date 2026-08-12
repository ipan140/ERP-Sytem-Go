package accounting

import (
	"ERP-System/common/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/accounting", middleware.Auth())
	api.POST("", CreateJournalEntryHandler)
	api.GET("", GetAllJournalEntryHandler)
	api.GET("/:id", GetJournalEntryByIDHandler)
	api.PUT("/:id", UpdateJournalEntryHandler)
	api.DELETE("/:id", DeleteJournalEntryHandler)
}
