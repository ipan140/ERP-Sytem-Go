package report

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateReport(data *Report) error {
	return config.DB.Create(data).Error
}

func GetAllReport() ([]Report, error) {
	var list []Report
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetAllPrintTemplates() ([]PrintTemplate, error) {
	var list []PrintTemplate
	err := config.DB.Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultPrintTemplates()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultPrintTemplates() error {
	defaults := []PrintTemplate{
		{Name: "Faktur Penjualan (Commercial Invoice A4)", Code: "TPL-INV-01", Module: "Finance & Sales", PaperSize: "A4", Orientation: "Portrait", HasHeader: true, IsDefault: true, Icon: "🧾"},
		{Name: "Slip Gaji Karyawan TER 2024 (Payslip)", Code: "TPL-PAY-02", Module: "HRD & Payroll", PaperSize: "A4 / Half", Orientation: "Portrait", HasHeader: true, IsDefault: true, Icon: "💵"},
		{Name: "Pesanan Pembelian Vendor (Purchase Order PO)", Code: "TPL-PO-03", Module: "Procurement SCM", PaperSize: "A4", Orientation: "Portrait", HasHeader: true, IsDefault: true, Icon: "📋"},
		{Name: "Surat Jalan / Bukti Pengiriman Barang (Delivery Order)", Code: "TPL-DO-04", Module: "Warehouse SCM", PaperSize: "Continuous 9.5x11\"", Orientation: "Landscape", HasHeader: false, IsDefault: true, Icon: "🚚"},
		{Name: "Bukti Kas Keluar / Payment Voucher", Code: "TPL-VCH-05", Module: "Finance Cash", PaperSize: "A5", Orientation: "Landscape", HasHeader: true, IsDefault: true, Icon: "🎫"},
		{Name: "Form Penawaran Resmi (Quotation Letter)", Code: "TPL-QUO-06", Module: "Sales CRM", PaperSize: "A4", Orientation: "Portrait", HasHeader: true, IsDefault: true, Icon: "💼"},
	}
	for _, tpl := range defaults {
		var existing PrintTemplate
		if err := config.DB.Where("code = ?", tpl.Code).First(&existing).Error; err != nil {
			_ = config.DB.Create(&tpl).Error
		}
	}
	return nil
}

func GetAllExportReports() ([]ExportReportItem, error) {
	var list []ExportReportItem
	err := config.DB.Order("id asc").Find(&list).Error
	if err == nil && len(list) == 0 {
		_ = SeedDefaultExportReports()
		_ = config.DB.Order("id asc").Find(&list).Error
	}
	return list, err
}

func SeedDefaultExportReports() error {
	defaults := []ExportReportItem{
		{Title: "Laporan Laba Rugi & Neraca Konsolidasi SAK", Frequency: "Bulanan / Kuartalan", Description: "Neraca lajur, arus kas operasional, dan performa profitabilitas bulanan.", LastRun: "31 Agustus 2025", Size: "1.4 MB", Icon: "📈"},
		{Title: "Rekapitulasi Absensi, Lembur & Potongan PPh 21 Karyawan", Frequency: "Bulanan (Periode 25-25)", Description: "Detail 248 pegawai aktif dengan tarif TER 2024, iuran BPJS TK & BPJS Kesehatan.", LastRun: "01 September 2025", Size: "890 KB", Icon: "👥"},
		{Title: "Laporan Mutasi Persediaan & Nilai Valuasi FIFO Gudang", Frequency: "Mingguan", Description: "Daftar pergerakan SKU keluar-masuk, safety stock, dan dead stock gudang Cikarang & Surabaya.", LastRun: "08 September 2025", Size: "2.8 MB", Icon: "📦"},
		{Title: "Ringkasan Pipeline Penjualan & Realisasi Omzet per Sales Rep", Frequency: "Mingguan", Description: "Konversi leads, win rate deal, dan target kuota revenue kuartal berjalan.", LastRun: "12 September 2025", Size: "620 KB", Icon: "🎯"},
	}
	for _, exp := range defaults {
		var existing ExportReportItem
		if err := config.DB.Where("title = ?", exp.Title).First(&existing).Error; err != nil {
			_ = config.DB.Create(&exp).Error
		}
	}
	return nil
}

func CreatePrintTemplate(data *PrintTemplate) error {
	return config.DB.Create(data).Error
}

func GetReportByID(id uint) (*Report, error) {
	var data Report
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateReport(data *Report) error {
	return config.DB.Save(data).Error
}

func DeleteReport(id uint) error {
	return config.DB.Delete(&Report{}, id).Error
}
