package budget

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/budget", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("", GetAllBudgetsHandler)
	api.POST("", CreateBudgetHandler)
	api.PUT("/:id", UpdateBudgetHandler)
	api.DELETE("/:id", DeleteBudgetHandler)
}
