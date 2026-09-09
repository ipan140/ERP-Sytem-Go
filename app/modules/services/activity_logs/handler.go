package activity_logs

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetActivityLogsHandler(c echo.Context) error {
	entityType := c.QueryParam("entity_type")
	entityIDStr := c.QueryParam("entity_id")
	if entityType == "" || entityIDStr == "" {
		return utils.SendError(c, http.StatusBadRequest, "Parameter entity_type dan entity_id wajib diisi", "")
	}

	entityID, err := strconv.Atoi(entityIDStr)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid entity_id", err.Error())
	}

	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	logs, err := GetActivityLogsService(entityType, uint(entityID), limit)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil log aktivitas", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Log aktivitas berhasil diambil", logs)
}
