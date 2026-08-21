package storage

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/core/storage", middleware.Auth(), middleware.RequireRoles(constants.RoleDirector, constants.RoleEmployee))
	api.POST("", CreateAttachmentHandler)
	api.GET("", GetAllAttachmentHandler)
	api.GET("/:id", GetAttachmentByIDHandler)
	api.PUT("/:id", UpdateAttachmentHandler)
	api.DELETE("/:id", DeleteAttachmentHandler)
}
