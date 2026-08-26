package calendar

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	g := e.Group("/api/calendar", middleware.Auth())
	g.GET("/events", GetEventsHandler)
	g.POST("/events", CreateEventHandler)
	g.PUT("/events/:id", UpdateEventHandler)
	g.DELETE("/events/:id", DeleteEventHandler)
}
