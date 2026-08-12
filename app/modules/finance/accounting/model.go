package accounting

import (
	"ERP-System/config"
	"time"
)

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Code      string    `gorm:"type:varchar(50);not null;unique" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"type:varchar(50);not null" json:"type"` // e.g. receivable, payable, bank, income, expense
	CreatedAt time.Time `json:"created_at"`
}

type Journal struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(10);not null;unique" json:"code"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Type string `gorm:"type:varchar(50);not null" json:"type"` // sale, purchase, cash, bank, general
}

type JournalEntry struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. INV/2026/001
	JournalID uint      `json:"journal_id"`
	Date      time.Time `json:"date"`
	State     string    `gorm:"type:varchar(20);default:'draft'" json:"state"` // draft, posted
	CreatedAt time.Time `json:"created_at"`
}

type JournalItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	EntryID   uint    `json:"entry_id"`
	AccountID uint    `json:"account_id"`
	Name      string  `gorm:"type:varchar(255)" json:"name"`
	Debit     float64 `gorm:"type:numeric(15,2);default:0" json:"debit"`
	Credit    float64 `gorm:"type:numeric(15,2);default:0" json:"credit"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Account{}, &Journal{}, &JournalEntry{}, &JournalItem{})
}
