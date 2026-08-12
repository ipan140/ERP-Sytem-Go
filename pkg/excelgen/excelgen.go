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
