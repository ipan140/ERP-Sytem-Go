package expenses

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/expenses", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateExpenseHandler)
	api.GET("", GetAllExpenseHandler)
	api.GET("/:id", GetExpenseByIDHandler)
	api.PUT("/:id", UpdateExpenseHandler)
	api.DELETE("/:id", DeleteExpenseHandler)

	api.POST("/expensesheet", CreateExpenseSheetHandler)
	api.GET("/expensesheet", GetAllExpenseSheetHandler)
	api.GET("/expensesheet/:id", GetExpenseSheetByIDHandler)
	api.PUT("/expensesheet/:id", UpdateExpenseSheetHandler)
	api.DELETE("/expensesheet/:id", DeleteExpenseSheetHandler)

	// Kas Kecil (Petty Cash - Imprest Fund)
	api.GET("/petty-cash", GetPettyCashHandler)
	api.POST("/petty-cash/expense", RecordPettyCashExpenseHandler)
	api.POST("/petty-cash/replenish", ReplenishPettyCashHandler)
}



