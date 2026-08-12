package main

import (
	"log"
	"net/http"

	"ERP-System/app/modules/core/artificial_intelligence"
	"ERP-System/app/modules/core/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/core/dashboards"
	"ERP-System/app/modules/core/discuss"
	"ERP-System/app/modules/core/documents"
	"ERP-System/app/modules/core/iot"
	"ERP-System/app/modules/core/knowledge"
	"ERP-System/app/modules/core/mailer"
	_ "ERP-System/app/modules/core/report"
	"ERP-System/app/modules/core/storage"
	"ERP-System/app/modules/core/voip"
	"ERP-System/app/modules/core/whatsapp"
	"ERP-System/app/modules/finance/accounting"
	"ERP-System/app/modules/finance/approvals"
	"ERP-System/app/modules/finance/consolidation"
	financeDocs "ERP-System/app/modules/finance/documents"
	"ERP-System/app/modules/finance/expenses"
	"ERP-System/app/modules/finance/invoicing"
	"ERP-System/app/modules/finance/sign"
	"ERP-System/app/modules/finance/spreadsheet_bi"
	"ERP-System/app/modules/hr/appraisals"
	"ERP-System/app/modules/hr/attendances"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/app/modules/hr/fleet"
	"ERP-System/app/modules/hr/lunch"
	"ERP-System/app/modules/hr/payroll"
	"ERP-System/app/modules/hr/recruitment"
	"ERP-System/app/modules/hr/referrals"
	"ERP-System/app/modules/hr/time_off"
	"ERP-System/app/modules/sales/crm"
	"ERP-System/app/modules/sales/point_of_sale"
	"ERP-System/app/modules/sales/rental"
	"ERP-System/app/modules/sales/sales_core"
	"ERP-System/app/modules/sales/subscriptions"
	"ERP-System/config"
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
	accounting.RegisterRoutes(e)
	approvals.RegisterRoutes(e)
	consolidation.RegisterRoutes(e)
	financeDocs.RegisterRoutes(e)
	expenses.RegisterRoutes(e)
	invoicing.RegisterRoutes(e)
	sign.RegisterRoutes(e)
	spreadsheet_bi.RegisterRoutes(e)
	appraisals.RegisterRoutes(e)
	attendances.RegisterRoutes(e)
	employees.RegisterRoutes(e)
	fleet.RegisterRoutes(e)
	lunch.RegisterRoutes(e)
	payroll.RegisterRoutes(e)
	recruitment.RegisterRoutes(e)
	referrals.RegisterRoutes(e)
	time_off.RegisterRoutes(e)
	crm.RegisterRoutes(e)
	point_of_sale.RegisterRoutes(e)
	rental.RegisterRoutes(e)
	sales_core.RegisterRoutes(e)
	subscriptions.RegisterRoutes(e)

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
