package pdfgen

import (
	"bytes"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// GeneratePDF merender string HTML menjadi byte PDF asli menggunakan wkhtmltopdf.
// Mendukung ukuran kertas (misal: "A4", "A5", "Letter") dan orientasi ("Portrait", "Landscape").
func GeneratePDF(htmlContent string, pageSize string, orientation string) ([]byte, error) {
	log.Println("Memulai render PDF dari HTML...")

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Printf("Gagal inisialisasi PDF Generator: %v\n", err)
		return nil, err
	}

	// 1. Pengaturan Kertas & Orientasi
	if pageSize == "" {
		pageSize = "A4"
	}
	if orientation == "" {
		orientation = "Portrait"
	}
	pdfg.PageSize.Set(pageSize)
	pdfg.Orientation.Set(orientation)
	
	// Pengaturan Margin Standar
	pdfg.MarginTop.Set(10)
	pdfg.MarginBottom.Set(10)
	pdfg.MarginLeft.Set(10)
	pdfg.MarginRight.Set(10)

	// 2. Buat halaman dari HTML
	page := wkhtmltopdf.NewPageReader(bytes.NewReader([]byte(htmlContent)))
	page.EnableLocalFileAccess.Set(true)
	pdfg.AddPage(page)

	// 3. Merender HTML menjadi PDF
	err = pdfg.Create()
	if err != nil {
		log.Printf("Gagal membuat PDF: %v\n", err)
		return nil, err
	}

	log.Println("File PDF berhasil dibuat!")
	return pdfg.Bytes(), nil
}
