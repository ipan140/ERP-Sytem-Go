package appraisals

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/appraisals", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateAppraisalHandler)
	api.GET("", GetAllAppraisalHandler)
	api.GET("/:id", GetAppraisalByIDHandler)
	api.PUT("/:id", UpdateAppraisalHandler)
	api.DELETE("/:id", DeleteAppraisalHandler)
}



