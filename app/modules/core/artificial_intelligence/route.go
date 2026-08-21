package artificial_intelligence

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/artificial_intelligence", middleware.Auth())
	api.POST("", CreateAIPromptHandler)
	api.GET("", GetAllAIPromptHandler)
	api.GET("/:id", GetAIPromptByIDHandler)
	api.PUT("/:id", UpdateAIPromptHandler)
	api.DELETE("/:id", DeleteAIPromptHandler)
}
