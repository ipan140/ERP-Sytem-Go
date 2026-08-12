package recruitment

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/recruitment", middleware.Auth())
	api.POST("", CreateJobApplicantHandler)
	api.GET("", GetAllJobApplicantHandler)
	api.GET("/:id", GetJobApplicantByIDHandler)
	api.PUT("/:id", UpdateJobApplicantHandler)
	api.DELETE("/:id", DeleteJobApplicantHandler)
}
