package timesheets

import (
	"ERP-System/app/modules/services/project"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Timesheet struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"` // Who worked
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	ProjectID   uint      `json:"project_id"`
	Project *project.Project `gorm:"foreignKey:ProjectID" json:"project,omitempty"` // Odoo relation mapped
	TaskID      *uint     `json:"task_id"` // Optional
	Task *project.Task `gorm:"foreignKey:TaskID" json:"task,omitempty"` // Odoo relation mapped
	Date        time.Time `json:"date"`
	Hours       float64   `gorm:"type:numeric(5,2);default:0" json:"hours"`
	Description string    `gorm:"type:varchar(255)" json:"description"`

	// 2. Billable Timesheet & 7. Project Profitability
	IsBillable bool    `gorm:"default:false" json:"is_billable"`         // Apakah jam ini akan ditagih ke klien?
	Cost       float64 `gorm:"type:numeric(15,2);default:0" json:"cost"` // Harga Pokok Tenaga Kerja
	CompanyID  uint    `gorm:"default:6" json:"company_id"`

	// Fase 2: Workflow Approval & Validation Gates
	Status          string     `gorm:"type:varchar(20);default:'approved'" json:"status"` // draft, submitted, approved, rejected, invoiced
	ApprovedByID    *uint      `json:"approved_by_id,omitempty"`
	ApprovedAt      *time.Time `json:"approved_at,omitempty"`
	RejectionReason string     `gorm:"type:text" json:"rejection_reason,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}


func (Timesheet) TableName() string {
	return "services.timesheets"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Timesheet{})
}
