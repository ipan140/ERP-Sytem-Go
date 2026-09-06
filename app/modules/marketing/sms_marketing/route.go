package sms_marketing

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/sms_marketing", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateSmsCampaignHandler)
	api.GET("", GetAllSmsCampaignHandler)
	api.GET("/:id", GetSmsCampaignByIDHandler)
	api.PUT("/:id", UpdateSmsCampaignHandler)
	api.DELETE("/:id", DeleteSmsCampaignHandler)
	api.POST("/:id/broadcast", ExecuteBroadcastHandler)
	api.POST("/:id/send-test", SendTestMessageHandler)

	// WhatsApp Meta Cloud API / WABA Routes
	api.GET("/templates", GetAllWaTemplatesHandler)
	api.POST("/templates", CreateWaTemplateHandler)
	api.PUT("/templates/:id", UpdateWaTemplateHandler)
	api.DELETE("/templates/:id", DeleteWaTemplateHandler)

	api.GET("/config", GetWaConfigHandler)
	api.POST("/config", SaveWaConfigHandler)
}



