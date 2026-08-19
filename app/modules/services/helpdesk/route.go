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

	api.POST("/helpdesksla", CreateHelpdeskSLAHandler)
	api.GET("/helpdesksla", GetAllHelpdeskSLAHandler)
	api.GET("/helpdesksla/:id", GetHelpdeskSLAByIDHandler)
	api.PUT("/helpdesksla/:id", UpdateHelpdeskSLAHandler)
	api.DELETE("/helpdesksla/:id", DeleteHelpdeskSLAHandler)
	api.POST("/helpdeskcannedresponse", CreateHelpdeskCannedResponseHandler)
	api.GET("/helpdeskcannedresponse", GetAllHelpdeskCannedResponseHandler)
	api.GET("/helpdeskcannedresponse/:id", GetHelpdeskCannedResponseByIDHandler)
	api.PUT("/helpdeskcannedresponse/:id", UpdateHelpdeskCannedResponseHandler)
	api.DELETE("/helpdeskcannedresponse/:id", DeleteHelpdeskCannedResponseHandler)

}
