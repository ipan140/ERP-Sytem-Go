package planning

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/planning", middleware.Auth())
	api.POST("", CreateShiftHandler)
	api.GET("", GetAllShiftHandler)
	api.GET("/:id", GetShiftByIDHandler)
	api.PUT("/:id", UpdateShiftHandler)
	api.DELETE("/:id", DeleteShiftHandler)
}
