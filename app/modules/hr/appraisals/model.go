package appraisals

import (
	"ERP-System/config"
	"time"
)

type Appraisal struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `json:"employee_id"`
	ManagerID  uint      `json:"manager_id"`
	Date       time.Time `json:"date"`
	Score      float64   `gorm:"type:numeric(5,2);default:0" json:"score"` // e.g. 1 to 5
	Feedback   string    `gorm:"type:text" json:"feedback"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, pending, done
	CreatedAt  time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Appraisal{})
}
