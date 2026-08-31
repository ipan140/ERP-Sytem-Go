package whatsapp

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/whatsapp", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreateWaTemplateHandler)
	api.GET("", GetAllWaTemplateHandler)
	api.GET("/:id", GetWaTemplateByIDHandler)
	api.PUT("/:id", UpdateWaTemplateHandler)
	api.DELETE("/:id", DeleteWaTemplateHandler)
}



