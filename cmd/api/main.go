package main

import (
	"log"
	"net/http"
	
	"ERP-System/config"
	"ERP-System/app/modules/core/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/core/dashboards"
	"ERP-System/app/modules/core/discuss"
	"ERP-System/app/modules/core/documents"
	"ERP-System/app/modules/core/iot"
	"ERP-System/app/modules/core/knowledge"
	"ERP-System/app/modules/core/storage"
	"ERP-System/app/modules/core/voip"
	"ERP-System/app/modules/core/whatsapp"
	"ERP-System/app/modules/core/artificial_intelligence"
	"ERP-System/app/modules/core/mailer"
	_ "ERP-System/app/modules/core/report"
	_ "ERP-System/docs" // Swagger docs

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Golang ERP System API
// @version 1.0
// @description This is the core API for Odoo-style Golang ERP.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Initialize configuration and database
	config.LoadEnv()
	config.ConnectDB()

	e := echo.New()

	// Global Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Register Module Routes
	auth.RegisterRoutes(e)
	base.RegisterRoutes(e)
	discuss.RegisterRoutes(e)
	documents.RegisterRoutes(e)
	dashboards.RegisterRoutes(e)
	iot.RegisterRoutes(e)
	knowledge.RegisterRoutes(e)
	voip.RegisterRoutes(e)
	whatsapp.RegisterRoutes(e)
	artificial_intelligence.RegisterRoutes(e)
	storage.RegisterRoutes(e)
	mailer.RegisterRoutes(e)

	// Register Swagger Route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to Golang ERP System (Odoo Core Engine)",
			"status":  "running",
		})
	})
	// Register Module Routes
	auth.RegisterRoutes(e)
	storage.RegisterRoutes(e)

	// 5. Start Server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
