package maintenance

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/maintenance", middleware.Auth(), middleware.RequireRoles(constants.RoleWarehouseManager, constants.RoleWarehouseWorker, constants.RolePurchasing, constants.RoleManufacturingManager, constants.RoleQualityManager, constants.RoleMaintenanceManager))
	api.POST("", CreateMaintenanceRequestHandler)
	api.GET("", GetAllMaintenanceRequestHandler)
	api.GET("/:id", GetMaintenanceRequestByIDHandler)
	api.PUT("/:id", UpdateMaintenanceRequestHandler)
	api.DELETE("/:id", DeleteMaintenanceRequestHandler)
}
