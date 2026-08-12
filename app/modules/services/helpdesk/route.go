package helpdesk

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/services/helpdesk", middleware.Auth())
	api.POST("", CreateTicketHandler)
	api.GET("", GetAllTicketHandler)
	api.GET("/:id", GetTicketByIDHandler)
	api.PUT("/:id", UpdateTicketHandler)
	api.DELETE("/:id", DeleteTicketHandler)
}
