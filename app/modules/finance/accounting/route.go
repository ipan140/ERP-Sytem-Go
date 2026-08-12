package accounting

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/accounting", middleware.Auth())
	api.POST("", CreateJournalEntryHandler)
	api.GET("", GetAllJournalEntryHandler)
	api.GET("/:id", GetJournalEntryByIDHandler)
	api.PUT("/:id", UpdateJournalEntryHandler)
	api.DELETE("/:id", DeleteJournalEntryHandler)
}
