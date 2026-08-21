package payroll

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/payroll", middleware.Auth())
	api.POST("", CreatePayslipHandler)
	api.GET("", GetAllPayslipHandler)
	api.GET("/:id", GetPayslipByIDHandler)
	api.PUT("/:id", UpdatePayslipHandler)
	api.DELETE("/:id", DeletePayslipHandler)

	api.POST("/payslipline", CreatePayslipLineHandler)
	api.GET("/payslipline", GetAllPayslipLineHandler)
	api.GET("/payslipline/:id", GetPayslipLineByIDHandler)
	api.PUT("/payslipline/:id", UpdatePayslipLineHandler)
	api.DELETE("/payslipline/:id", DeletePayslipLineHandler)

	api.POST("/salaryrule", CreateSalaryRuleHandler)
	api.GET("/salaryrule", GetAllSalaryRuleHandler)
	api.GET("/salaryrule/:id", GetSalaryRuleByIDHandler)
	api.PUT("/salaryrule/:id", UpdateSalaryRuleHandler)
	api.DELETE("/salaryrule/:id", DeleteSalaryRuleHandler)
}
