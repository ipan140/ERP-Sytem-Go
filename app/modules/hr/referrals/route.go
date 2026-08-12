package referrals

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/hr/referrals", middleware.Auth())
	api.POST("", CreateReferralHandler)
	api.GET("", GetAllReferralHandler)
	api.GET("/:id", GetReferralByIDHandler)
	api.PUT("/:id", UpdateReferralHandler)
	api.DELETE("/:id", DeleteReferralHandler)
}
