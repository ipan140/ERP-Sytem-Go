package sign

import (
	"ERP-System/common/utils"
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSignatureRequest godoc
// @Summary Create a new SignatureRequest
// @Description Create a new SignatureRequest in the system
// @Tags finance-sign
// @Accept json
// @Produce json
// @Success 201 {object} SignatureRequest
// @Param request body SignatureRequest true "Payload"
// @Router /api/finance/sign [post]
// @Security BearerAuth
func CreateSignatureRequestHandler(c echo.Context) error {
	var data SignatureRequest
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSignatureRequestService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSignatureRequest godoc
// @Summary Get all SignatureRequest
// @Description Retrieve a list of all SignatureRequest
// @Tags finance-sign
// @Produce json
// @Success 200 {object} []SignatureRequest
// @Router /api/finance/sign [get]
// @Security BearerAuth
func GetAllSignatureRequestHandler(c echo.Context) error {
	data, err := GetAllSignatureRequestService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSignatureRequestByID godoc
// @Summary Get a SignatureRequest by ID
// @Description Retrieve a specific SignatureRequest by its ID
// @Tags finance-sign
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} SignatureRequest
// @Router /api/finance/sign/{id} [get]
// @Security BearerAuth
func GetSignatureRequestByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSignatureRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSignatureRequest godoc
// @Summary Update a SignatureRequest
// @Description Update an existing SignatureRequest
// @Tags finance-sign
// @Accept json
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign/{id} [put]
// @Security BearerAuth
func UpdateSignatureRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSignatureRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSignatureRequestService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSignatureRequest godoc
// @Summary Delete a SignatureRequest
// @Description Delete a SignatureRequest by ID
// @Tags finance-sign
// @Produce json
// @Param id path int true "SignatureRequest ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/sign/{id} [delete]
// @Security BearerAuth
func DeleteSignatureRequestHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSignatureRequestService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// PreviewDocumentHandler godoc
// @Summary Preview dokumen laporan keuangan dengan tanda tangan digital resmi
// @Description Merender HTML template resmi dokumen finansial lengkap dengan segel dan SHA-256 digital signature
// @Tags finance-sign
// @Produce html
// @Param id path int true "Signature Request ID"
// @Success 200 {string} string "HTML Document Page"
// @Router /api/finance/sign/{id}/preview [get]
func PreviewDocumentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSignatureRequestByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Dokumen tidak ditemukan", err.Error())
	}

	tmplPath := "app/templates/financial_report_template.html"
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		// Fallback jika file template diakses dari root path yang berbeda
		tmpl, err = template.ParseGlob("**/financial_report_template.html")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Template dokumen tidak ditemukan: "+err.Error())
		}
	}

	type AccountItem struct {
		AccountName string
		Debit       string
		Credit      string
		Balance     string
	}

	docData := struct {
		ReportName    string
		Period        string
		Accounts      []AccountItem
		NetTotal      string
		SignerName    string
		SignerRole    string
		SignatureHash string
		IsSigned      bool
		SignDate      string
	}{
		ReportName: data.DocumentTitle,
		Period:     "Tahun Fiskal Berjalan 2026",
		Accounts: []AccountItem{
			{AccountName: "1-1001 Kas & Setara Kas (Bank BCA Giro)", Debit: "1.450.000.000", Credit: "0", Balance: "1.450.000.000"},
			{AccountName: "1-1002 Piutang Usaha (Trade Receivables)", Debit: "320.000.000", Credit: "0", Balance: "320.000.000"},
			{AccountName: "2-1001 Hutang Usaha Pihak Ketiga", Debit: "0", Credit: "210.000.000", Balance: "210.000.000"},
			{AccountName: "4-1001 Pendapatan Usaha Penjualan", Debit: "0", Credit: "2.350.000.000", Balance: "2.350.000.000"},
			{AccountName: "5-1001 Harga Pokok Penjualan (HPP)", Debit: "1.090.000.000", Credit: "0", Balance: "1.090.000.000"},
		},
		NetTotal:      "521.500.000",
		SignerName:    data.SignerName,
		SignerRole:    data.SignerRole,
		SignatureHash: data.SignatureHash,
		IsSigned:      data.Status == "signed",
		SignDate:      data.CreatedAt.Format("02 Jan 2006 15:04"),
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, docData); err != nil {
		return c.String(http.StatusInternalServerError, "Gagal merender template: "+err.Error())
	}

	return c.HTML(http.StatusOK, buf.String())
}


