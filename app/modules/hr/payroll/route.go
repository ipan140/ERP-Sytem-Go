package payroll

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/payroll", middleware.Auth())
	api.POST("", CreatePayslipHandler)
	api.GET("", GetAllPayslipHandler)
	api.GET("/:id", GetPayslipByIDHandler)
	api.PUT("/:id", UpdatePayslipHandler)
	api.DELETE("/:id", DeletePayslipHandler)
}
