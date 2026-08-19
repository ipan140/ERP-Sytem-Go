package payroll

import (
	"ERP-System/config"
	"time"
)

type SalaryRule struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(255);not null" json:"name"` // e.g. Basic Salary, Tax Deduction
	Code string `gorm:"type:varchar(50);not null;unique" json:"code"` // BASIC, TAX
}

type Payslip struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"` // SLIP/2026/08/001
	EmployeeID uint      `json:"employee_id"`
	ContractID uint      `json:"contract_id"`
	DateFrom   time.Time `json:"date_from"`
	DateTo     time.Time `json:"date_to"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, done
	CreatedAt  time.Time `json:"created_at"`
}

type PayslipLine struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	PayslipID    uint    `json:"payslip_id"`
	SalaryRuleID uint    `json:"salary_rule_id"`
	Amount       float64 `gorm:"type:numeric(15,2);default:0" json:"amount"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &SalaryRule{}, &Payslip{}, &PayslipLine{})
}
