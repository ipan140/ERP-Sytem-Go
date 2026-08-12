package voip

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/voip", middleware.Auth())
	api.POST("", CreateCallRecordHandler)
	api.GET("", GetAllCallRecordHandler)
	api.GET("/:id", GetCallRecordByIDHandler)
	api.PUT("/:id", UpdateCallRecordHandler)
	api.DELETE("/:id", DeleteCallRecordHandler)
}
