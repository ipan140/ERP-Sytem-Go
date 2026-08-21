package forum

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/forum", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager, constants.RoleCustomer, constants.RoleGuest))
	api.POST("", CreateForumPostHandler)
	api.GET("", GetAllForumPostHandler)
	api.GET("/:id", GetForumPostByIDHandler)
	api.PUT("/:id", UpdateForumPostHandler)
	api.DELETE("/:id", DeleteForumPostHandler)
}
