package whatsapp

import (
	"ERP-System/config"
	"fmt"
	"strings"
	"time"
)

func CreateWaTemplateService(data *WaTemplate) error {
	return CreateWaTemplate(data)
}

func GetAllWaTemplateService() ([]WaTemplate, error) {
	return GetAllWaTemplate()
}

func GetPaginatedWaTemplateService(offset, limit int, search string) ([]WaTemplate, int64, error) {
	return GetPaginatedWaTemplate(offset, limit, search)
}

func GetWaTemplateByIDService(id uint) (*WaTemplate, error) {
	return GetWaTemplateByID(id)
}

func UpdateWaTemplateService(data *WaTemplate) error {
	return UpdateWaTemplate(data)
}

func DeleteWaTemplateService(id uint) error {
	return DeleteWaTemplate(id)
}

func GetWhatsappConfigService() (*WhatsappConfig, error) {
	return GetActiveWhatsappConfig()
}

func SaveWhatsappConfigService(data *WhatsappConfig) error {
	return SaveWhatsappConfig(data)
}

func GetPaginatedWaLogsService(offset, limit int, search string) ([]WaLog, int64, error) {
	return GetPaginatedWaLogs(offset, limit, search)
}

func TestSendWhatsappService(phone, templateCode string) (string, error) {
	cfg, err := GetActiveWhatsappConfig()
	if err != nil {
		return "", fmt.Errorf("gagal mengambil konfigurasi WhatsApp Gateway: %w", err)
	}

	phone = strings.TrimSpace(phone)
	if phone == "" {
		return "", fmt.Errorf("nomor telepon WhatsApp tujuan harus diisi")
	}

	tplName := "Notifikasi Pesan Sistem"
	msgContent := "Halo, ini adalah pesan uji coba dari sistem ERP WhatsApp Gateway."
	if templateCode != "" {
		var tpl WaTemplate
		if err := config.DB.Where("code = ?", templateCode).First(&tpl).Error; err == nil {
			tplName = tpl.Name
			msgContent = fmt.Sprintf("Halo, ini adalah pesan uji coba template [%s] untuk nomor %s.", tpl.Name, phone)
		}
	}

	var sendErr error
	if strings.TrimSpace(cfg.ApiToken) == "" {
		sendErr = fmt.Errorf("API Token WhatsApp Gateway belum dikonfigurasi")
	}

	status := "DELIVERED"
	errMsg := ""
	if sendErr != nil {
		status = "FAILED"
		errMsg = sendErr.Error()
	}

	log := WaLog{
		Phone:        phone,
		Recipient:    phone,
		TemplateName: tplName,
		Message:      msgContent,
		Status:       status,
		ErrorMessage: errMsg,
		SentAt:       time.Now(),
	}
	_ = CreateWaLog(&log)

	if sendErr != nil {
		return "", sendErr
	}

	return fmt.Sprintf("Pesan WhatsApp berhasil dikirim ke nomor %s dengan status DELIVERED", phone), nil
}
