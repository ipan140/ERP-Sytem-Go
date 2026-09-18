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


type PrintTemplate struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	Code         string    `gorm:"type:varchar(100);default:''" json:"code"`
	Module       string    `gorm:"type:varchar(100);default:'Finance & Sales'" json:"module"`
	PaperSize    string    `gorm:"type:varchar(50);default:'A4'" json:"paper_size"`
	Orientation  string    `gorm:"type:varchar(50);default:'Portrait'" json:"orientation"`
	HasHeader    bool      `gorm:"default:true" json:"has_header"`
	IsDefault    bool      `gorm:"default:true" json:"is_default"`
	Icon         string    `gorm:"type:varchar(50);default:'🧾'" json:"icon"`
	HTMLTemplate string    `gorm:"type:text" json:"html_template"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type ExportReportItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Frequency   string    `gorm:"type:varchar(100);default:'Bulanan'" json:"frequency"`
	Description string    `gorm:"type:text" json:"description"`
	LastRun     string    `gorm:"type:varchar(100);default:'Baru saja'" json:"last_run"`
	Size        string    `gorm:"type:varchar(50);default:'1.2 MB'" json:"size"`
	Icon        string    `gorm:"type:varchar(50);default:'📈'" json:"icon"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (PrintTemplate) TableName() string {
	return "setting.print_templates"
}

func (ExportReportItem) TableName() string {
	return "setting.export_reports"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Report{}, &PrintTemplate{}, &ExportReportItem{})
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
