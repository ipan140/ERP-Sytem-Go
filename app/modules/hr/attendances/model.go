package attendances

import (
	"ERP-System/config"
	"time"
)

type Attendance struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	CheckIn     time.Time  `json:"check_in"`
	CheckOut    *time.Time `json:"check_out"`
	WorkedHours float64    `gorm:"type:numeric(15,2);default:0" json:"worked_hours"`
	CreatedAt   time.Time  `json:"created_at"`
}

type Overtime struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `json:"employee_id"`
	Date       time.Time `json:"date"`
	Hours      float64   `gorm:"type:numeric(15,2);default:0" json:"hours"`
	State      string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, approved
	CreatedAt  time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Attendance{}, &Overtime{})
}
