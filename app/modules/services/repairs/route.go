package repairs

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/repairs", middleware.Auth(), middleware.RequireRoles(constants.RoleProjectManager, constants.RoleSupportAgent))
	api.POST("", CreateRepairOrderHandler)
	api.GET("", GetAllRepairOrderHandler)
	api.GET("/:id", GetRepairOrderByIDHandler)
	api.PUT("/:id", UpdateRepairOrderHandler)
	api.DELETE("/:id", DeleteRepairOrderHandler)
}
