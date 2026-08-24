package lunch

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type LunchOrder struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EmployeeID uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	Date       time.Time `json:"date"`
	Total      float64   `gorm:"type:numeric(15,2);default:0" json:"total"`
	State      string    `gorm:"type:varchar(50);default:'new'" json:"state"` // new, confirmed, cancelled
	CreatedAt  time.Time `json:"created_at"`
}

type LunchCashmove struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	Date        time.Time `json:"date"`
	Amount      float64   `gorm:"type:numeric(15,2);not null" json:"amount"` // Positive for deposit/subsidy, Negative for payment
	Description string    `gorm:"type:varchar(255)" json:"description"`
}


func (LunchOrder) TableName() string {
	return "hrd.lunch_orders"
}

func (LunchCashmove) TableName() string {
	return "hrd.lunch_cashmoves"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &LunchOrder{}, &LunchCashmove{})
}
