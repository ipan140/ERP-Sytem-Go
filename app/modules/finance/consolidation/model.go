package consolidation

import (
	"ERP-System/config"
	"time"
)

type ConsolidationReport struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Konsolidasi Q1 2026
	Period    string    `gorm:"type:varchar(50)" json:"period"`
	CreatedAt time.Time `json:"created_at"`
}

type ConsolidatedAccount struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	ReportID  uint    `json:"report_id"`
	AccountID uint    `json:"account_id"`
	Balance   float64 `gorm:"type:numeric(15,2)" json:"balance"` // Saldo gabungan anak perusahaan
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ConsolidationReport{}, &ConsolidatedAccount{})
}
