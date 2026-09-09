package quality

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/quality", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.GET("/summary", GetQualitySummaryHandler)
	api.GET("/point", GetQualityPointsHandler)
	api.POST("/point", CreateQualityPointHandler)
	api.DELETE("/point/:id", DeleteQualityPointHandler)

	api.POST("", CreateQualityCheckHandler)
	api.GET("", GetAllQualityCheckHandler)
	api.GET("/:id", GetQualityCheckByIDHandler)
	api.PUT("/:id/process", ProcessQualityCheckHandler)
	api.POST("/:id/process", ProcessQualityCheckHandler)
	api.PUT("/:id", UpdateQualityCheckHandler)
	api.DELETE("/:id", DeleteQualityCheckHandler)
}



