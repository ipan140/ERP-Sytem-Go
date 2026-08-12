package pdfgen

import (
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// GeneratePDF merender string HTML menjadi byte PDF asli menggunakan wkhtmltopdf.
// Fungsi ini bersifat generik dan bisa dipanggil oleh modul mana pun (misal: Report Invoice, Payslip).
func GeneratePDF(htmlContent string) ([]byte, error) {
	log.Println("Memulai render PDF dari HTML...")

	// Inisialisasi generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Printf("❌ Gagal inisialisasi PDF Generator (Pastikan software wkhtmltopdf terinstal di OS Anda): %v\n", err)
		return nil, err
	}

	// Buat halaman (page) pembaca dari HTML string
	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))

	// Set opsi halaman agar terlihat rapi
	page.EnableLocalFileAccess.Set(true)

	// Tambahkan halaman ke generator
	pdfg.AddPage(page)

	// Proses pembuatan PDF (Merender HTML menjadi PDF)
	err = pdfg.Create()
	if err != nil {
		log.Printf("❌ Gagal membuat PDF: %v\n", err)
		return nil, err
	}

	log.Println("✅ File PDF berhasil dibuat dari HTML!")

	// Kembalikan hasil buffer sebagai byte array yang siap dikirim via API
	return pdfg.Bytes(), nil
}
