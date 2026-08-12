package invoicing

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateInvoice godoc
// @Summary Create a new Invoice
// @Description Create a new Invoice in the system
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/invoicing [post]
// @Security BearerAuth
func CreateInvoiceHandler(c echo.Context) error {
	var data Invoice
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateInvoiceService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllInvoice godoc
// @Summary Get all Invoice
// @Description Retrieve a list of all Invoice
// @Tags finance-invoicing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing [get]
// @Security BearerAuth
func GetAllInvoiceHandler(c echo.Context) error {
	data, err := GetAllInvoiceService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetInvoiceByID godoc
// @Summary Get a Invoice by ID
// @Description Retrieve a specific Invoice by its ID
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/{id} [get]
// @Security BearerAuth
func GetInvoiceByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetInvoiceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateInvoice godoc
// @Summary Update a Invoice
// @Description Update an existing Invoice
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/{id} [put]
// @Security BearerAuth
func UpdateInvoiceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetInvoiceByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateInvoiceService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteInvoice godoc
// @Summary Delete a Invoice
// @Description Delete a Invoice by ID
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/{id} [delete]
// @Security BearerAuth
func DeleteInvoiceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteInvoiceService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
