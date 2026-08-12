package whatsapp

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/whatsapp", middleware.Auth())
	api.POST("", CreateWaTemplateHandler)
	api.GET("", GetAllWaTemplateHandler)
	api.GET("/:id", GetWaTemplateByIDHandler)
	api.PUT("/:id", UpdateWaTemplateHandler)
	api.DELETE("/:id", DeleteWaTemplateHandler)
}
