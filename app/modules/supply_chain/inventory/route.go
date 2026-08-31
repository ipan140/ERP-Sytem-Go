package inventory

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/inventory", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateProductHandler)
	api.GET("", GetAllProductHandler)
	api.GET("/:id", GetProductByIDHandler)
	api.PUT("/:id", UpdateProductHandler)
	api.DELETE("/:id", DeleteProductHandler)

	api.POST("/productcategory", CreateProductCategoryHandler)
	api.GET("/productcategory", GetAllProductCategoryHandler)
	api.GET("/productcategory/:id", GetProductCategoryByIDHandler)
	api.PUT("/productcategory/:id", UpdateProductCategoryHandler)
	api.DELETE("/productcategory/:id", DeleteProductCategoryHandler)

	api.POST("/uomcategory", CreateUoMCategoryHandler)
	api.GET("/uomcategory", GetAllUoMCategoryHandler)
	api.GET("/uomcategory/:id", GetUoMCategoryByIDHandler)
	api.PUT("/uomcategory/:id", UpdateUoMCategoryHandler)
	api.DELETE("/uomcategory/:id", DeleteUoMCategoryHandler)

	api.POST("/uom", CreateUoMHandler)
	api.GET("/uom", GetAllUoMHandler)
	api.GET("/uom/:id", GetUoMByIDHandler)
	api.PUT("/uom/:id", UpdateUoMHandler)
	api.DELETE("/uom/:id", DeleteUoMHandler)

	api.POST("/producttemplate", CreateProductTemplateHandler)
	api.GET("/producttemplate", GetAllProductTemplateHandler)
	api.GET("/producttemplate/:id", GetProductTemplateByIDHandler)
	api.PUT("/producttemplate/:id", UpdateProductTemplateHandler)
	api.DELETE("/producttemplate/:id", DeleteProductTemplateHandler)

	api.POST("/productattribute", CreateProductAttributeHandler)
	api.GET("/productattribute", GetAllProductAttributeHandler)
	api.GET("/productattribute/:id", GetProductAttributeByIDHandler)
	api.PUT("/productattribute/:id", UpdateProductAttributeHandler)
	api.DELETE("/productattribute/:id", DeleteProductAttributeHandler)

	api.POST("/productattributevalue", CreateProductAttributeValueHandler)
	api.GET("/productattributevalue", GetAllProductAttributeValueHandler)
	api.GET("/productattributevalue/:id", GetProductAttributeValueByIDHandler)
	api.PUT("/productattributevalue/:id", UpdateProductAttributeValueHandler)
	api.DELETE("/productattributevalue/:id", DeleteProductAttributeValueHandler)

	api.POST("/stockwarehouse", CreateStockWarehouseHandler)
	api.GET("/stockwarehouse", GetAllStockWarehouseHandler)
	api.GET("/stockwarehouse/:id", GetStockWarehouseByIDHandler)
	api.PUT("/stockwarehouse/:id", UpdateStockWarehouseHandler)
	api.DELETE("/stockwarehouse/:id", DeleteStockWarehouseHandler)

	api.POST("/stocklocation", CreateStockLocationHandler)
	api.GET("/stocklocation", GetAllStockLocationHandler)
	api.GET("/stocklocation/:id", GetStockLocationByIDHandler)
	api.PUT("/stocklocation/:id", UpdateStockLocationHandler)
	api.DELETE("/stocklocation/:id", DeleteStockLocationHandler)

	api.POST("/stockpicking", CreateStockPickingHandler)
	api.GET("/stockpicking", GetAllStockPickingHandler)
	api.GET("/stockpicking/:id", GetStockPickingByIDHandler)
	api.PUT("/stockpicking/:id", UpdateStockPickingHandler)
	api.DELETE("/stockpicking/:id", DeleteStockPickingHandler)

	api.POST("/stocklot", CreateStockLotHandler)
	api.GET("/stocklot", GetAllStockLotHandler)
	api.GET("/stocklot/:id", GetStockLotByIDHandler)
	api.PUT("/stocklot/:id", UpdateStockLotHandler)
	api.DELETE("/stocklot/:id", DeleteStockLotHandler)

	api.POST("/stockquant", CreateStockQuantHandler)
	api.GET("/stockquant", GetAllStockQuantHandler)
	api.GET("/stockquant/:id", GetStockQuantByIDHandler)
	api.PUT("/stockquant/:id", UpdateStockQuantHandler)
	api.DELETE("/stockquant/:id", DeleteStockQuantHandler)

	api.POST("/stockputawayrule", CreateStockPutawayRuleHandler)
	api.GET("/stockputawayrule", GetAllStockPutawayRuleHandler)
	api.GET("/stockputawayrule/:id", GetStockPutawayRuleByIDHandler)
	api.PUT("/stockputawayrule/:id", UpdateStockPutawayRuleHandler)
	api.DELETE("/stockputawayrule/:id", DeleteStockPutawayRuleHandler)

	api.POST("/stockvaluationlayer", CreateStockValuationLayerHandler)
	api.GET("/stockvaluationlayer", GetAllStockValuationLayerHandler)
	api.GET("/stockvaluationlayer/:id", GetStockValuationLayerByIDHandler)
	api.PUT("/stockvaluationlayer/:id", UpdateStockValuationLayerHandler)
	api.DELETE("/stockvaluationlayer/:id", DeleteStockValuationLayerHandler)
}



