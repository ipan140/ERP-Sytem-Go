package timesheets

import (
	"ERP-System/config"
	"time"
)

type Timesheet struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"` // Who worked
	ProjectID   uint      `json:"project_id"`
	TaskID      *uint     `json:"task_id"` // Optional
	Date        time.Time `json:"date"`
	Hours       float64   `gorm:"type:numeric(5,2);default:0" json:"hours"`
	Description string    `gorm:"type:varchar(255)" json:"description"`

	// 2. Billable Timesheet & 7. Project Profitability
	IsBillable bool    `gorm:"default:false" json:"is_billable"`         // Apakah jam ini akan ditagih ke klien?
	Cost       float64 `gorm:"type:numeric(15,2);default:0" json:"cost"` // Harga Pokok Tenaga Kerja

	CreatedAt time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Timesheet{})
}
