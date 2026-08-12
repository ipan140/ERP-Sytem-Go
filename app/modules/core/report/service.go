package report

import (
	"ERP-System/pkg/pdfgen"
	"log"
)

// CreateInvoicePDF adalah contoh *service* untuk merakit dan mencetak PDF tagihan.
// Nantinya modul `invoicing` atau `sales` akan memanggil fungsi ini.
func CreateInvoicePDF(invoiceID uint) ([]byte, error) {
	log.Printf("Mengumpulkan data untuk Invoice #%d...", invoiceID)
	
	// 1. Ambil data invoice dari database (placeholder)
	// 2. Render template HTML dengan data invoice tersebut
	htmlContent := "<h1>Invoice #" + string(rune(invoiceID)) + "</h1><p>Total: Rp 100.000</p>"

	// 3. Panggil pkg pdfgen untuk mengubah HTML menjadi PDF secara nyata
	pdfBytes, err := pdfgen.GeneratePDF(htmlContent)
	if err != nil {
		log.Printf("❌ Gagal memroses PDF Invoice: %v", err)
		return nil, err
	}

	return pdfBytes, nil
}
