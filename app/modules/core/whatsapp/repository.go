package whatsapp

import (
	"ERP-System/config"
	"time"
	"gorm.io/gorm/clause"
)

func CreateWaTemplate(data *WaTemplate) error {
	return config.DB.Create(data).Error
}

func GetAllWaTemplate() ([]WaTemplate, error) {
	var list []WaTemplate
	err := config.DB.Preload(clause.Associations).Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultWaTemplates()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultWaTemplates() error {
	defaults := []WaTemplate{
		{
			Code:     "TPL_PAYROLL_01",
			Category: "Payroll & HRIS",
			Name:     "Notifikasi Slip Gaji Karyawan",
			Content:  "Halo {{1}}, Slip Gaji Elektronik Anda untuk periode {{2}} telah terbit.\nTotal Gaji Bersih: {{3}}.\nSilakan unduh dokumen terlampir pada portal ERP ESS. Terima kasih.",
			Status:   "APPROVED",
		},
		{
			Code:     "TPL_LEAVE_APPROVAL",
			Category: "HR Attendance",
			Name:     "Persetujuan Pengajuan Cuti",
			Content:  "Yth {{1}}, permohonan {{2}} Anda selama {{3}} hari telah DISETUJUI oleh atasan langsung.\nSisa cuti tahunan Anda: {{4}} hari.",
			Status:   "APPROVED",
		},
		{
			Code:     "TPL_INVOICE_BILLING",
			Category: "Finance & Accounting",
			Name:     "Tagihan Faktur Penjualan",
			Content:  "Yth. Pelanggan {{1}}, Invoice {{2}} senilai {{3}} telah jatuh tempo pada {{4}}.\nSilakan lakukan pembayaran ke rekening Virtual Account BCA / Mandiri terlampir.",
			Status:   "APPROVED",
		},
		{
			Code:     "TPL_AUTH_OTP",
			Category: "Security & Auth",
			Name:     "Kode OTP Verifikasi Login",
			Content:  "KODE KEAMANAN ERP: {{1}} adalah kode rahasia verifikasi login sistem Anda. Jangan berikan kepada siapapun termasuk staf IT. Berlaku 5 menit.",
			Status:   "APPROVED",
		},
	}
	for _, d := range defaults {
		var existing WaTemplate
		if err := config.DB.Where("code = ?", d.Code).First(&existing).Error; err != nil {
			d.CreatedAt = time.Now()
			d.UpdatedAt = time.Now()
			_ = config.DB.Create(&d).Error
		}
	}
	return nil
}

func GetPaginatedWaTemplate(offset, limit int, search string) ([]WaTemplate, int64, error) {
	var list []WaTemplate
	var total int64
	db := config.DB.Model(&WaTemplate{})
	if search != "" {
		db = db.Where("name ILIKE ? OR code ILIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Preload(clause.Associations).Order("id asc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetWaTemplateByID(id uint) (*WaTemplate, error) {
	var data WaTemplate
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateWaTemplate(data *WaTemplate) error {
	return config.DB.Save(data).Error
}

func DeleteWaTemplate(id uint) error {
	return config.DB.Delete(&WaTemplate{}, id).Error
}

// Gateway Config
func GetActiveWhatsappConfig() (*WhatsappConfig, error) {
	var cfg WhatsappConfig
	err := config.DB.Where("is_active = ?", true).Order("id desc").First(&cfg).Error
	if err != nil {
		cfg = WhatsappConfig{
			Provider:   "WABA_CLOUD",
			ApiToken:   "",
			PhoneID:    "",
			WebhookURL: "https://erp.perusahaan.co.id/api/whatsapp/webhook",
			IsActive:   true,
		}
		_ = config.DB.Create(&cfg).Error
		return &cfg, nil
	}
	return &cfg, nil
}

func SaveWhatsappConfig(data *WhatsappConfig) error {
	data.IsActive = true
	var existing WhatsappConfig
	if err := config.DB.Order("id desc").First(&existing).Error; err == nil {
		data.ID = existing.ID
		data.CreatedAt = existing.CreatedAt
		if data.ApiToken == "" {
			data.ApiToken = existing.ApiToken
		}
		return config.DB.Save(data).Error
	}
	return config.DB.Create(data).Error
}

// WaLog
func CreateWaLog(data *WaLog) error {
	if data.SentAt.IsZero() {
		data.SentAt = time.Now()
	}
	return config.DB.Create(data).Error
}

func GetPaginatedWaLogs(offset, limit int, search string) ([]WaLog, int64, error) {
	var list []WaLog
	var total int64
	db := config.DB.Model(&WaLog{})
	if search != "" {
		db = db.Where("phone ILIKE ? OR recipient ILIKE ? OR template_name ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}
