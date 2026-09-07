package sales_core

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	setupRoutes := func(api *echo.Group) {
		api.POST("", CreateSaleOrderHandler)
		api.GET("", GetAllSaleOrderHandler)
		api.GET("/:id", GetSaleOrderByIDHandler)
		api.PUT("/:id", UpdateSaleOrderHandler)
		api.DELETE("/:id", DeleteSaleOrderHandler)
		api.POST("/:id/confirm", ConfirmSaleOrderHandler)
		api.PUT("/:id/status", UpdateSaleOrderStatusHandler)
		api.GET("/:id/print", GenerateQuotationDocHandler)
		api.POST("/:id/create-invoice", CreateInvoiceFromSaleOrderHandler)
		api.POST("/:id/payment-link", GeneratePaymentLinkHandler)
		api.POST("/:id/approve-discount", ApproveDiscountHandler)
		api.POST("/:id/bypass-credit", BypassCreditHoldHandler)
		api.GET("/:id/export-efaktur", ExportEFakturCSVHandler)
		api.GET("/:id/deliveries", GetDeliveryOrdersHandler)
		api.POST("/:id/deliver", DeliverSaleOrderHandler)
		api.POST("/:id/sign", SignSaleOrderHandler)
		api.GET("/leaderboard", GetSalesLeaderboardHandler)

		api.POST("/pricelist", CreatePricelistHandler)
		api.GET("/pricelist", GetAllPricelistHandler)
		api.GET("/pricelist/:id", GetPricelistByIDHandler)
		api.PUT("/pricelist/:id", UpdatePricelistHandler)
		api.DELETE("/pricelist/:id", DeletePricelistHandler)

		api.POST("/pricelistitem", CreatePricelistItemHandler)
		api.GET("/pricelistitem", GetAllPricelistItemHandler)
		api.GET("/pricelistitem/:id", GetPricelistItemByIDHandler)
		api.PUT("/pricelistitem/:id", UpdatePricelistItemHandler)
		api.DELETE("/pricelistitem/:id", DeletePricelistItemHandler)

		api.POST("/quotationtemplate", CreateQuotationTemplateHandler)
		api.GET("/quotationtemplate", GetAllQuotationTemplateHandler)
		api.GET("/quotationtemplate/:id", GetQuotationTemplateByIDHandler)
		api.PUT("/quotationtemplate/:id", UpdateQuotationTemplateHandler)
		api.DELETE("/quotationtemplate/:id", DeleteQuotationTemplateHandler)

		api.POST("/deliverymethod", CreateDeliveryMethodHandler)
		api.GET("/deliverymethod", GetAllDeliveryMethodHandler)
		api.GET("/deliverymethod/:id", GetDeliveryMethodByIDHandler)
		api.PUT("/deliverymethod/:id", UpdateDeliveryMethodHandler)
		api.DELETE("/deliverymethod/:id", DeleteDeliveryMethodHandler)

		api.POST("/saleorderline", CreateSaleOrderLineHandler)
		api.GET("/saleorderline", GetAllSaleOrderLineHandler)
		api.GET("/saleorderline/:id", GetSaleOrderLineByIDHandler)
		api.PUT("/saleorderline/:id", UpdateSaleOrderLineHandler)
		api.DELETE("/saleorderline/:id", DeleteSaleOrderLineHandler)
	}

	apiCore := e.Group("/api/sales/sales_core", middleware.Auth(), middleware.GlobalAutoRBAC())
	setupRoutes(apiCore)

	apiAlias := e.Group("/api/sales/core", middleware.Auth(), middleware.GlobalAutoRBAC())
	setupRoutes(apiAlias)
}



