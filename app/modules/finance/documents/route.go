package documents

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/documents", middleware.Auth())
	api.POST("", CreateFinanceDocumentHandler)
	api.GET("", GetAllFinanceDocumentHandler)
	api.GET("/:id", GetFinanceDocumentByIDHandler)
	api.PUT("/:id", UpdateFinanceDocumentHandler)
	api.DELETE("/:id", DeleteFinanceDocumentHandler)
}
