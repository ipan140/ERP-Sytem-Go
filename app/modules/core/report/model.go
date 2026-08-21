package report

import (
	"ERP-System/config"
	"time"
)

type Report struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Report{})
}

type DynamicExcelRequest struct {
	FileName  string          `json:"file_name"`
	SheetName string          `json:"sheet_name"`
	Headers   []string        `json:"headers"`
	Data      [][]interface{} `json:"data"`
}

type DynamicPDFRequest struct {
	FileName    string `json:"file_name"`
	HTMLContent string `json:"html_content"`
	PageSize    string `json:"page_size"`   // Contoh: "A4", "A5", "Letter"
	Orientation string `json:"orientation"` // Contoh: "Portrait", "Landscape"
}
