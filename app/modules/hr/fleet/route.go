package fleet

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/fleet", middleware.Auth())
	api.POST("", CreateFleetVehicleHandler)
	api.GET("", GetAllFleetVehicleHandler)
	api.GET("/:id", GetFleetVehicleByIDHandler)
	api.PUT("/:id", UpdateFleetVehicleHandler)
	api.DELETE("/:id", DeleteFleetVehicleHandler)
}
