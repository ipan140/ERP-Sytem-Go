package subscriptions

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSubscription godoc
// @Summary Create a new Subscription
// @Description Create a new Subscription in the system
// @Tags sales-subscriptions
// @Accept json
// @Produce json
// @Success 201 {object} Subscription
// @Param request body Subscription true "Payload"
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
// @Success 200 {object} []Subscription
// @Router /api/sales/subscriptions [get]
// @Security BearerAuth
func GetAllSubscriptionHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllSubscriptionService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	state := c.QueryParam("state")

	data, total, err := GetPaginatedSubscriptionService(offset, limit, search, state)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetSubscriptionByID godoc
// @Summary Get a Subscription by ID
// @Description Retrieve a specific Subscription by its ID
// @Tags sales-subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} Subscription
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

// GenerateSubscriptionInvoiceHandler godoc
// @Summary Generate Recurring Invoice for Subscription
// @Description Trigger one-click recurring customer invoice in Finance Invoicing and advance next invoice date
// @Tags sales-subscriptions
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 201 {object} invoicing.Invoice
// @Router /api/sales/subscriptions/{id}/create-invoice [post]
// @Security BearerAuth
func GenerateSubscriptionInvoiceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	inv, err := GenerateSubscriptionInvoiceService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to generate subscription invoice", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Recurring invoice created successfully in Finance", inv)
}

// @Summary Create SubscriptionPlan
// @Description Create a new SubscriptionPlan
// @Tags sales-subscriptions
// @Accept json
// @Produce json
// @Success 201 {object} SubscriptionPlan
// @Param request body SubscriptionPlan true "Payload"
// @Router /api/sales/subscriptions/subscriptionplan [post]
// @Security BearerAuth
func CreateSubscriptionPlanHandler(c echo.Context) error {
	var data SubscriptionPlan
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSubscriptionPlanService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SubscriptionPlan
// @Description Retrieve a list of all SubscriptionPlan
// @Tags sales-subscriptions
// @Produce json
// @Success 200 {object} SubscriptionPlan
// @Router /api/sales/subscriptions/subscriptionplan [get]
// @Security BearerAuth
func GetAllSubscriptionPlanHandler(c echo.Context) error {
	data, err := GetAllSubscriptionPlanService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSubscriptionPlanByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSubscriptionPlanByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SubscriptionPlan
// @Description Update an existing SubscriptionPlan
// @Tags sales-subscriptions
// @Accept json
// @Produce json
// @Param id path int true "SubscriptionPlan ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions/subscriptionplan/{id} [put]
// @Security BearerAuth
func UpdateSubscriptionPlanHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSubscriptionPlanByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSubscriptionPlanService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SubscriptionPlan
// @Description Delete SubscriptionPlan by ID
// @Tags sales-subscriptions
// @Produce json
// @Param id path int true "SubscriptionPlan ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/sales/subscriptions/subscriptionplan/{id} [delete]
// @Security BearerAuth
func DeleteSubscriptionPlanHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSubscriptionPlanService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


