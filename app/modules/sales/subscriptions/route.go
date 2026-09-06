package subscriptions

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/subscriptions", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateSubscriptionHandler)
	api.GET("", GetAllSubscriptionHandler)
	api.GET("/:id", GetSubscriptionByIDHandler)
	api.PUT("/:id", UpdateSubscriptionHandler)
	api.DELETE("/:id", DeleteSubscriptionHandler)
	api.POST("/:id/create-invoice", GenerateSubscriptionInvoiceHandler)

	api.POST("/subscriptionplan", CreateSubscriptionPlanHandler)
	api.GET("/subscriptionplan", GetAllSubscriptionPlanHandler)
	api.GET("/subscriptionplan/:id", GetSubscriptionPlanByIDHandler)
	api.PUT("/subscriptionplan/:id", UpdateSubscriptionPlanHandler)
	api.DELETE("/subscriptionplan/:id", DeleteSubscriptionPlanHandler)
}



