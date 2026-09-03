package assets

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllCategoriesHandler godoc
// @Summary Ambil semua kategori aset tetap
// @Description Mengambil daftar kategori aset beserta default masa manfaat dan mapping akun GL
// @Tags finance-assets
// @Produce json
// @Success 200 {object} []AssetCategory
// @Router /api/finance/assets/categories [get]
// @Security BearerAuth
func GetAllCategoriesHandler(c echo.Context) error {
	list, err := GetAllCategoriesService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil kategori", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kategori aset retrieved", list)
}

// CreateCategoryHandler godoc
// @Summary Tambah kategori aset baru
// @Description Membuat kategori aset baru beserta default masa manfaat bulan
// @Tags finance-assets
// @Accept json
// @Produce json
// @Param request body AssetCategory true "Payload Kategori"
// @Success 201 {object} AssetCategory
// @Router /api/finance/assets/categories [post]
// @Security BearerAuth
func CreateCategoryHandler(c echo.Context) error {
	var cat AssetCategory
	if err := c.Bind(&cat); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateCategoryService(&cat); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat kategori", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Kategori berhasil dibuat", cat)
}

// UpdateCategoryHandler godoc
// @Summary Perbarui kategori aset
// @Description Mengubah data kategori aset atau masa manfaat default
// @Tags finance-assets
// @Accept json
// @Produce json
// @Param id path int true "ID Kategori"
// @Param request body AssetCategory true "Payload Kategori"
// @Success 200 {object} AssetCategory
// @Router /api/finance/assets/categories/{id} [put]
// @Security BearerAuth
func UpdateCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := GetCategoryByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Kategori tidak ditemukan", err.Error())
	}
	if err := c.Bind(cat); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateCategoryService(cat); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengupdate kategori", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kategori berhasil diupdate", cat)
}

// DeleteCategoryHandler godoc
// @Summary Hapus kategori aset
// @Description Menghapus kategori aset berdasarkan ID
// @Tags finance-assets
// @Produce json
// @Param id path int true "ID Kategori"
// @Success 200 {string} string "Kategori berhasil dihapus"
// @Router /api/finance/assets/categories/{id} [delete]
// @Security BearerAuth
func DeleteCategoryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCategoryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus kategori", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Kategori berhasil dihapus", nil)
}

// GetAllAssetsHandler godoc
// @Summary Ambil semua inventaris aset tetap
// @Description Mengambil daftar inventaris aset, akumulasi penyusutan, dan nilai buku bersih
// @Tags finance-assets
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/assets [get]
// @Security BearerAuth
func GetAllAssetsHandler(c echo.Context) error {
	list, totalAcq, totalAccum, totalNBV, err := GetAllAssetsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data aset", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Assets retrieved", map[string]interface{}{
		"items":              list,
		"total_acquisition":  totalAcq,
		"total_accumulated":  totalAccum,
		"total_net_book_val": totalNBV,
	})
}

// CreateAssetHandler godoc
// @Summary Daftarkan aset tetap baru
// @Description Mencatat aktiva tetap baru, menghitung penyusutan per bulan otomatis
// @Tags finance-assets
// @Accept json
// @Produce json
// @Param request body FixedAsset true "Payload Aset"
// @Success 201 {object} FixedAsset
// @Router /api/finance/assets [post]
// @Security BearerAuth
func CreateAssetHandler(c echo.Context) error {
	var asset FixedAsset
	if err := c.Bind(&asset); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAssetService(&asset); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mendaftarkan aset", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Aset berhasil didaftarkan", asset)
}

// UpdateAssetHandler godoc
// @Summary Perbarui data aset tetap
// @Description Mengubah rincian nama, harga perolehan, atau masa manfaat aset
// @Tags finance-assets
// @Accept json
// @Produce json
// @Param id path int true "ID Aset"
// @Param request body FixedAsset true "Payload Update"
// @Success 200 {object} FixedAsset
// @Router /api/finance/assets/{id} [put]
// @Security BearerAuth
func UpdateAssetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	asset, err := GetAssetByIDRepo(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Aset tidak ditemukan", err.Error())
	}
	if err := c.Bind(asset); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateAssetService(asset); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui aset", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Aset berhasil diperbarui", asset)
}

// DeleteAssetHandler godoc
// @Summary Hapus aset tetap
// @Description Menghapus pencatatan aset dari database
// @Tags finance-assets
// @Produce json
// @Param id path int true "ID Aset"
// @Success 200 {string} string "Aset berhasil dihapus"
// @Router /api/finance/assets/{id} [delete]
// @Security BearerAuth
func DeleteAssetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAssetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus aset", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Aset berhasil dihapus", nil)
}

// PostMonthlyDepreciationHandler godoc
// @Summary Posting jurnal depresiasi massal bulanan
// @Description Mengalkulasi dan memposting jurnal penyusutan seluruh aktiva aktif ke General Ledger
// @Tags finance-assets
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/assets/depreciate [post]
// @Security BearerAuth
func PostMonthlyDepreciationHandler(c echo.Context) error {
	totalDepreciated, err := ExecuteMonthlyDepreciationService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengeksekusi depresiasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Jurnal depresiasi bulanan berhasil diposting ke GL!", map[string]interface{}{
		"total_posted":      totalDepreciated,
		"gl_account_debit":  "6-1008 Beban Depresiasi Aset Tetap",
		"gl_account_credit": "1-1009 Akumulasi Penyusutan",
	})
}
