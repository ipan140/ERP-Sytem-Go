package quality

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/quality", middleware.Auth())
	api.POST("", CreateQualityCheckHandler)
	api.GET("", GetAllQualityCheckHandler)
	api.GET("/:id", GetQualityCheckByIDHandler)
	api.PUT("/:id", UpdateQualityCheckHandler)
	api.DELETE("/:id", DeleteQualityCheckHandler)
}
