package marketing_automation

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateAutomationCampaign godoc
// @Summary Create a new AutomationCampaign
// @Description Create a new AutomationCampaign in the system
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Success 201 {object} AutomationCampaign
// @Param request body AutomationCampaign true "Payload"
// @Router /api/marketing/marketing_automation [post]
// @Security BearerAuth
func CreateAutomationCampaignHandler(c echo.Context) error {
	var data AutomationCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAutomationCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllAutomationCampaign godoc
// @Summary Get all AutomationCampaign
// @Description Retrieve a list of all AutomationCampaign
// @Tags marketing-marketing_automation
// @Produce json
// @Success 200 {object} []AutomationCampaign
// @Router /api/marketing/marketing_automation [get]
// @Security BearerAuth
func GetAllAutomationCampaignHandler(c echo.Context) error {
	data, err := GetAllAutomationCampaignService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetAutomationCampaignByID godoc
// @Summary Get a AutomationCampaign by ID
// @Description Retrieve a specific AutomationCampaign by its ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Success 200 {object} AutomationCampaign
// @Router /api/marketing/marketing_automation/{id} [get]
// @Security BearerAuth
func GetAutomationCampaignByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAutomationCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateAutomationCampaign godoc
// @Summary Update a AutomationCampaign
// @Description Update an existing AutomationCampaign
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/marketing_automation/{id} [put]
// @Security BearerAuth
func UpdateAutomationCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAutomationCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAutomationCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteAutomationCampaign godoc
// @Summary Delete a AutomationCampaign
// @Description Delete a AutomationCampaign by ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "AutomationCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/marketing_automation/{id} [delete]
// @Security BearerAuth
func DeleteAutomationCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAutomationCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create WorkflowActivity
// @Description Create a new WorkflowActivity
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Success 201 {object} WorkflowActivity
// @Param request body WorkflowActivity true "Payload"
// @Router /api/marketing/marketing_automation/workflowactivity [post]
// @Security BearerAuth
func CreateWorkflowActivityHandler(c echo.Context) error {
	var data WorkflowActivity
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateWorkflowActivityService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all WorkflowActivity
// @Description Retrieve a list of all WorkflowActivity
// @Tags marketing-marketing_automation
// @Produce json
// @Success 200 {object} WorkflowActivity
// @Router /api/marketing/marketing_automation/workflowactivity [get]
// @Security BearerAuth
func GetAllWorkflowActivityHandler(c echo.Context) error {
	data, err := GetAllWorkflowActivityService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetWorkflowActivityByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkflowActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update WorkflowActivity
// @Description Update an existing WorkflowActivity
// @Tags marketing-marketing_automation
// @Accept json
// @Produce json
// @Param id path int true "WorkflowActivity ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/marketing_automation/workflowactivity/{id} [put]
// @Security BearerAuth
func UpdateWorkflowActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkflowActivityByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateWorkflowActivityService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete WorkflowActivity
// @Description Delete WorkflowActivity by ID
// @Tags marketing-marketing_automation
// @Produce json
// @Param id path int true "WorkflowActivity ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/marketing_automation/workflowactivity/{id} [delete]
// @Security BearerAuth
func DeleteWorkflowActivityHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWorkflowActivityService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}


