package events

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/events", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateEventHandler)
	api.GET("", GetAllEventHandler)
	api.GET("/:id", GetEventByIDHandler)
	api.PUT("/:id", UpdateEventHandler)
	api.DELETE("/:id", DeleteEventHandler)

	api.POST("/eventticket", CreateEventTicketHandler)
	api.POST("/eventticket/scan", ScanTicketHandler)
	api.GET("/eventticket", GetAllEventTicketHandler)
	api.GET("/eventticket/:id", GetEventTicketByIDHandler)
	api.PUT("/eventticket/:id", UpdateEventTicketHandler)
	api.DELETE("/eventticket/:id", DeleteEventTicketHandler)

}
