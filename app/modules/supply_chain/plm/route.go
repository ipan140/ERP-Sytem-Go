package plm

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/plm", middleware.Auth())
	api.POST("", CreateBomHandler)
	api.GET("", GetAllBomHandler)
	api.GET("/:id", GetBomByIDHandler)
	api.PUT("/:id", UpdateBomHandler)
	api.DELETE("/:id", DeleteBomHandler)
}
