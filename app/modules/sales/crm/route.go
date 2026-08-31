package crm

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/crm", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateLeadHandler)
	api.GET("", GetAllLeadHandler)
	api.GET("/:id", GetLeadByIDHandler)
	api.PUT("/:id", UpdateLeadHandler)
	api.DELETE("/:id", DeleteLeadHandler)

	api.POST("/salesteam", CreateSalesTeamHandler)
	api.GET("/salesteam", GetAllSalesTeamHandler)
	api.GET("/salesteam/:id", GetSalesTeamByIDHandler)
	api.PUT("/salesteam/:id", UpdateSalesTeamHandler)
	api.DELETE("/salesteam/:id", DeleteSalesTeamHandler)

	api.POST("/stage", CreateStageHandler)
	api.GET("/stage", GetAllStageHandler)
	api.GET("/stage/:id", GetStageByIDHandler)
	api.PUT("/stage/:id", UpdateStageHandler)
	api.DELETE("/stage/:id", DeleteStageHandler)

	api.POST("/activity", CreateActivityHandler)
	api.GET("/activity", GetAllActivityHandler)
	api.GET("/activity/:id", GetActivityByIDHandler)
	api.PUT("/activity/:id", UpdateActivityHandler)
	api.DELETE("/activity/:id", DeleteActivityHandler)

	api.POST("/salescommission", CreateSalesCommissionHandler)
	api.GET("/salescommission", GetAllSalesCommissionHandler)
	api.GET("/salescommission/:id", GetSalesCommissionByIDHandler)
	api.PUT("/salescommission/:id", UpdateSalesCommissionHandler)
	api.DELETE("/salescommission/:id", DeleteSalesCommissionHandler)
}



