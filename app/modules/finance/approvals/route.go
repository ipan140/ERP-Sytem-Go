package approvals

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/approvals", middleware.Auth(), middleware.RequireRoles(constants.RoleFinanceManager, constants.RoleFinanceBilling))
	api.POST("", CreateApprovalRequestHandler)
	api.GET("", GetAllApprovalRequestHandler)
	api.GET("/:id", GetApprovalRequestByIDHandler)
	api.PUT("/:id", UpdateApprovalRequestHandler)
	api.DELETE("/:id", DeleteApprovalRequestHandler)
}
