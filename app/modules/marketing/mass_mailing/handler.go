package mass_mailing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

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

// @Summary Create UtmTracker
// @Description Create a new UtmTracker
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker [post]
// @Security BearerAuth
func CreateUtmTrackerHandler(c echo.Context) error {
	var data UtmTracker
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateUtmTrackerService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all UtmTracker
// @Description Retrieve a list of all UtmTracker
// @Tags marketing-mass_mailing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker [get]
// @Security BearerAuth
func GetAllUtmTrackerHandler(c echo.Context) error {
	data, err := GetAllUtmTrackerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetUtmTrackerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUtmTrackerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update UtmTracker
// @Description Update an existing UtmTracker
// @Tags marketing-mass_mailing
// @Accept json
// @Produce json
// @Param id path int true "UtmTracker ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker/{id} [put]
// @Security BearerAuth
func UpdateUtmTrackerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetUtmTrackerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateUtmTrackerService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete UtmTracker
// @Description Delete UtmTracker by ID
// @Tags marketing-mass_mailing
// @Produce json
// @Param id path int true "UtmTracker ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/mass_mailing/utmtracker/{id} [delete]
// @Security BearerAuth
func DeleteUtmTrackerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteUtmTrackerService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}
