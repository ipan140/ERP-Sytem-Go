package documents

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/documents", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateWorkspaceHandler)
	api.GET("", GetAllWorkspaceHandler)
	api.GET("/:id", GetWorkspaceByIDHandler)
	api.PUT("/:id", UpdateWorkspaceHandler)
	api.DELETE("/:id", DeleteWorkspaceHandler)
}



