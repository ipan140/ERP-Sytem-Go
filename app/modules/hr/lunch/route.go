package lunch

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/lunch", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateLunchOrderHandler)
	api.GET("", GetAllLunchOrderHandler)
	api.GET("/:id", GetLunchOrderByIDHandler)
	api.PUT("/:id", UpdateLunchOrderHandler)
	api.DELETE("/:id", DeleteLunchOrderHandler)

	api.POST("/lunchorder", CreateLunchOrderHandler)
	api.GET("/lunchorder", GetAllLunchOrderHandler)
	api.GET("/lunchorder/:id", GetLunchOrderByIDHandler)
	api.PUT("/lunchorder/:id", UpdateLunchOrderHandler)
	api.DELETE("/lunchorder/:id", DeleteLunchOrderHandler)

	api.POST("/lunchcashmove", CreateLunchCashmoveHandler)
	api.GET("/lunchcashmove", GetAllLunchCashmoveHandler)
	api.GET("/lunchcashmove/:id", GetLunchCashmoveByIDHandler)
	api.PUT("/lunchcashmove/:id", UpdateLunchCashmoveHandler)
	api.DELETE("/lunchcashmove/:id", DeleteLunchCashmoveHandler)
}



