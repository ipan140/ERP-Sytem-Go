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

	CreatedAt time.Time `json:"created_at"`
}


func (Timesheet) TableName() string {
	return "services.timesheets"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Timesheet{})
}
