package approvals

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/approvals", middleware.Auth())
	api.POST("", CreateApprovalRequestHandler)
	api.GET("", GetAllApprovalRequestHandler)
	api.GET("/:id", GetApprovalRequestByIDHandler)
	api.PUT("/:id", UpdateApprovalRequestHandler)
	api.DELETE("/:id", DeleteApprovalRequestHandler)
}
