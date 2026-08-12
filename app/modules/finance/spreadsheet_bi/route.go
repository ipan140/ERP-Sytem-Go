package spreadsheet_bi

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/spreadsheet_bi", middleware.Auth())
	api.POST("", CreateSpreadsheetHandler)
	api.GET("", GetAllSpreadsheetHandler)
	api.GET("/:id", GetSpreadsheetByIDHandler)
	api.PUT("/:id", UpdateSpreadsheetHandler)
	api.DELETE("/:id", DeleteSpreadsheetHandler)
}
