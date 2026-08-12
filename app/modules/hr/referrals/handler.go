package referrals

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateReferral godoc
// @Summary Create a new Referral
// @Description Create a new Referral in the system
// @Tags hr-referrals
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/referrals [post]
// @Security BearerAuth
func CreateReferralHandler(c echo.Context) error {
	var data Referral
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateReferralService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllReferral godoc
// @Summary Get all Referral
// @Description Retrieve a list of all Referral
// @Tags hr-referrals
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/referrals [get]
// @Security BearerAuth
func GetAllReferralHandler(c echo.Context) error {
	data, err := GetAllReferralService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetReferralByID godoc
// @Summary Get a Referral by ID
// @Description Retrieve a specific Referral by its ID
// @Tags hr-referrals
// @Produce json
// @Param id path int true "Referral ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/referrals/{id} [get]
// @Security BearerAuth
func GetReferralByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReferralByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateReferral godoc
// @Summary Update a Referral
// @Description Update an existing Referral
// @Tags hr-referrals
// @Accept json
// @Produce json
// @Param id path int true "Referral ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/referrals/{id} [put]
// @Security BearerAuth
func UpdateReferralHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReferralByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateReferralService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteReferral godoc
// @Summary Delete a Referral
// @Description Delete a Referral by ID
// @Tags hr-referrals
// @Produce json
// @Param id path int true "Referral ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/referrals/{id} [delete]
// @Security BearerAuth
func DeleteReferralHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteReferralService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
