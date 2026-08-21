package report

import (
	"net/http"

	"ERP-System/common/utils"
	"ERP-System/pkg/excelgen"
	"ERP-System/pkg/pdfgen"

	"github.com/labstack/echo/v4"
)

// GenerateDynamicExcelHandler godoc
// @Summary Generate dynamic Excel file
// @Description Accepts dynamic headers and rows, returns an Excel file (.xlsx)
// @Tags Core - Report
// @Accept json
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Param body body DynamicExcelRequest true "Excel Data Payload"
// @Success 200 {file} file
// @Security BearerAuth
// @Router /core/report/excel [post]
func GenerateDynamicExcelHandler(c echo.Context) error {
	var req DynamicExcelRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format Data Salah", err.Error())
	}

	if req.FileName == "" {
		req.FileName = "Laporan_Dinamis.xlsx"
	}
	if req.SheetName == "" {
		req.SheetName = "Sheet1"
	}

	excelBytes, err := excelgen.GenerateExcel(req.SheetName, req.Headers, req.Data)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal merender Excel", err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+req.FileName)

	_, err = c.Response().Writer.Write(excelBytes)
	return err
}

// GenerateDynamicPDFHandler godoc
// @Summary Generate dynamic PDF file
// @Description Accepts raw HTML content and renders it into a PDF file
// @Tags Core - Report
// @Accept json
// @Produce application/pdf
// @Param body body DynamicPDFRequest true "PDF HTML Payload"
// @Success 200 {file} file
// @Security BearerAuth
// @Router /core/report/pdf [post]
func GenerateDynamicPDFHandler(c echo.Context) error {
	var req DynamicPDFRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format Data Salah", err.Error())
	}

	if req.FileName == "" {
		req.FileName = "Laporan_Dinamis.pdf"
	}

	pdfBytes, err := pdfgen.GeneratePDF(req.HTMLContent, req.PageSize, req.Orientation)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal merender PDF", err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/pdf")
	c.Response().Header().Set("Content-Disposition", "attachment; filename="+req.FileName)

	_, err = c.Response().Writer.Write(pdfBytes)
	return err
}

