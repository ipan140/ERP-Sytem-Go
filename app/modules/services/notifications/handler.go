package notifications

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetNotificationLogsHandler(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	logs, err := GetNotificationLogsService(limit)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil log notifikasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Log notifikasi berhasil diambil", logs)
}

func SendNotificationHandler(c echo.Context) error {
	var req struct {
		Channel       string `json:"channel"`
		Recipient     string `json:"recipient"`
		RecipientName string `json:"recipient_name"`
		EntityType    string `json:"entity_type"`
		EntityID      *uint  `json:"entity_id"`
		Subject       string `json:"subject"`
		Message       string `json:"message"`
	}
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Payload request tidak valid", err.Error())
	}
	if req.Recipient == "" || req.Message == "" {
		return utils.SendError(c, http.StatusBadRequest, "Recipient dan Message wajib diisi", "")
	}
	if req.Channel == "" {
		req.Channel = "whatsapp"
	}

	res, err := DispatchNotificationService(req.Channel, req.Recipient, req.RecipientName, req.EntityType, req.EntityID, req.Subject, req.Message)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengirim notifikasi", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Notifikasi berhasil dikirim", res)
}
