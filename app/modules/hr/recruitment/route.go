package recruitment

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/recruitment", middleware.Auth())
	api.POST("", CreateApplicantHandler)
	api.GET("", GetAllApplicantHandler)
	api.GET("/:id", GetApplicantByIDHandler)
	api.PUT("/:id", UpdateApplicantHandler)
	api.DELETE("/:id", DeleteApplicantHandler)
}
