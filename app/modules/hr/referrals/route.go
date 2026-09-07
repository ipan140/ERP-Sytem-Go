package referrals

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	setupRoutes := func(api *echo.Group) {
		api.POST("", CreateReferralRewardHandler)
		api.GET("", GetAllReferralRewardHandler)
		api.GET("/:id", GetReferralRewardByIDHandler)
		api.PUT("/:id", UpdateReferralRewardHandler)
		api.DELETE("/:id", DeleteReferralRewardHandler)

		api.POST("/referralpoint", CreateReferralPointHandler)
		api.GET("/referralpoint", GetAllReferralPointHandler)
		api.GET("/referralpoint/:id", GetReferralPointByIDHandler)
		api.PUT("/referralpoint/:id", UpdateReferralPointHandler)
		api.DELETE("/referralpoint/:id", DeleteReferralPointHandler)
	}

	apiLower := e.Group("/api/hr/referrals", middleware.Auth(), middleware.GlobalAutoRBAC())
	setupRoutes(apiLower)

	apiUpper := e.Group("/api/hr/ReferralRewards", middleware.Auth(), middleware.GlobalAutoRBAC())
	setupRoutes(apiUpper)
}



