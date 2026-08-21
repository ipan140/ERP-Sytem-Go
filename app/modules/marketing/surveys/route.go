package surveys

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/surveys", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager))
	api.POST("", CreateSurveyHandler)
	api.GET("", GetAllSurveyHandler)
	api.GET("/:id", GetSurveyByIDHandler)
	api.PUT("/:id", UpdateSurveyHandler)
	api.DELETE("/:id", DeleteSurveyHandler)
}
