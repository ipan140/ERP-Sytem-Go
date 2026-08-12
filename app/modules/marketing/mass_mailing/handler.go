package mass_mailing

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateMailingCampaign godoc
// @Summary Create a new MailingCampaign
// @Description Create a new MailingCampaign in the system
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing [post]
// @Security BearerAuth
func CreateMailingCampaignHandler(c echo.Context) error {
	var data MailingCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateMailingCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllMailingCampaign godoc
// @Summary Get all MailingCampaign
// @Description Retrieve a list of all MailingCampaign
// @Tags marketing-mass_mailing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing [get]
// @Security BearerAuth
func GetAllMailingCampaignHandler(c echo.Context) error {
	data, err := GetAllMailingCampaignService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetMailingCampaignByID godoc
// @Summary Get a MailingCampaign by ID
// @Description Retrieve a specific MailingCampaign by its ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id} [get]
// @Security BearerAuth
func GetMailingCampaignByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateMailingCampaign godoc
// @Summary Update a MailingCampaign
// @Description Update an existing MailingCampaign
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id} [put]
// @Security BearerAuth
func UpdateMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetMailingCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateMailingCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteMailingCampaign godoc
// @Summary Delete a MailingCampaign
// @Description Delete a MailingCampaign by ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "MailingCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/{id} [delete]
// @Security BearerAuth
func DeleteMailingCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteMailingCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
