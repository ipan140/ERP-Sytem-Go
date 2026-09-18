package whatsapp

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	registerWhatsappGroup(e.Group("/api/whatsapp", middleware.Auth(), middleware.GlobalAutoRBAC()))
	registerWhatsappGroup(e.Group("/api/core/whatsapp", middleware.Auth(), middleware.GlobalAutoRBAC()))
}

func registerWhatsappGroup(api *echo.Group) {
	// Config, Test & Logs
	api.GET("/config", GetWhatsappConfigHandler)
	api.POST("/config", SaveWhatsappConfigHandler)
	api.POST("/test", TestSendWhatsappHandler)
	api.GET("/logs", GetWaLogsHandler)

	// Templates
	api.POST("/templates", CreateWaTemplateHandler)
	api.GET("/templates", GetAllWaTemplateHandler)
	api.GET("/templates/:id", GetWaTemplateByIDHandler)
	api.PUT("/templates/:id", UpdateWaTemplateHandler)
	api.DELETE("/templates/:id", DeleteWaTemplateHandler)

	api.POST("", CreateWaTemplateHandler)
	api.GET("", GetAllWaTemplateHandler)
	api.GET("/:id", GetWaTemplateByIDHandler)
	api.PUT("/:id", UpdateWaTemplateHandler)
	api.DELETE("/:id", DeleteWaTemplateHandler)
}



