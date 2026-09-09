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
	CompanyID     uint                `gorm:"default:6" json:"company_id"`

	// Fase 2: e-BAST Validation Gate
	BastValidated bool                `gorm:"default:false" json:"bast_validated"`
	ValidatedByID *uint               `json:"validated_by_id,omitempty"`
	ValidatedAt   *time.Time          `json:"validated_at,omitempty"`

	// Fase 4: GPS Geotagging
	CheckInLat    *float64            `gorm:"type:numeric(10,7)" json:"check_in_lat,omitempty"`
	CheckInLng    *float64            `gorm:"type:numeric(10,7)" json:"check_in_lng,omitempty"`
	CheckInAt     *time.Time          `json:"check_in_at,omitempty"`
	CheckOutLat   *float64            `gorm:"type:numeric(10,7)" json:"check_out_lat,omitempty"`
	CheckOutLng   *float64            `gorm:"type:numeric(10,7)" json:"check_out_lng,omitempty"`
	CheckOutAt    *time.Time          `json:"check_out_at,omitempty"`

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