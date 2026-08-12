package barcode

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/barcode", middleware.Auth())
	api.POST("", CreateBarcodeConfigHandler)
	api.GET("", GetAllBarcodeConfigHandler)
	api.GET("/:id", GetBarcodeConfigByIDHandler)
	api.PUT("/:id", UpdateBarcodeConfigHandler)
	api.DELETE("/:id", DeleteBarcodeConfigHandler)
}
