package elearning

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/elearning", middleware.Auth())
	api.POST("", CreateCourseHandler)
	api.GET("", GetAllCourseHandler)
	api.GET("/:id", GetCourseByIDHandler)
	api.PUT("/:id", UpdateCourseHandler)
	api.DELETE("/:id", DeleteCourseHandler)
}
