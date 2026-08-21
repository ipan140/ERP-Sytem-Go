package spreadsheet_bi

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/spreadsheet_bi", middleware.Auth(), middleware.RequireRoles(constants.RoleFinanceManager, constants.RoleFinanceBilling))
	api.POST("", CreateSpreadsheetHandler)
	api.GET("", GetAllSpreadsheetHandler)
	api.GET("/:id", GetSpreadsheetByIDHandler)
	api.PUT("/:id", UpdateSpreadsheetHandler)
	api.DELETE("/:id", DeleteSpreadsheetHandler)
}
