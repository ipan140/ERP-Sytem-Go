package project

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/project", middleware.Auth())
	api.POST("", CreateProjectHandler)
	api.GET("", GetAllProjectHandler)
	api.GET("/:id", GetProjectByIDHandler)
	api.PUT("/:id", UpdateProjectHandler)
	api.DELETE("/:id", DeleteProjectHandler)
}
