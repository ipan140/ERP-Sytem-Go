package consolidation

import (
	"ERP-System/common/utils"
	"net/http"

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
