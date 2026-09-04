package marketing_automation

import (
	"net/http"
	"strconv"

	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateAutomationCampaign godoc
// @Summary Create a new AutomationCampaign
// @Description Create a new marketing automation campaign workflow scenario
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param request body AutomationCampaign true "Automation Campaign Payload"
// @Success 201 {object} AutomationCampaign
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation [post]
// @Security BearerAuth
func CreateAutomationCampaignHandler(c echo.Context) error {
	var data AutomationCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAutomationCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create automation campaign", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Automation campaign created successfully", data)
}

// GetAllAutomationCampaign godoc
// @Summary Get all AutomationCampaign
// @Description Retrieve a list of all configured marketing automation campaigns
// @Tags marketing-marketing_automation
// @Produce json
// @Success 200 {object} []AutomationCampaign
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation [get]
// @Security BearerAuth
func GetAllAutomationCampaignHandler(c echo.Context) error {
	data, err := GetAllAutomationCampaignService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve automation campaigns", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Automation campaigns retrieved successfully", data)
}

// GetAutomationCampaignByID godoc
// @Summary Get an AutomationCampaign by ID
// @Description Retrieve a specific automation campaign scenario by its ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Success 200 {object} AutomationCampaign
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/{id} [get]
// @Security BearerAuth
func GetAutomationCampaignByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid campaign ID", "ID must be a positive integer")
	}
	data, err := GetAutomationCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Automation campaign not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Automation campaign retrieved successfully", data)
}

// UpdateAutomationCampaign godoc
// @Summary Update an AutomationCampaign
// @Description Update an existing automation campaign scenario
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Param request body AutomationCampaign true "Updated Automation Campaign Payload"
// @Success 200 {object} AutomationCampaign
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/{id} [put]
// @Security BearerAuth
func UpdateAutomationCampaignHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid campaign ID", "ID must be a positive integer")
	}
	data, err := GetAutomationCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Automation campaign not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	data.ID = uint(id)
	if err := UpdateAutomationCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update automation campaign", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Automation campaign updated successfully", data)
}

// DeleteAutomationCampaign godoc
// @Summary Delete an AutomationCampaign
// @Description Delete an automation campaign scenario by ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/{id} [delete]
// @Security BearerAuth
func DeleteAutomationCampaignHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid campaign ID", "ID must be a positive integer")
	}
	if err := DeleteAutomationCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete automation campaign", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Automation campaign deleted successfully", nil)
}

// CreateWorkflowActivityHandler godoc
// @Summary Create WorkflowActivity
// @Description Create a new action step (activity) in an automation campaign workflow
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param request body WorkflowActivity true "Workflow Activity Payload"
// @Success 201 {object} WorkflowActivity
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/workflowactivity [post]
// @Security BearerAuth
func CreateWorkflowActivityHandler(c echo.Context) error {
	var data WorkflowActivity
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateWorkflowActivityService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create workflow activity", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Workflow activity created successfully", data)
}

// GetAllWorkflowActivityHandler godoc
// @Summary Get all WorkflowActivities
// @Description Retrieve a list of all action steps / workflow activities
// @Tags marketing-marketing_automation
// @Produce json
// @Success 200 {object} []WorkflowActivity
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/workflowactivity [get]
// @Security BearerAuth
func GetAllWorkflowActivityHandler(c echo.Context) error {
	data, err := GetAllWorkflowActivityService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve workflow activities", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Workflow activities retrieved successfully", data)
}

// GetWorkflowActivityByIDHandler godoc
// @Summary Get WorkflowActivity by ID
// @Description Retrieve a specific workflow activity step by its ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "WorkflowActivity ID"
// @Success 200 {object} WorkflowActivity
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/workflowactivity/{id} [get]
// @Security BearerAuth
func GetWorkflowActivityByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid activity ID", "ID must be a positive integer")
	}
	data, err := GetWorkflowActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Workflow activity not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Workflow activity retrieved successfully", data)
}

// UpdateWorkflowActivityHandler godoc
// @Summary Update WorkflowActivity
// @Description Update an existing workflow activity action step
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param id path int true "WorkflowActivity ID"
// @Param request body WorkflowActivity true "Updated Workflow Activity Payload"
// @Success 200 {object} WorkflowActivity
// @Failure 400 {object} utils.ErrorResponse
// @Failure 404 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/workflowactivity/{id} [put]
// @Security BearerAuth
func UpdateWorkflowActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid activity ID", "ID must be a positive integer")
	}
	data, err := GetWorkflowActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Workflow activity not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	data.ID = uint(id)
	if err := UpdateWorkflowActivityService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update workflow activity", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Workflow activity updated successfully", data)
}

// DeleteWorkflowActivityHandler godoc
// @Summary Delete WorkflowActivity
// @Description Delete a workflow activity action step by ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "WorkflowActivity ID"
// @Success 200 {object} utils.SuccessResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/marketing/marketing_automation/workflowactivity/{id} [delete]
// @Security BearerAuth
func DeleteWorkflowActivityHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Invalid activity ID", "ID must be a positive integer")
	}
	if err := DeleteWorkflowActivityService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete workflow activity", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Workflow activity deleted successfully", nil)
}


