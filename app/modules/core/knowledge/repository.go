package knowledge

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateArticle(data *Article) error {
	if data.Title == "" && data.Name != "" {
		data.Title = data.Name
	}
	if data.Name == "" && data.Title != "" {
		data.Name = data.Title
	}
	return config.DB.Create(data).Error
}

func GetAllArticle() ([]Article, error) {
	var list []Article
	err := config.DB.Preload(clause.Associations).Order("id desc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultArticles()
		_ = config.DB.Order("id desc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultArticles() error {
	defaults := []Article{
		{
			Title:       "Buku Panduan Karyawan Baru (Employee Handbook 2025)",
			Name:        "Buku Panduan Karyawan Baru (Employee Handbook 2025)",
			Category:    "HR & Kepegawaian",
			Summary:     "Pedoman lengkap orientasi dan budaya korporat bagi 100-1000 karyawan: jam kerja, dress code, tunjangan BPJS, dan hak asuransi.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Jam Kerja & Presensi</h4><p>Jam kerja reguler adalah Senin s.d. Jumat pukul 08.30 - 17.30 WIB dengan istirahat 1 jam. Presensi wajib dilakukan melalui mesin biometrik atau mobile GPS portal sebelum pukul 08.45 WIB.</p><h4 class=\"font-bold text-gray-900 dark:text-white mt-3\">2. Tunjangan & Asuransi</h4><p>Seluruh pegawai tetap dan PKWT didaftarkan pada BPJS Ketenagakerjaan (JKK, JKM, JHT, JP) dan BPJS Kesehatan sejak hari pertama kerja. Plafon rawat inap tambahan berlaku setelah masa percobaan 3 bulan.</p>",
			Author:      "HR People & Culture Lead",
			ReadingTime: "12 Menit",
			Status:      "PUBLISHED",
		},
		{
			Title:       "SOP Pengajuan Cuti Tahunan, Izin Sakit & Cuti Melahirkan",
			Name:        "SOP Pengajuan Cuti Tahunan, Izin Sakit & Cuti Melahirkan",
			Category:    "HR & Kepegawaian",
			Summary:     "Mekanisme mandiri pengajuan kuota 12 hari cuti tahunan, dispensasi khusus berbayar, dan SLA persetujuan atasan langsung.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Prosedur Cuti Tahunan</h4><p>Pengajuan wajib diajukan minimal 3 hari kerja sebelum tanggal pelaksanaan melalui menu Self-Service Portal Karyawan.</p><h4 class=\"font-bold text-gray-900 dark:text-white mt-3\">2. Izin Sakit</h4><p>Izin sakit lebih dari 1 (satu) hari kalender wajib melampirkan Surat Keterangan Dokter berizin resmi dengan diagnosa medis.</p>",
			Author:      "HR Operations",
			ReadingTime: "5 Menit",
			Status:      "PUBLISHED",
		},
		{
			Title:       "Pedoman Klaim Reimbursement Rawat Jalan & Perjalanan Dinas",
			Name:        "Pedoman Klaim Reimbursement Rawat Jalan & Perjalanan Dinas",
			Category:    "Finance & Pajak",
			Summary:     "Ketentuan nota kuitansi bermeterai, plafon kacamata tahunan Rp 1.500.000, serta uang saku perjalanan dinas luar kota.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Nota & Bukti Pembayaran</h4><p>Semua kuitansi rawat jalan wajib mencantumkan nama pegawai dan cap stempel klinik/rumah sakit. Batas pengajuan adalah tanggal 20 setiap bulannya.</p><h4 class=\"font-bold text-gray-900 dark:text-white mt-3\">2. Perjalanan Dinas</h4><p>Tiket transportasi dan hotel bintang 3/4 dipesan terpusat oleh General Affairs. Uang saku harian ditransfer via Payroll pada hari kepulangan.</p>",
			Author:      "Finance & Accounting Lead",
			ReadingTime: "7 Menit",
			Status:      "PUBLISHED",
		},
		{
			Title:       "SOP Pengadaan Barang & Approval Purchase Order (PO) > Rp 50 Juta",
			Name:        "SOP Pengadaan Barang & Approval Purchase Order (PO) > Rp 50 Juta",
			Category:    "SCM & Gudang",
			Summary:     "Kebijakan 3 vendor pembanding (tender komparatif) dan matriks tanda tangan otorisasi berjenjang untuk integritas pengeluaran.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Batas Nominal & Matriks Approval</h4><p>Nilai pengadaan Rp 1 - 50 Juta disetujui Department Head. Nilai > Rp 50 Juta wajib melampirkan minimal 3 penawaran vendor dan disetujui oleh Direktur Keuangan.</p>",
			Author:      "Procurement Specialist",
			ReadingTime: "8 Menit",
			Status:      "PUBLISHED",
		},
		{
			Title:       "Standar K3, APD, dan Keselamatan Gudang Distribusi Cikarang",
			Name:        "Standar K3, APD, dan Keselamatan Gudang Distribusi Cikarang",
			Category:    "SCM & Gudang",
			Summary:     "Prosedur keselamatan kerja area bongkar muat, kewajiban sepatu safety boot, helm pelindung, dan batas kecepatan forklift 10 km/jam.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Alat Pelindung Diri (APD)</h4><p>Setiap orang yang memasuki zona gudang wajib mengenakan helm keselamatan, rompi high-visibility, dan sepatu boot berujung baja (steel-toe).</p>",
			Author:      "Warehouse HSE Officer",
			ReadingTime: "10 Menit",
			Status:      "PUBLISHED",
		},
		{
			Title:       "Protokol Keamanan Password, Otentikasi 2FA, & Anti-Phishing",
			Name:        "Protokol Keamanan Password, Otentikasi 2FA, & Anti-Phishing",
			Category:    "IT & Keamanan",
			Summary:     "Standar kata sandi minimum 12 karakter alfanumerik, aktivasi TOTP authenticator untuk hak akses ERP, dan cara lapor insiden siber.",
			Content:     "<h4 class=\"font-bold text-gray-900 dark:text-white\">1. Kebijakan Kredensial</h4><p>Kata sandi akun ERP kedaluwarsa otomatis setiap 90 hari. Dilarang menggunakan password berulang atau membagikan akun kepada rekan kerja.</p>",
			Author:      "Chief Information Security Officer",
			ReadingTime: "6 Menit",
			Status:      "PUBLISHED",
		},
	}
	for _, a := range defaults {
		var existing Article
		if err := config.DB.Where("title = ? OR name = ?", a.Title, a.Name).First(&existing).Error; err != nil {
			_ = config.DB.Create(&a).Error
		}
	}
	return nil
}

func GetArticleByID(id uint) (*Article, error) {
	var data Article
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateArticle(data *Article) error {
	return config.DB.Save(data).Error
}

func DeleteArticle(id uint) error {
	return config.DB.Delete(&Article{}, id).Error
}
