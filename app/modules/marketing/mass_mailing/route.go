package mass_mailing

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/mass_mailing", middleware.Auth())
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
