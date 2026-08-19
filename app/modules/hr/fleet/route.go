package fleet

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/fleet", middleware.Auth())
	api.POST("", CreateVehicleHandler)
	api.GET("", GetAllVehicleHandler)
	api.GET("/:id", GetVehicleByIDHandler)
	api.PUT("/:id", UpdateVehicleHandler)
	api.DELETE("/:id", DeleteVehicleHandler)
}
