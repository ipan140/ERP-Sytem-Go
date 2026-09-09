package repairs

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/repairs", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateRepairOrderHandler)
	api.GET("", GetAllRepairOrderHandler)
	api.GET("/:id", GetRepairOrderByIDHandler)
	api.PUT("/:id", UpdateRepairOrderHandler)
	api.DELETE("/:id", DeleteRepairOrderHandler)

	// Fase 2: QC Gate Route
	api.POST("/:id/qc-pass", PassQCHandler)

	// Fase 4: Public Tracking Portal (No Auth)
	public := e.Group("/api/public/services/tracking")
	public.GET("", PublicTrackingHandler)
	public.POST("/approve-estimate", PublicApproveEstimateHandler)
}




