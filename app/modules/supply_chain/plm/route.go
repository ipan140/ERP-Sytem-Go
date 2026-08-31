package plm

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/plm", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateBomHandler)
	api.GET("", GetAllBomHandler)
	api.GET("/:id", GetBomByIDHandler)
	api.PUT("/:id", UpdateBomHandler)
	api.DELETE("/:id", DeleteBomHandler)
}



