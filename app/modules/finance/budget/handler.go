package budget

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllBudgetsHandler godoc
// @Summary Ambil semua data anggaran per departemen
// @Description Mengambil ringkasan pagu anggaran, realisasi pengeluaran, sisa kuota, dan status (Aman/Peringatan/Kritis)
// @Tags finance-budget
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/budget [get]
// @Security BearerAuth
func GetAllBudgetsHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		list, totalLimit, totalSpent, remaining, overallUsage, err := GetAllBudgetsService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data anggaran", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Budgets retrieved", map[string]interface{}{
			"items":             list,
			"total_allocated":   totalLimit,
			"total_realized":    totalSpent,
			"total_remaining":   remaining,
			"overall_usage_pct": overallUsage,
		})
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)

	list, total, totalLimit, totalSpent, remaining, overallUsage, err := GetPaginatedBudgetsService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data anggaran", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Budgets retrieved", map[string]interface{}{
		"items":             list,
		"total_allocated":   totalLimit,
		"total_realized":    totalSpent,
		"total_remaining":   remaining,
		"overall_usage_pct": overallUsage,
	}, meta)
}

// CreateBudgetHandler godoc
// @Summary Tetapkan pagu anggaran baru untuk departemen
// @Description Menetapkan limit anggaran belanja per divisi untuk periode kuartal tertentu
// @Tags finance-budget
// @Accept json
// @Produce json
// @Param request body DepartmentBudget true "Payload Anggaran"
// @Success 201 {object} DepartmentBudget
// @Router /api/finance/budget [post]
// @Security BearerAuth
func CreateBudgetHandler(c echo.Context) error {
	var item DepartmentBudget
	if err := c.Bind(&item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateBudgetService(&item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menambahkan anggaran", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Anggaran departemen berhasil didaftarkan", item)
}

// UpdateBudgetHandler godoc
// @Summary Perbarui alokasi atau realisasi anggaran
// @Description Mengubah batas pagu atau mengupdate realisasi anggaran per departemen
// @Tags finance-budget
// @Accept json
// @Produce json
// @Param id path int true "ID Anggaran"
// @Param request body DepartmentBudget true "Payload Update"
// @Success 200 {object} DepartmentBudget
// @Router /api/finance/budget/{id} [put]
// @Security BearerAuth
func UpdateBudgetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	item, err := GetBudgetByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data anggaran tidak ditemukan", err.Error())
	}
	if err := c.Bind(item); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateBudgetService(item); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui anggaran", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Anggaran departemen berhasil diperbarui", item)
}

// DeleteBudgetHandler godoc
// @Summary Hapus data anggaran departemen
// @Description Menghapus pencatatan pagu anggaran departemen
// @Tags finance-budget
// @Produce json
// @Param id path int true "ID Anggaran"
// @Success 200 {string} string "Anggaran berhasil dihapus"
// @Router /api/finance/budget/{id} [delete]
// @Security BearerAuth
func DeleteBudgetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteBudgetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus anggaran", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Anggaran berhasil dihapus", nil)
}
