package time_off

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/time_off", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateLeaveRequestHandler)
	api.GET("", GetAllLeaveRequestHandler)
	api.GET("/:id", GetLeaveRequestByIDHandler)
	api.PUT("/:id", UpdateLeaveRequestHandler)
	api.DELETE("/:id", DeleteLeaveRequestHandler)

	api.POST("/leavetype", CreateLeaveTypeHandler)
	api.GET("/leavetype", GetAllLeaveTypeHandler)
	api.GET("/leavetype/:id", GetLeaveTypeByIDHandler)
	api.PUT("/leavetype/:id", UpdateLeaveTypeHandler)
	api.DELETE("/leavetype/:id", DeleteLeaveTypeHandler)

	api.POST("/leaveallocation", CreateLeaveAllocationHandler)
	api.GET("/leaveallocation", GetAllLeaveAllocationHandler)
	api.GET("/leaveallocation/:id", GetLeaveAllocationByIDHandler)
	api.PUT("/leaveallocation/:id", UpdateLeaveAllocationHandler)
	api.DELETE("/leaveallocation/:id", DeleteLeaveAllocationHandler)
}



