package user_roles

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// ?? SANGAT PENTING: Grup ini HANYA BISA DIAKSES oleh SUPERADMIN! ??
	// Bahkan HR Manager atau Director sekalipun tidak bisa mengubah hak akses.
	api := e.Group("/api/core/user_roles", middleware.Auth(), middleware.RequireRoles(constants.RoleSuperadmin))

	api.GET("", GetAllUsersRolesHandler)
	api.POST("/assign", AssignRoleHandler)

	// CRUD Dinamis untuk tabel Roles (Setting Roles)
	rolesApi := e.Group("/api/core/roles", middleware.Auth(), middleware.RequireRoles(constants.RoleSuperadmin))
	rolesApi.POST("", CreateRoleHandler)
	rolesApi.GET("", GetAllRolesHandler)
	rolesApi.PUT("/:id", UpdateRoleHandler)
	rolesApi.DELETE("/:id", DeleteRoleHandler)
}
