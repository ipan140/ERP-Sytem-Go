package consolidation

import (
	"ERP-System/app/modules/finance/accounting"
	"ERP-System/config"
	"time"
)

type ConsolidationReport struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Konsolidasi Q1 2026
	Period    string    `gorm:"type:varchar(50)" json:"period"`
	Branches  string    `gorm:"type:text" json:"branches"` // Simpan JSON array string cabang terpilih
	CreatedAt time.Time `json:"created_at"`
}

type ConsolidatedAccount struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ReportID  uint    `json:"report_id"`
	Report *ConsolidationReport `gorm:"foreignKey:ReportID" json:"report,omitempty"` // Odoo relation mapped
	AccountID uint    `json:"account_id"`
	Account *accounting.Account `gorm:"foreignKey:AccountID" json:"account,omitempty"` // Odoo relation mapped
	Balance   float64 `gorm:"type:numeric(15,2)" json:"balance"` // Saldo gabungan anak perusahaan
}


func (ConsolidationReport) TableName() string {
	return "finance.consolidation_reports"
}

func (ConsolidatedAccount) TableName() string {
	return "finance.consolidated_accounts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ConsolidationReport{}, &ConsolidatedAccount{})
}
