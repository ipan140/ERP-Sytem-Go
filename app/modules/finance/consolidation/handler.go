package consolidation

import (
	"ERP-System/common/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GenerateConsolidationHandler godoc
// @Summary Generate Multi-Company Consolidation
// @Description Aggregate ledgers from all companies into one report
// @Tags finance-consolidation
// @Produce json
// @Param name query string true "Report Name"
// @Param period query string true "Period"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/consolidation/generate [post]
// @Security BearerAuth
func GenerateConsolidationHandler(c echo.Context) error {
	name := c.QueryParam("name")
	period := c.QueryParam("period")

	if err := GenerateConsolidationService(name, period); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat konsolidasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Laporan Konsolidasi berhasil dicetak", nil)
}

func GetAllConsolidationReportsHandler(c echo.Context) error {
	data, err := GetAllConsolidationReportsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Berhasil mengambil data", data)
}

func GetConsolidationReportByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetConsolidationReportByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data tidak ditemukan", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Berhasil mengambil data", data)
}

func CreateConsolidationReportHandler(c echo.Context) error {
	type Req struct {
		Name     string   `json:"name"`
		Period   string   `json:"period"`
		Branches []string `json:"branches"`
	}
	var req Req
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	branchesJson, _ := json.Marshal(req.Branches)
	payload := ConsolidationReport{
		Name:     req.Name,
		Period:   req.Period,
		Branches: string(branchesJson),
	}

	if err := CreateConsolidationReportService(&payload); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat laporan konsolidasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Laporan konsolidasi berhasil dibuat", payload)
}

func UpdateConsolidationReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetConsolidationReportByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data tidak ditemukan", err.Error())
	}

	type Req struct {
		Name     string   `json:"name"`
		Period   string   `json:"period"`
		Branches []string `json:"branches"`
	}
	var req Req
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	branchesJson, _ := json.Marshal(req.Branches)
	payload := ConsolidationReport{
		ID:       uint(id),
		Name:     req.Name,
		Period:   req.Period,
		Branches: string(branchesJson),
	}

	if err := UpdateConsolidationReportService(&payload); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengupdate laporan", err.Error())
	}
	data.Name = payload.Name
	data.Period = payload.Period
	data.Branches = payload.Branches
	return utils.SendSuccess(c, http.StatusOK, "Laporan berhasil diperbarui", data)
}

func DeleteConsolidationReportHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteConsolidationReportService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data berhasil dihapus", nil)
}


