package purchase

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/purchase", middleware.Auth())
	api.POST("", CreatePurchaseOrderHandler)
	api.GET("", GetAllPurchaseOrderHandler)
	api.GET("/:id", GetPurchaseOrderByIDHandler)
	api.PUT("/:id", UpdatePurchaseOrderHandler)
	api.DELETE("/:id", DeletePurchaseOrderHandler)
}
