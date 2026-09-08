package planning

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Shift struct {
	ID         uint                `gorm:"primaryKey" json:"id"`
	Name       string              `gorm:"type:varchar(255)" json:"name"`
	EmployeeID *uint               `json:"employee_id"`
	Employee   *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Role       string              `gorm:"type:varchar(100)" json:"role"`
	Date       string              `gorm:"type:varchar(50)" json:"date"`
	StartTime  string              `gorm:"type:varchar(50)" json:"start_time"`
	EndTime    string              `gorm:"type:varchar(50)" json:"end_time"`
	Hours      float64             `gorm:"type:numeric(5,2);default:0" json:"hours"`
	State      string              `gorm:"type:varchar(50);default:'draft'" json:"state"`
	Notes      string              `gorm:"type:text" json:"notes"`
	CreatedAt  time.Time           `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Shift{})
}


// ---- Auto-Generated TableName methods ----
func (Shift) TableName() string {
	return "services.shifts"
}