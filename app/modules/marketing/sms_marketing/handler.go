package sms_marketing

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateSmsCampaign godoc
// @Summary Create a new SmsCampaign
// @Description Create a new SMS/WhatsApp marketing campaign in the system
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param request body SmsCampaign true "Sms Campaign Payload"
// @Success 201 {object} utils.SuccessResponse{data=SmsCampaign} "Campaign created successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload"
// @Failure 500 {object} utils.ErrorResponse "Failed to create data"
// @Router /api/marketing/sms_marketing [post]
// @Security BearerAuth
func CreateSmsCampaignHandler(c echo.Context) error {
	var data SmsCampaign
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSmsCampaignService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSmsCampaign godoc
// @Summary Get all SmsCampaign
// @Description Retrieve a list of all SMS/WhatsApp campaigns
// @Tags marketing-sms_marketing
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]SmsCampaign} "List of SMS campaigns"
// @Failure 500 {object} utils.ErrorResponse "Failed to retrieve data"
// @Router /api/marketing/sms_marketing [get]
// @Security BearerAuth
func GetAllSmsCampaignHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllSmsCampaignService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedSmsCampaignService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetSmsCampaignByID godoc
// @Summary Get a SmsCampaign by ID
// @Description Retrieve a specific SMS/WhatsApp campaign by its ID
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} utils.SuccessResponse{data=SmsCampaign} "Campaign found"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 404 {object} utils.ErrorResponse "Data not found"
// @Router /api/marketing/sms_marketing/{id} [get]
// @Security BearerAuth
func GetSmsCampaignByIDHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	data, err := GetSmsCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSmsCampaign godoc
// @Summary Update a SmsCampaign
// @Description Update an existing SMS/WhatsApp campaign
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Param request body SmsCampaign true "Updated SMS Campaign Payload"
// @Success 200 {object} utils.SuccessResponse{data=SmsCampaign} "Campaign updated successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid request payload or ID"
// @Failure 404 {object} utils.ErrorResponse "Data not found"
// @Failure 500 {object} utils.ErrorResponse "Failed to update data"
// @Router /api/marketing/sms_marketing/{id} [put]
// @Security BearerAuth
func UpdateSmsCampaignHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	data, err := GetSmsCampaignByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSmsCampaignService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSmsCampaign godoc
// @Summary Delete a SmsCampaign
// @Description Delete a SMS/WhatsApp campaign by ID
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} utils.SuccessResponse{data=nil} "Data deleted successfully"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 500 {object} utils.ErrorResponse "Failed to delete data"
// @Router /api/marketing/sms_marketing/{id} [delete]
// @Security BearerAuth
func DeleteSmsCampaignHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	if err := DeleteSmsCampaignService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// ExecuteBroadcastHandler godoc
// @Summary Execute real broadcast to audience
// @Description Broadcast SMS / WhatsApp messages to targeted customer audience
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Success 200 {object} utils.SuccessResponse{data=BroadcastResult} "Broadcast execution summary"
// @Failure 400 {object} utils.ErrorResponse "Invalid ID parameter"
// @Failure 500 {object} utils.ErrorResponse "Failed to execute broadcast"
// @Router /api/marketing/sms_marketing/{id}/broadcast [post]
// @Security BearerAuth
func ExecuteBroadcastHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	result, err := ExecuteBroadcastService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to execute broadcast", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Broadcast berhasil dieksekusi ke pelanggan", result)
}

// TestSendPayload payload for sending test message
type TestSendPayload struct {
	TargetPhone string `json:"target_phone" example:"081234567890"`
	Channel     string `json:"channel" example:"WhatsApp"`
}

// SendTestMessageHandler godoc
// @Summary Send test broadcast message to user's phone / WhatsApp
// @Description Dispatches a test SMS/WhatsApp message and generates a direct WhatsApp link
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param id path int true "SmsCampaign ID"
// @Param request body TestSendPayload true "Test Dispatch Target"
// @Success 200 {object} utils.SuccessResponse{data=map[string]interface{}} "Test message dispatch response"
// @Failure 400 {object} utils.ErrorResponse "Invalid payload or missing phone number"
// @Failure 500 {object} utils.ErrorResponse "Failed to send test message"
// @Router /api/marketing/sms_marketing/{id}/send-test [post]
// @Security BearerAuth
func SendTestMessageHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "ID tidak valid", "")
	}
	var req TestSendPayload
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if req.TargetPhone == "" {
		return utils.SendError(c, http.StatusBadRequest, "Nomor HP / WhatsApp wajib diisi", "")
	}

	result, err := SendTestMessageService(uint(id), req.TargetPhone, req.Channel)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to send test message", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Pesan tes berhasil dikirim", result)
}

