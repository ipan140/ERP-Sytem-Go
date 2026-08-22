package sms_marketing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSmsCampaign godoc
// @Summary Create a new SmsCampaign
// @Description Create a new SmsCampaign in the system
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Success 201 {object} SmsCampaign
// @Param request body SmsCampaign true "Payload"
// @Router /api/marketing/sms_marketing [post]
// @Security BearerAuth
func CreateSmsCampaignHandler(c echo.Context) error {
	var data SmsCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSmsCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSmsCampaign godoc
// @Summary Get all SmsCampaign
// @Description Retrieve a list of all SmsCampaign
// @Tags marketing-sms_marketing
// @Produce json
// @Success 200 {object} []SmsCampaign
// @Router /api/marketing/sms_marketing [get]
// @Security BearerAuth
func GetAllSmsCampaignHandler(c echo.Context) error {
	data, err := GetAllSmsCampaignService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSmsCampaignByID godoc
// @Summary Get a SmsCampaign by ID
// @Description Retrieve a specific SmsCampaign by its ID
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} SmsCampaign
// @Router /api/marketing/sms_marketing/{id} [get]
// @Security BearerAuth
func GetSmsCampaignByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSmsCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSmsCampaign godoc
// @Summary Update a SmsCampaign
// @Description Update an existing SmsCampaign
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/sms_marketing/{id} [put]
// @Security BearerAuth
func UpdateSmsCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSmsCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSmsCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSmsCampaign godoc
// @Summary Delete a SmsCampaign
// @Description Delete a SmsCampaign by ID
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/marketing/sms_marketing/{id} [delete]
// @Security BearerAuth
func DeleteSmsCampaignHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSmsCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


