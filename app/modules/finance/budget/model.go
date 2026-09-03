package budget

import (
	"ERP-System/config"
	"time"
)

type DepartmentBudget struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	DepartmentName string    `gorm:"type:varchar(100);not null" json:"department_name"` // Pemasaran, IT, HR, Operasional
	FiscalPeriod   string    `gorm:"type:varchar(50);not null" json:"fiscal_period"`    // Q1 2026
	AllocatedLimit float64   `gorm:"type:numeric(15,2);not null" json:"allocated_limit"` // Pagu Anggaran
	RealizedSpent  float64   `gorm:"type:numeric(15,2);default:0" json:"realized_spent"`  // Realisasi
	RemainingBudget float64  `gorm:"type:numeric(15,2);default:0" json:"remaining_budget"`
	UsagePercent   float64   `gorm:"type:numeric(5,2);default:0" json:"usage_percent"`
	Status         string    `gorm:"type:varchar(30);default:'Aman'" json:"status"` // Aman, Peringatan, Overbudget
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &DepartmentBudget{})
}

func (DepartmentBudget) TableName() string {
	return "finance.department_budgets"
}
