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

	api.POST("/course", CreateCourseHandler)
	api.GET("/course", GetAllCourseHandler)
	api.GET("/course/:id", GetCourseByIDHandler)
	api.PUT("/course/:id", UpdateCourseHandler)
	api.DELETE("/course/:id", DeleteCourseHandler)

	api.POST("/slide", CreateSlideHandler)
	api.GET("/slide", GetAllSlideHandler)
	api.GET("/slide/:id", GetSlideByIDHandler)
	api.PUT("/slide/:id", UpdateSlideHandler)
	api.DELETE("/slide/:id", DeleteSlideHandler)

	api.POST("/certification", CreateCertificationHandler)
	api.GET("/certification", GetAllCertificationHandler)
	api.GET("/certification/:id", GetCertificationByIDHandler)
	api.PUT("/certification/:id", UpdateCertificationHandler)
	api.DELETE("/certification/:id", DeleteCertificationHandler)

}
