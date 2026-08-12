package expenses

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/expenses", middleware.Auth())
	api.POST("", CreateExpenseHandler)
	api.GET("", GetAllExpenseHandler)
	api.GET("/:id", GetExpenseByIDHandler)
	api.PUT("/:id", UpdateExpenseHandler)
	api.DELETE("/:id", DeleteExpenseHandler)
}
