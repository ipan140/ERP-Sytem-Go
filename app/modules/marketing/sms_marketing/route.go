package sms_marketing

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/sms_marketing", middleware.Auth())
	api.POST("", CreateSmsCampaignHandler)
	api.GET("", GetAllSmsCampaignHandler)
	api.GET("/:id", GetSmsCampaignByIDHandler)
	api.PUT("/:id", UpdateSmsCampaignHandler)
	api.DELETE("/:id", DeleteSmsCampaignHandler)
}
