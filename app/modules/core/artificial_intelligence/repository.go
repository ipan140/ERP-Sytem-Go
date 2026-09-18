package artificial_intelligence

import (
	"strings"
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateAIPrompt(data *AIPrompt) error {
	if data.Title == "" && data.Name != "" {
		data.Title = data.Name
	}
	if data.Name == "" && data.Title != "" {
		data.Name = data.Title
	}
	return config.DB.Create(data).Error
}

func GetAllAIPrompt() ([]AIPrompt, error) {
	var list []AIPrompt
	err := config.DB.Preload(clause.Associations).Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultAIPrompts()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultAIPrompts() error {
	defaults := []AIPrompt{
		{
			Title:        "Evaluasi Kinerja & Analisis KPI Pegawai Periode Kuartal",
			Name:         "Evaluasi Kinerja & Analisis KPI Pegawai Periode Kuartal",
			Module:       "HR & Payroll",
			Model:        "Gemini 1.5 Pro",
			Description:  "Menganalisis pencapaian target kerja, presensi absensi, dan memberikan umpan balik konstruktif bagi evaluasi 360 derajat.",
			SystemPrompt: "Anda adalah HR Talent Development Specialist kelas dunia di perusahaan dengan 500+ karyawan.",
			SampleInput:  "Nama Karyawan: Siti Aminah\nJabatan: Senior Account Executive\nPencapaian Target: 114% (Rp 1.14 Miliar dari target Rp 1 Miliar)\nKehadiran: 98% Tepat Waktu, Cuti 3 Hari\nFeedback Tim: Kolaboratif, proaktif memandu junior, kemampuan negosiasi tinggi.",
			SampleOutput: "<h4 class=\"font-bold text-gray-900 dark:text-white mb-1\">🎯 Rekomendasi Evaluasi Kinerja (Performance Review):</h4><p><strong>Status Penilaian:</strong> <span class=\"text-emerald-600 font-bold\">EXCEEDS EXPECTATIONS (Rating A)</span></p><ul class=\"list-disc pl-5 mt-2 space-y-1 text-xs\"><li><strong>Pencapaian Omzet:</strong> Lampaui kuota 114% (+Rp 140 Juta) menunjukkan kemampuan penetrasi pasar yang sangat tangguh.</li><li><strong>Disiplin & Budaya Kerja:</strong> Rasio presensi 98% membuktikan komitmen operasional yang stabil.</li><li><strong>Rencana Pengembangan:</strong> Direkomendasikan masuk dalam program <em>Fast-Track Leadership Development</em> dan nominasi kenaikan tunjangan jabatan pada peninjauan akhir tahun.</li></ul>",
			Icon:         "👥",
		},
		{
			Title:        "Deteksi Dini Anomali & Rekonsiliasi Jurnal Keuangan",
			Name:         "Deteksi Dini Anomali & Rekonsiliasi Jurnal Keuangan",
			Module:       "Finance & Accounting",
			Model:        "Gemini 1.5 Pro",
			Description:  "Memeriksa pola mutasi kas bank, membandingkan faktur pajak keluaran vs masukan, dan mendeteksi potensi duplikasi entri ledger.",
			SystemPrompt: "Anda adalah Chief Financial Auditor berpengalaman dalam mendeteksi fraud dan anomali akuntansi standar PSAK.",
			SampleInput:  "Mutasi Rekening Bank BCA: Debit Rp 125.000.000 (Vendor PT Baja Logam)\nJurnal Ledger Sistem: Kredit Rp 125.000.000 pada akun Hutang Usaha\nFaktur Pajak Masukan: Ditemukan 2 nomor FP yang persis sama diupload selisih 14 menit.",
			SampleOutput: "<h4 class=\"font-bold text-gray-900 dark:text-white mb-1\">⚠️ Peringatan Audit Keuangan:</h4><p><strong>Tingkat Risiko:</strong> <span class=\"text-amber-600 font-bold\">SEDANG (Potensi Duplikasi Faktur Masukan)</span></p><ul class=\"list-disc pl-5 mt-2 space-y-1 text-xs\"><li><strong>Temuan:</strong> Terdapat indikasi <em>Double Posting</em> faktur pajak masukan senilai PPN Rp 13.750.000.</li><li><strong>Dampak:</strong> Pelaporan SPT Masa PPN berpotensi mengalami selisih faktur ganda di DJP e-Faktur.</li><li><strong>Rekomendasi Tindakan:</strong> Batalkan salah satu draft faktur masukan sebelum proses closing SPT akhir bulan. Rekonsiliasi mutasi kas sudah sesuai 100%.</li></ul>",
			Icon:         "🔍",
		},
		{
			Title:        "Prediksi Kebutuhan Safety Stock Menghadapi Tren Musiman",
			Name:         "Prediksi Kebutuhan Safety Stock Menghadapi Tren Musiman",
			Module:       "Supply Chain",
			Model:        "Gemini 1.5 Pro",
			Description:  "Menghitung buffer stock gudang, lead time vendor, dan reorder point optimal menjelang puncak permintaan kuartal.",
			SystemPrompt: "Anda adalah Kepala Perencanaan Logistik & Supply Chain Terpadu.",
			SampleInput:  "SKU: MAT-STEEL-01 (Bahan Baku Utama)\nStok Saat Ini: 420 Unit\nKonsumsi Rata-rata: 50 Unit / Hari\nLead Time Pengiriman Vendor: 14 Hari\nProyeksi Kenaikan Order Q4: +35%",
			SampleOutput: "<h4 class=\"font-bold text-gray-900 dark:text-white mb-1\">📦 Rekomendasi Pengadaan Persediaan (SCM):</h4><p><strong>1. Estimasi Kebutuhan Harian Baru:</strong> 50 unit x 1.35 = <strong>67.5 Unit / Hari</strong>.</p><p><strong>2. Lead Time Demand:</strong> 67.5 x 14 hari = <strong>945 Unit</strong>.</p><p><strong>3. Status Stok Kritis:</strong> Stok saat ini (420 unit) hanya cukup untuk <strong>6.2 hari ke depan</strong>. Risiko <em>stockout</em> dalam 8 hari ke depan!</p><p><strong>4. Tindakan Segera:</strong> Rilis Purchase Order (PO) Darurat minimal <strong>1.200 Unit</strong> ke vendor rekanan terverifikasi hari ini.</p>",
			Icon:         "📦",
		},
		{
			Title:        "Strategi Negosiasi & Closing Deals Penjualan B2B Nilai Tinggi",
			Name:         "Strategi Negosiasi & Closing Deals Penjualan B2B Nilai Tinggi",
			Module:       "Sales & CRM",
			Model:        "Gemini 1.5 Pro",
			Description:  "Menyusun naskah penawaran win-win, mengatasi keberatan diskon harga, dan menyusun klausul kontrak tahunan.",
			SystemPrompt: "Anda adalah VP of Enterprise Sales berpengalaman menutup kontrak B2B multi-miliar.",
			SampleInput:  "Klien: PT Manufaktur Logam Prima\nNilai Deal: Rp 1.2 Miliar / Tahun\nKeberatan: \"Kompetitor menawarkan harga 15% lebih murah, anggaran kami terbatas kuartal ini.\"",
			SampleOutput: "<h4 class=\"font-bold text-gray-900 dark:text-white mb-1\">💼 Strategi Negosiasi B2B:</h4><p><strong>Taktik: Jangan Menurunkan Harga Pokok, Ubah Struktur Nilai (Value Restructuring)</strong></p><p><strong>Draf Respon ke Klien:</strong></p><blockquote class=\"border-l-2 border-brand-500 pl-3 italic text-gray-600 dark:text-gray-300 mt-1\">\"Kami sangat memahami pertimbangan anggaran PT Manufaktur Logam Prima. Solusi kami mencakup garansi SLA 99.9% dan dukungan onsite engineer 24/7 yang rata-rata menghemat downtime senilai Rp 300 Juta/tahun dibandingkan opsi lain. Sebagai solusi kuartal ini, kami tawarkan skema pembayaran kuartalan (Rp 300 Juta x 4 termin) ditambah bonus gratis modul IoT Biometrik.\"</blockquote>",
			Icon:         "🎯",
		},
	}
	for _, p := range defaults {
		var existing AIPrompt
		if err := config.DB.Where("title = ? OR name = ?", p.Title, p.Name).First(&existing).Error; err != nil {
			_ = config.DB.Create(&p).Error
		}
	}
	return nil
}

func GetAIPromptByID(id uint) (*AIPrompt, error) {
	var data AIPrompt
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateAIPrompt(data *AIPrompt) error {
	return config.DB.Save(data).Error
}

func DeleteAIPrompt(id uint) error {
	return config.DB.Delete(&AIPrompt{}, id).Error
}

func GetActiveAIConfig() (*AIConfig, error) {
	var cfg AIConfig
	err := config.DB.First(&cfg).Error
	if err != nil {
		cfg = AIConfig{
			ActiveProvider:  "GEMINI",
			GeminiModel:     "gemini-1.5-flash",
			OpenAiModel:     "gpt-4o-mini",
			OpenAiBaseUrl:   "https://api.openai.com/v1",
			ClaudeModel:     "claude-3-5-sonnet-20241022",
			DeepseekModel:   "deepseek-chat",
			DeepseekBaseUrl: "https://api.deepseek.com/v1",
			OllamaBaseUrl:   "http://localhost:11434",
			OllamaModel:     "llama3",
		}
		_ = config.DB.Create(&cfg).Error
	}
	return &cfg, nil
}

func SaveAIConfig(data *AIConfig) error {
	var existing AIConfig
	err := config.DB.First(&existing).Error
	if err != nil {
		return config.DB.Create(data).Error
	}

	preserveIfEmpty := func(newVal, oldVal string) string {
		trimmed := strings.TrimSpace(newVal)
		if trimmed == "" || strings.HasPrefix(trimmed, "•") || strings.HasPrefix(trimmed, "*") {
			return oldVal
		}
		return trimmed
	}

	if data.ActiveProvider != "" {
		existing.ActiveProvider = data.ActiveProvider
	}
	if data.GeminiModel != "" {
		existing.GeminiModel = data.GeminiModel
	}
	existing.GeminiApiKey = preserveIfEmpty(data.GeminiApiKey, existing.GeminiApiKey)

	if data.OpenAiModel != "" {
		existing.OpenAiModel = data.OpenAiModel
	}
	if data.OpenAiBaseUrl != "" {
		existing.OpenAiBaseUrl = data.OpenAiBaseUrl
	}
	existing.OpenAiApiKey = preserveIfEmpty(data.OpenAiApiKey, existing.OpenAiApiKey)

	if data.ClaudeModel != "" {
		existing.ClaudeModel = data.ClaudeModel
	}
	existing.ClaudeApiKey = preserveIfEmpty(data.ClaudeApiKey, existing.ClaudeApiKey)

	if data.DeepseekModel != "" {
		existing.DeepseekModel = data.DeepseekModel
	}
	if data.DeepseekBaseUrl != "" {
		existing.DeepseekBaseUrl = data.DeepseekBaseUrl
	}
	existing.DeepseekApiKey = preserveIfEmpty(data.DeepseekApiKey, existing.DeepseekApiKey)

	if data.OllamaModel != "" {
		existing.OllamaModel = data.OllamaModel
	}
	if data.OllamaBaseUrl != "" {
		existing.OllamaBaseUrl = data.OllamaBaseUrl
	}

	return config.DB.Save(&existing).Error
}
