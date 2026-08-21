package report

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/report", middleware.Auth(), middleware.RequireRoles(constants.RoleDirector, constants.RoleEmployee))
	api.POST("", CreateReportHandler)
	api.GET("", GetAllReportHandler)
	api.GET("/:id", GetReportByIDHandler)
	api.PUT("/:id", UpdateReportHandler)
	api.DELETE("/:id", DeleteReportHandler)
}
