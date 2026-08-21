package ecommerce

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/website/ecommerce", middleware.Auth())
	api.POST("", CreateCartHandler)
	api.GET("", GetAllCartHandler)
	api.GET("/:id", GetCartByIDHandler)
	api.PUT("/:id", UpdateCartHandler)
	api.DELETE("/:id", DeleteCartHandler)

	api.POST("/portaluser", CreatePortalUserHandler)
	api.GET("/portaluser", GetAllPortalUserHandler)
	api.GET("/portaluser/:id", GetPortalUserByIDHandler)
	api.PUT("/portaluser/:id", UpdatePortalUserHandler)
	api.DELETE("/portaluser/:id", DeletePortalUserHandler)
	api.POST("/shoppingcart", CreateShoppingCartHandler)
	api.GET("/shoppingcart", GetAllShoppingCartHandler)
	api.GET("/shoppingcart/:id", GetShoppingCartByIDHandler)
	api.PUT("/shoppingcart/:id", UpdateShoppingCartHandler)
	api.DELETE("/shoppingcart/:id", DeleteShoppingCartHandler)
	api.POST("/cartitem", CreateCartItemHandler)
	api.GET("/cartitem", GetAllCartItemHandler)
	api.GET("/cartitem/:id", GetCartItemByIDHandler)
	api.PUT("/cartitem/:id", UpdateCartItemHandler)
	api.DELETE("/cartitem/:id", DeleteCartItemHandler)

}
