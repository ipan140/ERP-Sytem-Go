package lunch

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/lunch", middleware.Auth())
	api.POST("", CreateLunchOrderHandler)
	api.GET("", GetAllLunchOrderHandler)
	api.GET("/:id", GetLunchOrderByIDHandler)
	api.PUT("/:id", UpdateLunchOrderHandler)
	api.DELETE("/:id", DeleteLunchOrderHandler)
}
