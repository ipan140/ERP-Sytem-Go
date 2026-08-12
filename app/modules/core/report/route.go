package report

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/report", middleware.Auth())
	api.POST("", CreateReportHandler)
	api.GET("", GetAllReportHandler)
	api.GET("/:id", GetReportByIDHandler)
	api.PUT("/:id", UpdateReportHandler)
	api.DELETE("/:id", DeleteReportHandler)
}
