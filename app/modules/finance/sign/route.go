package sign

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/sign", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateSignatureRequestHandler)
	api.GET("", GetAllSignatureRequestHandler)
	api.GET("/:id", GetSignatureRequestByIDHandler)
	api.PUT("/:id", UpdateSignatureRequestHandler)
	api.DELETE("/:id", DeleteSignatureRequestHandler)
}



