package artificial_intelligence

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/artificial_intelligence", middleware.Auth())
	api.POST("", CreateAIPromptHandler)
	api.GET("", GetAllAIPromptHandler)
	api.GET("/:id", GetAIPromptByIDHandler)
	api.PUT("/:id", UpdateAIPromptHandler)
	api.DELETE("/:id", DeleteAIPromptHandler)
}
