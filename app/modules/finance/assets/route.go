package assets

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/assets", middleware.Auth(), middleware.GlobalAutoRBAC())
	// Assets
	api.GET("", GetAllAssetsHandler)
	api.POST("", CreateAssetHandler)
	api.PUT("/:id", UpdateAssetHandler)
	api.DELETE("/:id", DeleteAssetHandler)
	api.POST("/depreciate", PostMonthlyDepreciationHandler)

	// Categories
	api.GET("/categories", GetAllCategoriesHandler)
	api.POST("/categories", CreateCategoryHandler)
	api.PUT("/categories/:id", UpdateCategoryHandler)
	api.DELETE("/categories/:id", DeleteCategoryHandler)
}