// GetAllWaTemplatesHandler godoc
// @Summary Get all WhatsApp WABA Templates
// @Description Retrieve list of all approved Meta Cloud API message templates
// @Tags marketing-sms_marketing
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=[]WaTemplate} "List of WhatsApp Templates"
// @Failure 500 {object} utils.ErrorResponse "Gagal mengambil template"
// @Router /api/marketing/sms_marketing/templates [get]
// @Security BearerAuth
func GetAllWaTemplatesHandler(c echo.Context) error {
	list, err := GetAllWaTemplatesService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil template WhatsApp", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data template WhatsApp berhasil diambil", list)
}

// CreateWaTemplateHandler godoc
// @Summary Create / Register a WhatsApp Template
// @Description Submit a new WhatsApp message template to Meta Cloud API
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param request body WaTemplate true "WaTemplate Payload"
// @Success 201 {object} utils.SuccessResponse{data=WaTemplate} "Template WhatsApp berhasil diajukan"
// @Failure 400 {object} utils.ErrorResponse "Invalid payload"
// @Failure 500 {object} utils.ErrorResponse "Gagal membuat template"
// @Router /api/marketing/sms_marketing/templates [post]
// @Security BearerAuth
func CreateWaTemplateHandler(c echo.Context) error {
	var data WaTemplate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateWaTemplateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat template WhatsApp", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Template WhatsApp berhasil diajukan ke Meta", data)
}

// UpdateWaTemplateHandler godoc
// @Summary Update a WhatsApp Template
// @Description Update existing WhatsApp message template
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param id path int true "Template ID"
// @Param request body WaTemplate true "WaTemplate Payload"
// @Success 200 {object} utils.SuccessResponse{data=WaTemplate} "Template WhatsApp berhasil diperbarui"
// @Failure 404 {object} utils.ErrorResponse "Template tidak ditemukan"
// @Router /api/marketing/sms_marketing/templates/{id} [put]
// @Security BearerAuth
func UpdateWaTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWaTemplateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Template tidak ditemukan", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateWaTemplateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui template", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Template WhatsApp berhasil diperbarui", data)
}

// DeleteWaTemplateHandler godoc
// @Summary Delete a WhatsApp Template
// @Description Delete WhatsApp message template by ID
// @Tags marketing-sms_marketing
// @Produce json
// @Param id path int true "Template ID"
// @Success 200 {object} utils.SuccessResponse "Template WhatsApp berhasil dihapus"
// @Failure 500 {object} utils.ErrorResponse "Gagal menghapus template"
// @Router /api/marketing/sms_marketing/templates/{id} [delete]
// @Security BearerAuth
func DeleteWaTemplateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWaTemplateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menghapus template", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Template WhatsApp berhasil dihapus", nil)
}

// GetWaConfigHandler godoc
// @Summary Get WhatsApp WABA Configuration
// @Description Retrieve current WhatsApp Business Account (WABA) credentials
// @Tags marketing-sms_marketing
// @Produce json
// @Success 200 {object} utils.SuccessResponse{data=WaConfig} "Konfigurasi WhatsApp WABA berhasil dimuat"
// @Router /api/marketing/sms_marketing/config [get]
// @Security BearerAuth
func GetWaConfigHandler(c echo.Context) error {
	cfg, err := GetWaConfigService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengambil konfigurasi WhatsApp", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi WhatsApp WABA berhasil dimuat", cfg)
}

// SaveWaConfigHandler godoc
// @Summary Save WhatsApp WABA Configuration
// @Description Store or update Meta Cloud API WhatsApp credentials
// @Tags marketing-sms_marketing
// @Accept json
// @Produce json
// @Param request body WaConfig true "WaConfig Payload"
// @Success 200 {object} utils.SuccessResponse{data=WaConfig} "Konfigurasi WhatsApp Cloud API berhasil disimpan"
// @Router /api/marketing/sms_marketing/config [post]
// @Security BearerAuth
func SaveWaConfigHandler(c echo.Context) error {
	var cfg WaConfig
	if err := c.Bind(&cfg); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := SaveWaConfigService(&cfg); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menyimpan konfigurasi WhatsApp", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Konfigurasi WhatsApp Cloud API berhasil disimpan", cfg)
}


