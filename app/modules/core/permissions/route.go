package permissions

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Hanya SUPERADMIN yang boleh mengakses halaman "Settings / Pengatur Toggle" ini
	api := e.Group("/api/core/permissions", middleware.Auth(), middleware.GlobalAutoRBAC(), middleware.RequireRoles(constants.RoleSuperadmin))

	api.GET("", GetAllPermissionsHandler)
	api.POST("/toggle", TogglePermissionHandler)
	api.GET("/modules", GetAvailableModulesHandler)
}

