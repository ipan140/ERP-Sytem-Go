package surveys

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Protected Admin Routes
	api := e.Group("/api/marketing/surveys", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateSurveyHandler)
	api.GET("", GetAllSurveyHandler)
	api.GET("/:id", GetSurveyByIDHandler)
	api.PUT("/:id", UpdateSurveyHandler)
	api.DELETE("/:id", DeleteSurveyHandler)

	// Public Routes (Bisa diakses publik atau webhook Google Form/Sheets tanpa token)
	publicApi := e.Group("/api/public/surveys")
	publicApi.GET("/:id", GetPublicSurveyHandler)
	publicApi.POST("/:id/respond", SubmitSurveyResponseHandler)
	publicApi.POST("/:id/webhook-gform", SubmitSurveyResponseHandler)
}



