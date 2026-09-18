package artificial_intelligence

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	for _, path := range []string{"/api/artificial_intelligence", "/api/core/ai"} {
		api := e.Group(path, middleware.Auth(), middleware.GlobalAutoRBAC())
		api.GET("/config", GetAIConfigHandler)
		api.POST("/config", SaveAIConfigHandler)
		api.POST("/generate", GenerateAISimulationHandler)
		api.POST("", CreateAIPromptHandler)
		api.GET("", GetAllAIPromptHandler)
		api.GET("/:id", GetAIPromptByIDHandler)
		api.PUT("/:id", UpdateAIPromptHandler)
		api.DELETE("/:id", DeleteAIPromptHandler)
	}
}



