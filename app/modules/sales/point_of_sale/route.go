package point_of_sale

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/point_of_sale", middleware.Auth())
	api.POST("", CreatePosSessionHandler)
	api.GET("", GetAllPosSessionHandler)
	api.GET("/:id", GetPosSessionByIDHandler)
	api.PUT("/:id", UpdatePosSessionHandler)
	api.DELETE("/:id", DeletePosSessionHandler)

	api.POST("/posconfig", CreatePosConfigHandler)
	api.GET("/posconfig", GetAllPosConfigHandler)
	api.GET("/posconfig/:id", GetPosConfigByIDHandler)
	api.PUT("/posconfig/:id", UpdatePosConfigHandler)
	api.DELETE("/posconfig/:id", DeletePosConfigHandler)

	api.POST("/posorder", CreatePosOrderHandler)
	api.GET("/posorder", GetAllPosOrderHandler)
	api.GET("/posorder/:id", GetPosOrderByIDHandler)
	api.PUT("/posorder/:id", UpdatePosOrderHandler)
	api.DELETE("/posorder/:id", DeletePosOrderHandler)

	api.POST("/posorderline", CreatePosOrderLineHandler)
	api.GET("/posorderline", GetAllPosOrderLineHandler)
	api.GET("/posorderline/:id", GetPosOrderLineByIDHandler)
	api.PUT("/posorderline/:id", UpdatePosOrderLineHandler)
	api.DELETE("/posorderline/:id", DeletePosOrderLineHandler)

	api.POST("/pospayment", CreatePosPaymentHandler)
	api.GET("/pospayment", GetAllPosPaymentHandler)
	api.GET("/pospayment/:id", GetPosPaymentByIDHandler)
	api.PUT("/pospayment/:id", UpdatePosPaymentHandler)
	api.DELETE("/pospayment/:id", DeletePosPaymentHandler)

	api.POST("/loyaltyprogram", CreateLoyaltyProgramHandler)
	api.GET("/loyaltyprogram", GetAllLoyaltyProgramHandler)
	api.GET("/loyaltyprogram/:id", GetLoyaltyProgramByIDHandler)
	api.PUT("/loyaltyprogram/:id", UpdateLoyaltyProgramHandler)
	api.DELETE("/loyaltyprogram/:id", DeleteLoyaltyProgramHandler)
}
