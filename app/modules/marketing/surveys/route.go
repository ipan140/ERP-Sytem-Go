package surveys

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/surveys", middleware.Auth())
	api.POST("", CreateSurveyHandler)
	api.GET("", GetAllSurveyHandler)
	api.GET("/:id", GetSurveyByIDHandler)
	api.PUT("/:id", UpdateSurveyHandler)
	api.DELETE("/:id", DeleteSurveyHandler)
}
