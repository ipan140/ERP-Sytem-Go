package field_service

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/field_service", middleware.Auth())
	api.POST("", CreateFieldServiceTaskHandler)
	api.GET("", GetAllFieldServiceTaskHandler)
	api.GET("/:id", GetFieldServiceTaskByIDHandler)
	api.PUT("/:id", UpdateFieldServiceTaskHandler)
	api.DELETE("/:id", DeleteFieldServiceTaskHandler)
}
