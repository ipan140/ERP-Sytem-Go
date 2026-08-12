package main

import (
	"log"
	"net/http"
	
	"ERP-System/config"
	"ERP-System/app/modules/core/auth"
	"ERP-System/app/modules/core/storage"
	_ "ERP-System/app/modules/core/mailer"
	_ "ERP-System/app/modules/core/report"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Load configuration and connect to database
	config.LoadEnv()
	config.ConnectDB()

	// 2. Initialize Echo instance
	e := echo.New()

	// 3. Add global middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 4. Setup Routes
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
