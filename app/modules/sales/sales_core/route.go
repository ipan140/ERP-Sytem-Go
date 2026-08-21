package sales_core

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/sales/sales_core", middleware.Auth())
	api.POST("", CreateSaleOrderHandler)
	api.GET("", GetAllSaleOrderHandler)
	api.GET("/:id", GetSaleOrderByIDHandler)
	api.PUT("/:id", UpdateSaleOrderHandler)
	api.DELETE("/:id", DeleteSaleOrderHandler)

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
