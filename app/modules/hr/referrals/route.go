package referrals

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/ReferralRewards", middleware.Auth(), middleware.GlobalAutoRBAC())
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



