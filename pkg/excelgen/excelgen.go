package excelgen

import (
	"bytes"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func GenerateExcel(sheetName string, headers []string, rows [][]interface{}) ([]byte, error) {
	log.Println("Membuat dokumen Excel untuk:", sheetName)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Error closing Excel file:", err)
		}
	}()

	// 1. Ubah nama Sheet utama
	f.SetSheetName("Sheet1", sheetName)

	// 2. Tulis baris Header (Baris 1)
	for colIndex, header := range headers {
		// CoordinatesToCellName(1, 1) -> "A1"
		cell, _ := excelize.CoordinatesToCellName(colIndex+1, 1)
		f.SetCellValue(sheetName, cell, header)
	}

	// 3. Tulis Isi Data (Mulai dari Baris 2)
	for rowIndex, row := range rows {
		for colIndex, cellValue := range row {
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, rowIndex+2)
			f.SetCellValue(sheetName, cell, cellValue)
		}
	}

	var buffer bytes.Buffer
	if err := f.Write(&buffer); err != nil {
		return nil, fmt.Errorf("gagal merender file excel: %v", err)
	}

	return buffer.Bytes(), nil
}

// GenerateExcelFromTemplate merender dokumen Excel menggunakan file template .xlsx yang sudah ada.
func GenerateExcelFromTemplate(templatePath string, sheetName string, staticData map[string]interface{}, tableData [][]interface{}, tableStartRow int) ([]byte, error) {
	log.Println("Membuka template Excel:", templatePath)

	// Membuka file excel template yang ada di direktori project
	f, err := excelize.OpenFile(templatePath)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka template excel: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("Error closing Excel file:", err)
		}
	}()

	// 1. Tulis data spesifik / statis ke sel tertentu 
	// (misal mengisi kop surat, nama kustomer di cell B3, tanggal di C3, dsb.)
	for cell, value := range staticData {
		f.SetCellValue(sheetName, cell, value)
	}

	// 2. Tulis baris isi Data / Tabel (Mulai dari baris yang ditentukan: tableStartRow)
	for rowIndex, row := range tableData {
		for colIndex, cellValue := range row {
			// colIndex+1 = Kolom A (1), B (2), dst.
			cell, _ := excelize.CoordinatesToCellName(colIndex+1, tableStartRow+rowIndex)
			f.SetCellValue(sheetName, cell, cellValue)
		}
	}

	// 3. Simpan perubahan ke buffer untuk didownload
	var buffer bytes.Buffer
	if err := f.Write(&buffer); err != nil {
		return nil, fmt.Errorf("gagal menyimpan hasil render excel: %v", err)
	}

	return buffer.Bytes(), nil
}
