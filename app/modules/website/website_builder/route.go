package website_builder

import (
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Static file serving for uploads (resumes & documents)
	e.Static("/uploads", "uploads")

	// Public & Corporate Portal Endpoints
	e.GET("/api/website/careers", GetCareersHandler)
	e.POST("/api/website/careers/upload-cv", UploadResumeHandler)
	e.POST("/api/website/careers/apply", ApplyCareerHandler)
	e.GET("/api/website/announcements", GetAnnouncementsHandler)
	e.GET("/api/website/stats", GetEnterpriseStatsHandler)

	// Fase 1: Customer & Vendor Self-Service Portal
	e.GET("/api/website/partner-portal", GetPartnerPortalDataHandler)

	// Fase 2: Whistleblowing System (WBS)
	e.POST("/api/website/whistleblowing", SubmitWBSHandler)
	e.GET("/api/website/whistleblowing/:ticket", GetWBSStatusHandler)

	// Fase 3: E-Commerce Midtrans Snap
	e.POST("/api/website/ecommerce/snap", CreateSnapTransactionHandler)

	// Protected CMS Admin Endpoints
	api := e.Group("/api/website/website_builder", middleware.Auth(), middleware.GlobalAutoRBAC())
	api.POST("", CreatePageHandler)
	api.GET("", GetAllPageHandler)
	api.GET("/:id", GetPageByIDHandler)
	api.PUT("/:id", UpdatePageHandler)
	api.DELETE("/:id", DeletePageHandler)
	api.POST("/announcements", CreateAnnouncementHandler)
}



