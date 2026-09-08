package field_service

import (
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type ChecklistItem struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type FieldServiceTask struct {
	ID            uint                `gorm:"primaryKey" json:"id"`
	Name          string              `gorm:"type:varchar(255);not null" json:"name"`
	PartnerID     *uint               `json:"partner_id"`
	Partner       *base.Partner       `gorm:"foreignKey:PartnerID" json:"partner,omitempty"`
	EmployeeID    *uint               `json:"employee_id"`
	Employee      *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	ScheduledDate *time.Time          `json:"scheduled_date,omitempty"`
	Address       string              `gorm:"type:text" json:"address"`
	Priority      string              `gorm:"type:varchar(50);default:'low'" json:"priority"`
	State         string              `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, assigned, in_progress, completed, cancelled
	Notes         string              `gorm:"type:text" json:"notes"`
	Signature     string              `gorm:"type:text" json:"signature"`
	Checklist     []ChecklistItem     `gorm:"serializer:json" json:"checklist"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &FieldServiceTask{})
}

// ---- Auto-Generated TableName methods ----
func (FieldServiceTask) TableName() string {
	return "services.field_service_tasks"
}