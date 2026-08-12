package sign

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/finance/sign", middleware.Auth())
	api.POST("", CreateSignatureRequestHandler)
	api.GET("", GetAllSignatureRequestHandler)
	api.GET("/:id", GetSignatureRequestByIDHandler)
	api.PUT("/:id", UpdateSignatureRequestHandler)
	api.DELETE("/:id", DeleteSignatureRequestHandler)
}
