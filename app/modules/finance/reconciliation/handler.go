package reconciliation

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllBankStatementsHandler godoc
// @Summary Ambil semua mutasi rekening koran bank
// @Description Mengambil daftar seluruh mutasi bank (BCA, Mandiri, dll) beserta status rekonsiliasi
// @Tags finance-reconciliation
// @Produce json
// @Success 200 {object} []BankStatementItem
// @Router /api/finance/reconciliation [get]
// @Security BearerAuth
func GetAllBankStatementsHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		list, err := GetAllBankStatementsService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Bank statement items retrieved", list)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	bank := c.QueryParam("bank")
	status := c.QueryParam("status")

	list, total, err := GetPaginatedBankStatementsService(offset, limit, search, bank, status)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Bank statement items retrieved", list, meta)
}

// CreateBankStatementHandler godoc
// @Summary Catat / input mutasi rekening koran baru
// @Description Menambahkan data mutasi kas bank masuk (kredit) atau keluar (debit)
// @Tags finance-reconciliation
// @Accept json
// @Produce json
// @Param request body BankStatementItem true "Payload Mutasi"
// @Success 201 {object} BankStatementItem
// @Router /api/finance/reconciliation [post]
// @Security BearerAuth
func CreateBankStatementHandler(c echo.Context) error {
	var item BankStatementItem
	if err := c.Bind(&item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateBankStatementService(&item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menambahkan mutasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Mutasi rekening koran berhasil ditambahkan", item)
}

// UpdateBankStatementHandler godoc
// @Summary Perbarui data mutasi bank
// @Description Mengedit informasi deskripsi, nominal, atau no referensi mutasi bank
// @Tags finance-reconciliation
// @Accept json
// @Produce json
// @Param id path int true "ID Mutasi Bank"
// @Param request body BankStatementItem true "Payload Update"
// @Success 200 {object} BankStatementItem
// @Router /api/finance/reconciliation/{id} [put]
// @Security BearerAuth
func UpdateBankStatementHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := GetBankStatementByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data tidak ditemukan", err.Error())
	}
	if err := c.Bind(item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateBankStatementService(item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data mutasi berhasil diperbarui", item)
}

// DeleteBankStatementHandler godoc
// @Summary Hapus data mutasi bank
// @Description Menghapus baris mutasi bank berdasarkan ID
// @Tags finance-reconciliation
// @Produce json
// @Param id path int true "ID Mutasi Bank"
// @Success 200 {string} string "Mutasi berhasil dihapus"
// @Router /api/finance/reconciliation/{id} [delete]
// @Security BearerAuth
func DeleteBankStatementHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBankStatementService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus mutasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Mutasi berhasil dihapus", nil)
}

// AutoReconcileHandler godoc
// @Summary Jalankan pencocokan otomatis (Auto-Match)
// @Description Menjalankan algoritma auto-match untuk mencocokkan mutasi kas masuk dengan invoice terbuka
// @Tags finance-reconciliation
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/reconciliation/auto [post]
// @Security BearerAuth
func AutoReconcileHandler(c echo.Context) error {
	matchedCount, err := AutoReconcileService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal auto-reconcile", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Auto-reconcile selesai diproses", map[string]interface{}{
		"matched_items": matchedCount,
	})
}

// ManualReconcileHandler godoc
// @Summary Rekonsiliasi manual mutasi bank dengan invoice
// @Description Menghubungkan mutasi bank dengan invoice atau pos akun biaya tertentu
// @Tags finance-reconciliation
// @Accept json
// @Produce json
// @Param id path int true "ID Mutasi Bank"
// @Param request body map[string]string true "Nomor Invoice Tagihan"
// @Success 200 {object} BankStatementItem
// @Router /api/finance/reconciliation/{id}/manual [put]
// @Security BearerAuth
func ManualReconcileHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	type Req struct {
		MatchedInvoice string `json:"matched_invoice"`
	}
	var req Req
	c.Bind(&req)

	item, err := ManualReconcileService(uint(id), req.MatchedInvoice)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mencocokkan mutasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Item berhasil direkonsiliasi", item)
}

// UploadBankStatementCsvHandler godoc
// @Summary Unggah file rekening koran bank (.CSV)
// @Description Mengunggah dan mem-parsing file CSV rekening koran (format KlikBCA / Mandiri MCM)
// @Tags finance-reconciliation
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File CSV Rekening Koran"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/reconciliation/upload-csv [post]
// @Security BearerAuth
func UploadBankStatementCsvHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "File CSV wajib diunggah", err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuka file", err.Error())
	}
	defer src.Close()

	count, err := ImportBankStatementCsvService(src)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Gagal memproses format file CSV: "+err.Error(), err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "File rekening koran berhasil diimpor", map[string]interface{}{
		"imported_rows": count,
		"filename":      file.Filename,
	})
}
