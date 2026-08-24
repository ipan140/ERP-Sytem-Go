package appointments

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Appointment struct {
	ID         uint                `gorm:"primaryKey" json:"id"`
	Name       string              `gorm:"type:varchar(255);not null" json:"name"`
	EmployeeID uint                `json:"employee_id"`
	Employee   *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	PartnerID  uint                `json:"partner_id"`
	Partner    *base.Partner       `gorm:"foreignKey:PartnerID" json:"partner,omitempty"`
	Date       time.Time           `json:"date"`
	Notes      string              `gorm:"type:text" json:"notes"`
	State      string              `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, confirmed, done
	CreatedAt  time.Time           `json:"created_at"`
}

func (Appointment) TableName() string {
	return "services.appointments"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Appointment{})
}
