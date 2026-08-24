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


func (ExpenseSheet) TableName() string {
	return "finance.expense_sheets"
}

func (Expense) TableName() string {
	return "finance.expenses"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ExpenseSheet{}, &Expense{})
}
