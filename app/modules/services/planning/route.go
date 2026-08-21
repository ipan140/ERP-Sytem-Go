package planning

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/planning", middleware.Auth(), middleware.RequireRoles(constants.RoleProjectManager, constants.RoleSupportAgent))
	api.POST("", CreateShiftHandler)
	api.GET("", GetAllShiftHandler)
	api.GET("/:id", GetShiftByIDHandler)
	api.PUT("/:id", UpdateShiftHandler)
	api.DELETE("/:id", DeleteShiftHandler)
}
