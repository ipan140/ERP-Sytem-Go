package reconciliation

import (
	"ERP-System/config"
	"time"
)

type BankStatementItem struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Date           time.Time `json:"date"`
	Description    string    `gorm:"type:varchar(255)" json:"description"`
	RefNumber      string    `gorm:"type:varchar(100)" json:"ref_number"`
	Debit          float64   `gorm:"type:numeric(15,2);default:0" json:"debit"`
	Credit         float64   `gorm:"type:numeric(15,2);default:0" json:"credit"`
	BankName       string    `gorm:"type:varchar(50)" json:"bank_name"`
	IsReconciled   bool      `gorm:"default:false" json:"is_reconciled"`
	MatchedInvoice string    `gorm:"type:varchar(100)" json:"matched_invoice"`
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &BankStatementItem{})
}

func (BankStatementItem) TableName() string {
	return "finance.bank_statement_items"
}
