package appraisals

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/appraisals", middleware.Auth(), middleware.RequireRoles(constants.RoleHRManager, constants.RoleFleetManager, constants.RoleEmployee))
	api.POST("", CreateAppraisalHandler)
	api.GET("", GetAllAppraisalHandler)
	api.GET("/:id", GetAppraisalByIDHandler)
	api.PUT("/:id", UpdateAppraisalHandler)
	api.DELETE("/:id", DeleteAppraisalHandler)
}
