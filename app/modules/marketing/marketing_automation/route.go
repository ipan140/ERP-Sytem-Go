package marketing_automation

import (
	"ERP-System/common/constants"
	"ERP-System/common/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/marketing/marketing_automation", middleware.Auth(), middleware.RequireRoles(constants.RoleMarketingManager))
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
