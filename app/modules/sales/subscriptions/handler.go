package subscriptions

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateSubscription godoc
// @Summary Create a new Subscription
// @Description Create a new Subscription in the system
// @Tags sales-subscriptions
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/sales/subscriptions [post]
// @Security BearerAuth
func CreateSubscriptionHandler(c echo.Context) error {
	var data Subscription
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSubscriptionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSubscription godoc
// @Summary Get all Subscription
// @Description Retrieve a list of all Subscription
// @Tags sales-subscriptions
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions [get]
// @Security BearerAuth
func GetAllSubscriptionHandler(c echo.Context) error {
	data, err := GetAllSubscriptionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSubscriptionByID godoc
// @Summary Get a Subscription by ID
// @Description Retrieve a specific Subscription by its ID
// @Tags sales-subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions/{id} [get]
// @Security BearerAuth
func GetSubscriptionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSubscriptionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSubscription godoc
// @Summary Update a Subscription
// @Description Update an existing Subscription
// @Tags sales-subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions/{id} [put]
// @Security BearerAuth
func UpdateSubscriptionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSubscriptionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSubscriptionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSubscription godoc
// @Summary Delete a Subscription
// @Description Delete a Subscription by ID
// @Tags sales-subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions/{id} [delete]
// @Security BearerAuth
func DeleteSubscriptionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSubscriptionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
