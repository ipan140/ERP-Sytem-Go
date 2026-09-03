package tax

import (
	"ERP-System/config"
	"time"
)

type TaxMasterConfig struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaxCode     string    `gorm:"type:varchar(50);unique;not null" json:"tax_code"` // e.g., PPH23, PPN11, PPH42
	TaxName     string    `gorm:"type:varchar(100);not null" json:"tax_name"`
	Rate        float64   `gorm:"type:numeric(5,2);not null" json:"rate"`
	LegalBasis  string    `gorm:"type:varchar(255)" json:"legal_basis"`             // Dasar Hukum / UU
	Description string    `gorm:"type:text" json:"description"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaxReportSummary struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaxPeriod   string    `gorm:"type:varchar(50)" json:"tax_period"` // e.g. Maret 2026
	TaxType     string    `gorm:"type:varchar(50)" json:"tax_type"`   // PPh 21 TER, PPh 23, PPh 4(2), PPN
	TaxBase     float64   `gorm:"type:numeric(15,2)" json:"tax_base"` // DPP
	TaxRate     float64   `gorm:"type:numeric(5,2)" json:"tax_rate"`   // Persentase Tarif
	TaxAmount   float64   `gorm:"type:numeric(15,2)" json:"tax_amount"`
	PartnerName string    `gorm:"type:varchar(150)" json:"partner_name"`
	NPWP        string    `gorm:"type:varchar(50)" json:"npwp"`
	Status      string    `gorm:"type:varchar(30);default:'Siap Lapor'" json:"status"` // Siap Lapor, Dilaporkan
	CreatedAt   time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &TaxReportSummary{}, &TaxMasterConfig{})
}

func (TaxReportSummary) TableName() string {
	return "finance.tax_reports"
}

func (TaxMasterConfig) TableName() string {
	return "finance.tax_master_configs"
}
