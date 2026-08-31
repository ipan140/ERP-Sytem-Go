package purchase

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/purchase", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreatePurchaseOrderHandler)
	api.GET("", GetAllPurchaseOrderHandler)
	api.GET("/:id", GetPurchaseOrderByIDHandler)
	api.PUT("/:id", UpdatePurchaseOrderHandler)
	api.DELETE("/:id", DeletePurchaseOrderHandler)

	api.POST("/purchaserequisition", CreatePurchaseRequisitionHandler)
	api.GET("/purchaserequisition", GetAllPurchaseRequisitionHandler)
	api.GET("/purchaserequisition/:id", GetPurchaseRequisitionByIDHandler)
	api.PUT("/purchaserequisition/:id", UpdatePurchaseRequisitionHandler)
	api.DELETE("/purchaserequisition/:id", DeletePurchaseRequisitionHandler)

	api.POST("/productsupplierinfo", CreateProductSupplierInfoHandler)
	api.GET("/productsupplierinfo", GetAllProductSupplierInfoHandler)
	api.GET("/productsupplierinfo/:id", GetProductSupplierInfoByIDHandler)
	api.PUT("/productsupplierinfo/:id", UpdateProductSupplierInfoHandler)
	api.DELETE("/productsupplierinfo/:id", DeleteProductSupplierInfoHandler)

	api.POST("/purchaseorderline", CreatePurchaseOrderLineHandler)
	api.GET("/purchaseorderline", GetAllPurchaseOrderLineHandler)
	api.GET("/purchaseorderline/:id", GetPurchaseOrderLineByIDHandler)
	api.PUT("/purchaseorderline/:id", UpdatePurchaseOrderLineHandler)
	api.DELETE("/purchaseorderline/:id", DeletePurchaseOrderLineHandler)
}



