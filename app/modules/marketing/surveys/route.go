package surveys

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/surveys", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateSurveyHandler)
	api.GET("", GetAllSurveyHandler)
	api.GET("/:id", GetSurveyByIDHandler)
	api.PUT("/:id", UpdateSurveyHandler)
	api.DELETE("/:id", DeleteSurveyHandler)
}



