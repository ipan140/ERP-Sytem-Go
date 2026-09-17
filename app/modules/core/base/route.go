package base

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/base", middleware.Auth(), middleware.GlobalAutoRBAC())
	// Currencies
	api.POST("/currency/sync-bi", SyncBankIndonesiaRatesHandler)
	api.GET("/currency", GetAllCurrencyHandler)
	api.POST("/currency", CreateCurrencyHandler)
	api.GET("/currency/:id", GetCurrencyByIDHandler)
	api.PUT("/currency/:id", UpdateCurrencyHandler)
	api.DELETE("/currency/:id", DeleteCurrencyHandler)
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



