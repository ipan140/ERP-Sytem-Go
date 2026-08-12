package maintenance

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/maintenance", middleware.Auth())
	api.POST("", CreateMaintenanceRequestHandler)
	api.GET("", GetAllMaintenanceRequestHandler)
	api.GET("/:id", GetMaintenanceRequestByIDHandler)
	api.PUT("/:id", UpdateMaintenanceRequestHandler)
	api.DELETE("/:id", DeleteMaintenanceRequestHandler)
}
