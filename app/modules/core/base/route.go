package base

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/base", middleware.Auth())
	api.POST("", CreateCurrencyHandler)
	api.GET("", GetAllCurrencyHandler)
	api.GET("/:id", GetCurrencyByIDHandler)
	api.PUT("/:id", UpdateCurrencyHandler)
	api.DELETE("/:id", DeleteCurrencyHandler)
}
