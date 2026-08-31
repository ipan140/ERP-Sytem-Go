package planning

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/planning", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateShiftHandler)
	api.GET("", GetAllShiftHandler)
	api.GET("/:id", GetShiftByIDHandler)
	api.PUT("/:id", UpdateShiftHandler)
	api.DELETE("/:id", DeleteShiftHandler)
}



