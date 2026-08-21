package barcode

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/barcode", middleware.Auth())
	api.POST("", CreateBarcodeConfigHandler)
	api.GET("", GetAllBarcodeConfigHandler)
	api.GET("/:id", GetBarcodeConfigByIDHandler)
	api.PUT("/:id", UpdateBarcodeConfigHandler)
	api.DELETE("/:id", DeleteBarcodeConfigHandler)
}
