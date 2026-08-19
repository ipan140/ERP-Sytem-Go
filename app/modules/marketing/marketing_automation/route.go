package marketing_automation

import (
	"github.com/labstack/echo/v4"
	"ERP-System/common/middleware"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/marketing_automation", middleware.Auth())
	api.POST("", CreateAutomationCampaignHandler)
	api.GET("", GetAllAutomationCampaignHandler)
	api.GET("/:id", GetAutomationCampaignByIDHandler)
	api.PUT("/:id", UpdateAutomationCampaignHandler)
	api.DELETE("/:id", DeleteAutomationCampaignHandler)

	api.POST("/workflowactivity", CreateWorkflowActivityHandler)
	api.GET("/workflowactivity", GetAllWorkflowActivityHandler)
	api.GET("/workflowactivity/:id", GetWorkflowActivityByIDHandler)
	api.PUT("/workflowactivity/:id", UpdateWorkflowActivityHandler)
	api.DELETE("/workflowactivity/:id", DeleteWorkflowActivityHandler)

}
