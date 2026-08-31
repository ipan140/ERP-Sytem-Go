package quality

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/quality", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateQualityCheckHandler)
	api.GET("", GetAllQualityCheckHandler)
	api.GET("/:id", GetQualityCheckByIDHandler)
	api.PUT("/:id", UpdateQualityCheckHandler)
	api.DELETE("/:id", DeleteQualityCheckHandler)
}



