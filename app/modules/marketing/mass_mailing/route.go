package mass_mailing

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/mass_mailing", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateMailingCampaignHandler)
	api.GET("", GetAllMailingCampaignHandler)
	api.GET("/:id", GetMailingCampaignByIDHandler)
	api.PUT("/:id", UpdateMailingCampaignHandler)
	api.DELETE("/:id", DeleteMailingCampaignHandler)

	api.POST("/utmtracker", CreateUtmTrackerHandler)
	api.GET("/utmtracker", GetAllUtmTrackerHandler)
	api.GET("/utmtracker/:id", GetUtmTrackerByIDHandler)
	api.PUT("/utmtracker/:id", UpdateUtmTrackerHandler)
	api.DELETE("/utmtracker/:id", DeleteUtmTrackerHandler)

}



