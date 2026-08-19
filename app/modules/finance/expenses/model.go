package expenses

import (
	"ERP-System/config"
	"time"
)

type ExpenseSheet struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Business Trip to Jakarta
	EmployeeID uint      `json:"employee_id"`
	Total      float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, submit, approve, post, done
	CreatedAt  time.Time `json:"created_at"`
}

type Expense struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. Flight Ticket
	ExpenseSheetID *uint     `json:"expense_sheet_id"`
	EmployeeID     uint      `json:"employee_id"`
	TotalAmount    float64   `gorm:"type:numeric(15,2);default:0" json:"total_amount"`
	State          string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, reported
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ExpenseSheet{}, &Expense{})
}
