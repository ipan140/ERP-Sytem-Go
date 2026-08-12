package point_of_sale

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/point_of_sale", middleware.Auth())
	api.POST("", CreatePosSessionHandler)
	api.GET("", GetAllPosSessionHandler)
	api.GET("/:id", GetPosSessionByIDHandler)
	api.PUT("/:id", UpdatePosSessionHandler)
	api.DELETE("/:id", DeletePosSessionHandler)
}
