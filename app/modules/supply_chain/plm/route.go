package plm

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/supply_chain/plm", middleware.Auth(), middleware.GlobalAutoRBAC())

	// Specific endpoints first
	api.GET("/summary", GetPlmSummaryHandler)
	api.GET("/type", GetAllEcoTypesHandler)

	// CRUD & Workflow
	api.POST("", CreateEcoHandler)
	api.GET("", GetAllEcoHandler)
	api.GET("/:id", GetEcoByIDHandler)
	api.PUT("/:id/state", UpdateEcoStateHandler)
	api.DELETE("/:id", DeleteEcoHandler)
}
