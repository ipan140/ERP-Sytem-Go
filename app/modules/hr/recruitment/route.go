package recruitment

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/recruitment", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateApplicantHandler)
	api.GET("", GetAllApplicantHandler)
	api.GET("/:id", GetApplicantByIDHandler)
	api.PUT("/:id", UpdateApplicantHandler)
	api.DELETE("/:id", DeleteApplicantHandler)
}



