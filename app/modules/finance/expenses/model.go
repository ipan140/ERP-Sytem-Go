package expenses

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type ExpenseSheet struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Business Trip to Jakarta
	EmployeeID uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	Total      float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, submit, approve, post, done
	CreatedAt  time.Time `json:"created_at"`
}

type Expense struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Flight Ticket
	ExpenseSheetID *uint     `json:"expense_sheet_id"`
	ExpenseSheet *ExpenseSheet `gorm:"foreignKey:ExpenseSheetID"` // Auto-added relation
	EmployeeID     uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	TotalAmount    float64   `gorm:"type:numeric(15,2);default:0" json:"total_amount"`
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, reported
	CreatedAt      time.Time `json:"created_at"`
}


type PettyCashFund struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. Kas Kecil Operasional Kantor
	Custodian    string    `gorm:"type:varchar(100);not null" json:"custodian"` // Pemegang kasir kas kecil
	PlafondLimit float64   `gorm:"type:numeric(15,2);default:5000000" json:"plafond_limit"` // Plafon sistem imprest
	CurrentBalance float64 `gorm:"type:numeric(15,2);default:5000000" json:"current_balance"`
	GLAccountID  uint      `json:"gl_account_id"` // Akun Kas Kecil (1-1001)
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PettyCashTransaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FundID      uint      `json:"fund_id"`
	Fund        *PettyCashFund `gorm:"foreignKey:FundID" json:"fund,omitempty"`
	TxType      string    `gorm:"type:varchar(20);not null" json:"tx_type"` // "expense" (pengeluaran), "replenish" (pengisian kembali)
	Amount      float64   `gorm:"type:numeric(15,2);not null" json:"amount"`
	Description string    `gorm:"type:varchar(255);not null" json:"description"` // e.g. Beli air galon & snack meeting
	ReceiptRef  string    `gorm:"type:varchar(100)" json:"receipt_ref"`
	RecordedBy  string    `gorm:"type:varchar(100)" json:"recorded_by"`
	CreatedAt   time.Time `json:"created_at"`
}

func (ExpenseSheet) TableName() string {
	return "finance.expense_sheets"
}

func (Expense) TableName() string {
	return "finance.expenses"
}

func (PettyCashFund) TableName() string {
	return "finance.petty_cash_funds"
}

func (PettyCashTransaction) TableName() string {
	return "finance.petty_cash_transactions"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ExpenseSheet{}, &Expense{}, &PettyCashFund{}, &PettyCashTransaction{})
}
