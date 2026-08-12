package purchase

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreatePurchaseOrder godoc
// @Summary Create a new PurchaseOrder
// @Description Create a new PurchaseOrder in the system
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/supply_chain/purchase [post]
// @Security BearerAuth
func CreatePurchaseOrderHandler(c echo.Context) error {
	var data PurchaseOrder
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreatePurchaseOrderService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllPurchaseOrder godoc
// @Summary Get all PurchaseOrder
// @Description Retrieve a list of all PurchaseOrder
// @Tags supply_chain-purchase
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase [get]
// @Security BearerAuth
func GetAllPurchaseOrderHandler(c echo.Context) error {
	data, err := GetAllPurchaseOrderService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetPurchaseOrderByID godoc
// @Summary Get a PurchaseOrder by ID
// @Description Retrieve a specific PurchaseOrder by its ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/{id} [get]
// @Security BearerAuth
func GetPurchaseOrderByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdatePurchaseOrder godoc
// @Summary Update a PurchaseOrder
// @Description Update an existing PurchaseOrder
// @Tags supply_chain-purchase
// @Accept json
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/{id} [put]
// @Security BearerAuth
func UpdatePurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPurchaseOrderByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdatePurchaseOrderService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeletePurchaseOrder godoc
// @Summary Delete a PurchaseOrder
// @Description Delete a PurchaseOrder by ID
// @Tags supply_chain-purchase
// @Produce json
// @Param id path int true "PurchaseOrder ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/supply_chain/purchase/{id} [delete]
// @Security BearerAuth
func DeletePurchaseOrderHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePurchaseOrderService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
