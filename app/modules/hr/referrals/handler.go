package referrals

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateReferralReward godoc
// @Summary Create a new ReferralReward
// @Description Create a new ReferralReward in the system
// @Tags hr-ReferralRewards
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/ReferralRewards [post]
// @Security BearerAuth
func CreateReferralRewardHandler(c echo.Context) error {
	var data ReferralReward
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateReferralRewardService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllReferralReward godoc
// @Summary Get all ReferralReward
// @Description Retrieve a list of all ReferralReward
// @Tags hr-ReferralRewards
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/ReferralRewards [get]
// @Security BearerAuth
func GetAllReferralRewardHandler(c echo.Context) error {
	data, err := GetAllReferralRewardService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetReferralRewardByID godoc
// @Summary Get a ReferralReward by ID
// @Description Retrieve a specific ReferralReward by its ID
// @Tags hr-ReferralRewards
// @Produce json
// @Param id path int true "ReferralReward ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/ReferralRewards/{id} [get]
// @Security BearerAuth
func GetReferralRewardByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReferralRewardByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateReferralReward godoc
// @Summary Update a ReferralReward
// @Description Update an existing ReferralReward
// @Tags hr-ReferralRewards
// @Accept json
// @Produce json
// @Param id path int true "ReferralReward ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/ReferralRewards/{id} [put]
// @Security BearerAuth
func UpdateReferralRewardHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetReferralRewardByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateReferralRewardService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteReferralReward godoc
// @Summary Delete a ReferralReward
// @Description Delete a ReferralReward by ID
// @Tags hr-ReferralRewards
// @Produce json
// @Param id path int true "ReferralReward ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/ReferralRewards/{id} [delete]
// @Security BearerAuth
func DeleteReferralRewardHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteReferralRewardService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
