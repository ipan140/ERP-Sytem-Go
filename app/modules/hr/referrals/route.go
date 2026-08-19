package referrals

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/ReferralRewards", middleware.Auth())
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
