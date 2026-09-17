package mailer

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateEmailLog godoc
// @Summary Create a new EmailLog
// @Description Create a new EmailLog in the system
// @Tags core-mailer
// @Accept json
// @Produce json
// @Success 201 {object} EmailLog
// @Param request body EmailLog true "Payload"
// @Router /api/core/mailer [post]
// @Security BearerAuth
func CreateEmailLogHandler(c echo.Context) error {
	var data EmailLog
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateEmailLogService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllEmailLog godoc
// @Summary Get all EmailLog
// @Description Retrieve a list of all EmailLog
// @Tags core-mailer
// @Produce json
// @Success 200 {object} []EmailLog
// @Router /api/core/mailer [get]
// @Security BearerAuth
func GetAllEmailLogHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllEmailLogService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedEmailLogService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetEmailLogByID godoc
// @Summary Get a EmailLog by ID
// @Description Retrieve a specific EmailLog by its ID
// @Tags core-mailer
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} EmailLog
// @Router /api/core/mailer/{id} [get]
// @Security BearerAuth
func GetEmailLogByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmailLogByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateEmailLog godoc
// @Summary Update a EmailLog
// @Description Update an existing EmailLog
// @Tags core-mailer
// @Accept json
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/mailer/{id} [put]
// @Security BearerAuth
func UpdateEmailLogHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmailLogByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateEmailLogService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteEmailLog godoc
// @Summary Delete a EmailLog
// @Description Delete a EmailLog by ID
// @Tags core-mailer
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/mailer/{id} [delete]
// @Security BearerAuth
func DeleteEmailLogHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmailLogService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// GetSmtpConfigHandler godoc
// @Summary Get active SMTP configuration
// @Description Retrieve current active SMTP outgoing configuration
// @Tags core-mailer
// @Produce json
// @Success 200 {object} SmtpConfig
// @Router /api/core/mailer/config [get]
// @Security BearerAuth
func GetSmtpConfigHandler(c echo.Context) error {
	data, err := GetSmtpConfigService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil konfigurasi SMTP", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi SMTP berhasil dimuat", data)
}

// SaveSmtpConfigHandler godoc
// @Summary Save or update SMTP configuration
// @Description Save or update SMTP outgoing configuration
// @Tags core-mailer
// @Accept json
// @Produce json
// @Param request body SmtpConfig true "SMTP Payload"
// @Success 200 {object} SmtpConfig
// @Router /api/core/mailer/config [post]
// @Security BearerAuth
func SaveSmtpConfigHandler(c echo.Context) error {
	var payload SmtpConfig
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format payload tidak valid", err.Error())
	}
	if err := SaveSmtpConfigService(&payload); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan konfigurasi SMTP", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi SMTP berhasil disimpan", payload)
}

// TestSendEmailHandler godoc
// @Summary Test SMTP connection and send test email
// @Description Connect to SMTP server and send a verification test message
// @Tags core-mailer
// @Accept json
// @Produce json
// @Param request body map[string]string true "Recipient Payload"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/mailer/test [post]
// @Security BearerAuth
func TestSendEmailHandler(c echo.Context) error {
	var body struct {
		Recipient string `json:"recipient"`
	}
	_ = c.Bind(&body)

	msg, err := TestSendEmailService(body.Recipient)
	if err != nil {
		return utils.SendError(c, http.StatusBadGateway, "Gagal mengirim email uji coba", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, msg, map[string]interface{}{
		"success": true,
		"message": msg,
	})
}


