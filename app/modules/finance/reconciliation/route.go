package reconciliation

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/reconciliation", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("", GetAllBankStatementsHandler)
	api.POST("", CreateBankStatementHandler)
	api.POST("/upload-csv", UploadBankStatementCsvHandler)
	api.POST("/auto", AutoReconcileHandler)
	api.PUT("/:id", UpdateBankStatementHandler)
	api.PUT("/:id/manual", ManualReconcileHandler)
	api.DELETE("/:id", DeleteBankStatementHandler)
}
