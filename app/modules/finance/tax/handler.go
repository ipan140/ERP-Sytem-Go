package tax

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// GetAllTaxMasterConfigsHandler godoc
// @Summary Ambil semua master konfigurasi regulasi pajak
// @Description Mengambil daftar master tarif pajak (PPN, PPh 23, PPh 21 TER, dll)
// @Tags finance-tax
// @Produce json
// @Success 200 {object} []TaxMasterConfig
// @Router /api/finance/tax/configs [get]
// @Security BearerAuth
func GetAllTaxMasterConfigsHandler(c echo.Context) error {
	list, err := GetAllTaxMasterConfigsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil konfigurasi pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Tax configurations retrieved", list)
}

// CreateTaxMasterConfigHandler godoc
// @Summary Tambah master regulasi pajak baru (Admin)
// @Description Mendaftarkan skema regulasi pajak baru beserta tarif dan dasar hukumnya
// @Tags finance-tax
// @Accept json
// @Produce json
// @Param request body TaxMasterConfig true "Payload Master Pajak"
// @Success 201 {object} TaxMasterConfig
// @Router /api/finance/tax/configs [post]
// @Security BearerAuth
func CreateTaxMasterConfigHandler(c echo.Context) error {
	var item TaxMasterConfig
	if err := c.Bind(&item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateTaxMasterConfigService(&item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menambahkan regulasi pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Regulasi pajak berhasil didaftarkan", item)
}

// UpdateTaxMasterConfigHandler godoc
// @Summary Perbarui master regulasi pajak (Admin)
// @Description Mengubah persentase tarif pajak atau deskripsi dasar hukum
// @Tags finance-tax
// @Accept json
// @Produce json
// @Param id path int true "ID Master Pajak"
// @Param request body TaxMasterConfig true "Payload Update"
// @Success 200 {object} TaxMasterConfig
// @Router /api/finance/tax/configs/{id} [put]
// @Security BearerAuth
func UpdateTaxMasterConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := GetTaxMasterConfigByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data tidak ditemukan", err.Error())
	}
	if err := c.Bind(item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateTaxMasterConfigService(item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui regulasi pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Regulasi pajak berhasil diperbarui", item)
}

// DeleteTaxMasterConfigHandler godoc
// @Summary Hapus skema master pajak (Admin)
// @Description Menghapus skema master tarif pajak
// @Tags finance-tax
// @Produce json
// @Param id path int true "ID Master Pajak"
// @Success 200 {string} string "Regulasi pajak berhasil dihapus"
// @Router /api/finance/tax/configs/{id} [delete]
// @Security BearerAuth
func DeleteTaxMasterConfigHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTaxMasterConfigService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus regulasi pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Regulasi pajak berhasil dihapus", nil)
}

// GetAllTaxReportsHandler godoc
// @Summary Ambil semua rekapitulasi bukti potong pajak
// @Description Mengambil daftar bukti pemotongan PPh/PPN transaksi masa pajak berjalan
// @Tags finance-tax
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/tax [get]
// @Security BearerAuth
func GetAllTaxReportsHandler(c echo.Context) error {
	list, totalTaxAmount, err := GetAllTaxReportsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Tax reports retrieved", map[string]interface{}{
		"items":            list,
		"total_tax_amount": totalTaxAmount,
		"currency":         "IDR",
		"tax_office":       "KPP Pratama Jakarta Kebayoran Baru",
	})
}

// CreateTaxReportHandler godoc
// @Summary Catat bukti pemotongan pajak transaksi baru
// @Description Menyimpan bukti potong PPh 21/23/4(2)/PPN atas rekanan/vendor
// @Tags finance-tax
// @Accept json
// @Produce json
// @Param request body TaxReportSummary true "Payload Bukti Potong"
// @Success 201 {object} TaxReportSummary
// @Router /api/finance/tax [post]
// @Security BearerAuth
func CreateTaxReportHandler(c echo.Context) error {
	var item TaxReportSummary
	if err := c.Bind(&item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateTaxReportService(&item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data pajak berhasil disimpan", item)
}

// UpdateTaxReportHandler godoc
// @Summary Perbarui data bukti potong pajak
// @Description Mengubah rincian nama rekanan, NPWP, atau DPP transaksi
// @Tags finance-tax
// @Accept json
// @Produce json
// @Param id path int true "ID Bukti Potong"
// @Param request body TaxReportSummary true "Payload Update"
// @Success 200 {object} TaxReportSummary
// @Router /api/finance/tax/{id} [put]
// @Security BearerAuth
func UpdateTaxReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := GetTaxReportByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data tidak ditemukan", err.Error())
	}
	if err := c.Bind(item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateTaxReportService(item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data pajak berhasil diperbarui", item)
}

// DeleteTaxReportHandler godoc
// @Summary Hapus bukti potong pajak
// @Description Menghapus pencatatan bukti potong dari daftar SPT Masa
// @Tags finance-tax
// @Produce json
// @Param id path int true "ID Bukti Potong"
// @Success 200 {string} string "Data pajak berhasil dihapus"
// @Router /api/finance/tax/{id} [delete]
// @Security BearerAuth
func DeleteTaxReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteTaxReportService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus data pajak", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data pajak berhasil dihapus", nil)
}

// ExportEBupotCsvHandler godoc
// @Summary Ekspor CSV format e-Bupot Unifikasi DJP
// @Description Mengunduh file CSV siap impor ke aplikasi e-Bupot Unifikasi / Coretax Ditjen Pajak
// @Tags finance-tax
// @Produce text/csv
// @Success 200 {string} string "CSV Content"
// @Router /api/finance/tax/export-ebupot [get]
func ExportEBupotCsvHandler(c echo.Context) error {
	csvContent, err := GenerateEBupotCsvStringService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal men-generate CSV", err.Error())
	}
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=e-Bupot_SPT_Masa_"+time.Now().Format("20060102")+".csv")
	return c.String(http.StatusOK, csvContent)
}
