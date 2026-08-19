package events

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/events", middleware.Auth())
	api.POST("", CreateEventHandler)
	api.GET("", GetAllEventHandler)
	api.GET("/:id", GetEventByIDHandler)
	api.PUT("/:id", UpdateEventHandler)
	api.DELETE("/:id", DeleteEventHandler)

	api.POST("/eventticket", CreateEventTicketHandler)
	api.GET("/eventticket", GetAllEventTicketHandler)
	api.GET("/eventticket/:id", GetEventTicketByIDHandler)
	api.PUT("/eventticket/:id", UpdateEventTicketHandler)
	api.DELETE("/eventticket/:id", DeleteEventTicketHandler)

}
