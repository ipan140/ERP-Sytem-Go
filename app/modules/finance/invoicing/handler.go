package invoicing

import (
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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

// PostInvoiceHandler godoc
// @Summary Confirm and Post Invoice
// @Description Post an invoice and trigger auto-journal entry in accounting
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/{id}/post [post]
// @Security BearerAuth
func PostInvoiceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := PostInvoiceService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal melakukan posting faktur", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Faktur berhasil diposting dan Jurnal telah dibuat", nil)
}

// RefundInvoiceHandler godoc
// @Summary Create a Refund / Credit Note
// @Description Refund a posted invoice and create a reversal journal
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/{id}/refund [post]
// @Security BearerAuth
func RefundInvoiceHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := RefundInvoiceService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal melakukan refund faktur", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Refund berhasil, jurnal pembalik telah dibuat", nil)
}

// TriggerDunningHandler godoc
// @Summary Trigger Dunning Process
// @Description Scan overdue invoices and increment their threat level
// @Tags finance-invoicing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/dunning [post]
// @Security BearerAuth
func TriggerDunningHandler(c echo.Context) error {
	if err := RunDunningProcess(); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menjalankan mesin dunning", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Mesin Dunning berhasil menyapu seluruh faktur tunggakan", nil)
}

// @Summary Create PaymentTermLine
// @Description Create a new PaymentTermLine
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/invoicing/paymenttermline [post]
// @Security BearerAuth
func CreatePaymentTermLineHandler(c echo.Context) error { var data PaymentTermLine; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreatePaymentTermLineService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
// @Summary Get all PaymentTermLine
// @Description Retrieve a list of all PaymentTermLine
// @Tags finance-invoicing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/paymenttermline [get]
// @Security BearerAuth
func GetAllPaymentTermLineHandler(c echo.Context) error { data, err := GetAllPaymentTermLineService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetPaymentTermLineByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentTermLineByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Update PaymentTermLine
// @Description Update an existing PaymentTermLine
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Param id path int true "PaymentTermLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/paymenttermline/{id} [put]
// @Security BearerAuth
func UpdatePaymentTermLineHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetPaymentTermLineByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdatePaymentTermLineService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Delete PaymentTermLine
// @Description Delete PaymentTermLine by ID
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "PaymentTermLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/paymenttermline/{id} [delete]
// @Security BearerAuth
func DeletePaymentTermLineHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeletePaymentTermLineService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }

// @Summary Create TaxRepartitionLine
// @Description Create a new TaxRepartitionLine
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/invoicing/taxrepartitionline [post]
// @Security BearerAuth
func CreateTaxRepartitionLineHandler(c echo.Context) error { var data TaxRepartitionLine; if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }; if err := CreateTaxRepartitionLineService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusCreated, "Success", data) }
// @Summary Get all TaxRepartitionLine
// @Description Retrieve a list of all TaxRepartitionLine
// @Tags finance-invoicing
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/taxrepartitionline [get]
// @Security BearerAuth
func GetAllTaxRepartitionLineHandler(c echo.Context) error { data, err := GetAllTaxRepartitionLineService(); if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
func GetTaxRepartitionLineByIDHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetTaxRepartitionLineByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Update TaxRepartitionLine
// @Description Update an existing TaxRepartitionLine
// @Tags finance-invoicing
// @Accept json
// @Produce json
// @Param id path int true "TaxRepartitionLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/taxrepartitionline/{id} [put]
// @Security BearerAuth
func UpdateTaxRepartitionLineHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); data, err := GetTaxRepartitionLineByIDService(uint(id)); if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }; if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }; if err := UpdateTaxRepartitionLineService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", data) }
// @Summary Delete TaxRepartitionLine
// @Description Delete TaxRepartitionLine by ID
// @Tags finance-invoicing
// @Produce json
// @Param id path int true "TaxRepartitionLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/invoicing/taxrepartitionline/{id} [delete]
// @Security BearerAuth
func DeleteTaxRepartitionLineHandler(c echo.Context) error { id, _ := strconv.Atoi(c.Param("id")); if err := DeleteTaxRepartitionLineService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }; return utils.SendSuccess(c, http.StatusOK, "Success", nil) }
