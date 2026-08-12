package subscriptions

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/subscriptions", middleware.Auth())
	api.POST("", CreateSubscriptionHandler)
	api.GET("", GetAllSubscriptionHandler)
	api.GET("/:id", GetSubscriptionByIDHandler)
	api.PUT("/:id", UpdateSubscriptionHandler)
	api.DELETE("/:id", DeleteSubscriptionHandler)
}
