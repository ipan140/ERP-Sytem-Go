package maintenance

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/maintenance", middleware.Auth(), middleware.GlobalAutoRBAC())

	// Executive Summary & Equipment management (must precede /:id)
	api.GET("/summary", GetMaintenanceSummaryHandler)
	api.GET("/equipment", GetMaintenanceEquipmentsHandler)
	api.POST("/equipment", CreateMaintenanceEquipmentHandler)
	api.DELETE("/equipment/:id", DeleteMaintenanceEquipmentHandler)

	// Maintenance Requests CRUD & State Workflow
	api.POST("", CreateMaintenanceRequestHandler)
	api.GET("", GetAllMaintenanceRequestHandler)
	api.GET("/:id", GetMaintenanceRequestByIDHandler)
	api.PUT("/:id/state", UpdateMaintenanceStateHandler)
	api.PUT("/:id", UpdateMaintenanceRequestHandler)
	api.DELETE("/:id", DeleteMaintenanceRequestHandler)
}




