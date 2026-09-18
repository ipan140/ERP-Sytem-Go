package voip

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	for _, path := range []string{"/api/voip", "/api/core/voip"} {
		api := e.Group(path, middleware.Auth(), middleware.GlobalAutoRBAC())
		api.GET("/extensions", GetVoipExtensionsHandler)
		api.POST("/extensions", CreateVoipExtensionHandler)
		api.POST("/call", InitiateCallHandler)
		api.POST("", CreateCallRecordHandler)
		api.GET("", GetAllCallRecordHandler)
		api.GET("/:id", GetCallRecordByIDHandler)
		api.PUT("/:id", UpdateCallRecordHandler)
		api.DELETE("/:id", DeleteCallRecordHandler)
	}
}



