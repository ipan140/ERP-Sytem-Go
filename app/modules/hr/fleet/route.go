package fleet

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/fleet", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateVehicleHandler)
	api.GET("", GetAllVehicleHandler)
	api.GET("/:id", GetVehicleByIDHandler)
	api.PUT("/:id", UpdateVehicleHandler)
	api.DELETE("/:id", DeleteVehicleHandler)

	api.POST("/vehiclelogcontract", CreateVehicleLogContractHandler)
	api.GET("/vehiclelogcontract", GetAllVehicleLogContractHandler)
	api.GET("/vehiclelogcontract/:id", GetVehicleLogContractByIDHandler)
	api.PUT("/vehiclelogcontract/:id", UpdateVehicleLogContractHandler)
	api.DELETE("/vehiclelogcontract/:id", DeleteVehicleLogContractHandler)

	api.POST("/vehiclelogfuel", CreateVehicleLogFuelHandler)
	api.GET("/vehiclelogfuel", GetAllVehicleLogFuelHandler)
	api.GET("/vehiclelogfuel/:id", GetVehicleLogFuelByIDHandler)
	api.PUT("/vehiclelogfuel/:id", UpdateVehicleLogFuelHandler)
	api.DELETE("/vehiclelogfuel/:id", DeleteVehicleLogFuelHandler)

	api.POST("/vehiclelogservices", CreateVehicleLogServicesHandler)
	api.GET("/vehiclelogservices", GetAllVehicleLogServicesHandler)
	api.GET("/vehiclelogservices/:id", GetVehicleLogServicesByIDHandler)
	api.PUT("/vehiclelogservices/:id", UpdateVehicleLogServicesHandler)
	api.DELETE("/vehiclelogservices/:id", DeleteVehicleLogServicesHandler)
}



