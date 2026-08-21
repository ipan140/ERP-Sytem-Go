package base

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/base", middleware.Auth(), middleware.RequireRoles(constants.RoleDirector, constants.RoleEmployee))
	api.POST("", CreateCurrencyHandler)
	api.GET("", GetAllCurrencyHandler)
	api.GET("/:id", GetCurrencyByIDHandler)
	api.PUT("/:id", UpdateCurrencyHandler)
	api.DELETE("/:id", DeleteCurrencyHandler)

	api.POST("/country", CreateCountryHandler)
	api.GET("/country", GetAllCountryHandler)
	api.GET("/country/:id", GetCountryByIDHandler)
	api.PUT("/country/:id", UpdateCountryHandler)
	api.DELETE("/country/:id", DeleteCountryHandler)

	api.POST("/countrystate", CreateCountryStateHandler)
	api.GET("/countrystate", GetAllCountryStateHandler)
	api.GET("/countrystate/:id", GetCountryStateByIDHandler)
	api.PUT("/countrystate/:id", UpdateCountryStateHandler)
	api.DELETE("/countrystate/:id", DeleteCountryStateHandler)

	api.POST("/partner", CreatePartnerHandler)
	api.GET("/partner", GetAllPartnerHandler)
	api.GET("/partner/:id", GetPartnerByIDHandler)
	api.PUT("/partner/:id", UpdatePartnerHandler)
	api.DELETE("/partner/:id", DeletePartnerHandler)
}
