package appraisals

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/appraisals", middleware.Auth())
	api.POST("", CreateAppraisalHandler)
	api.GET("", GetAllAppraisalHandler)
	api.GET("/:id", GetAppraisalByIDHandler)
	api.PUT("/:id", UpdateAppraisalHandler)
	api.DELETE("/:id", DeleteAppraisalHandler)
}
