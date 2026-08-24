package appraisals

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Appraisal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	ManagerID  uint      `json:"manager_id"`
	Manager *Appraisal `gorm:"foreignKey:ManagerID"` // Auto-added relation
	Date       time.Time `json:"date"`
	Score      float64   `gorm:"type:numeric(5,2);default:0" json:"score"` // e.g. 1 to 5
	Feedback   string    `gorm:"type:text" json:"feedback"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, pending, done
	CreatedAt  time.Time `json:"created_at"`
}


func (Appraisal) TableName() string {
	return "hrd.appraisals"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Appraisal{})
}
