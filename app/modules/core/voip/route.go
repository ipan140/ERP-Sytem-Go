package voip

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/voip", middleware.Auth(), middleware.RequireRoles(constants.RoleDirector, constants.RoleEmployee))
	api.POST("", CreateCallRecordHandler)
	api.GET("", GetAllCallRecordHandler)
	api.GET("/:id", GetCallRecordByIDHandler)
	api.PUT("/:id", UpdateCallRecordHandler)
	api.DELETE("/:id", DeleteCallRecordHandler)
}
