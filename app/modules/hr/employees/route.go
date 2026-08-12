package employees

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/employees", middleware.Auth())
	api.POST("", CreateEmployeeHandler)
	api.GET("", GetAllEmployeeHandler)
	api.GET("/:id", GetEmployeeByIDHandler)
	api.PUT("/:id", UpdateEmployeeHandler)
	api.DELETE("/:id", DeleteEmployeeHandler)
}
