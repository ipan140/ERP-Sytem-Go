package user_roles

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/user_roles", middleware.Auth(), middleware.GlobalAutoRBAC(), middleware.RequireRoles(constants.RoleSuperadmin))

	api.GET("", GetAllUsersRolesHandler)
	api.GET("/constants", GetRoleConstantsHandler)
	api.POST("/assign", AssignRoleHandler)

	rolesApi := e.Group("/api/core/roles", middleware.Auth(), middleware.GlobalAutoRBAC(), middleware.RequireRoles(constants.RoleSuperadmin))
	rolesApi.POST("", CreateRoleHandler)
	rolesApi.GET("", GetAllRolesHandler)
	rolesApi.PUT("/:id", UpdateRoleHandler)
	rolesApi.DELETE("/:id", DeleteRoleHandler)
}

