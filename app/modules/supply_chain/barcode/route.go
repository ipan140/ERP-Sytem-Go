package barcode

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/barcode", middleware.Auth(), middleware.RequireRoles(constants.RoleWarehouseManager, constants.RoleWarehouseWorker, constants.RolePurchasing, constants.RoleManufacturingManager, constants.RoleQualityManager, constants.RoleMaintenanceManager))
	api.POST("", CreateBarcodeConfigHandler)
	api.GET("", GetAllBarcodeConfigHandler)
	api.GET("/:id", GetBarcodeConfigByIDHandler)
	api.PUT("/:id", UpdateBarcodeConfigHandler)
	api.DELETE("/:id", DeleteBarcodeConfigHandler)
}
