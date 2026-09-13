package recruitment

import (
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Stage struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(100);not null" json:"name"` // Initial Qualification, Interview, Offered
	Sequence int    `gorm:"default:10" json:"sequence"`
}

type Applicant struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"type:varchar(255);not null" json:"name"`
	Email          string    `gorm:"type:varchar(100)" json:"email"`
	Phone          string    `gorm:"type:varchar(50)" json:"phone"`
	JobPositionID  uint      `json:"job_position_id"`
	JobPosition *employees.JobPosition `gorm:"foreignKey:JobPositionID" json:"jobposition,omitempty"` // Odoo relation mapped
	StageID        uint      `json:"stage_id"`
	Stage *Stage `gorm:"foreignKey:StageID"` // Auto-added relation
	ExpectedSalary float64   `gorm:"type:numeric(15,2);default:0" json:"expected_salary"`
	ResumeURL      string    `gorm:"type:varchar(500)" json:"resume_url"`
	Notes          string    `gorm:"type:text" json:"notes"`
	State          string    `gorm:"type:varchar(50);default:'in_progress'" json:"state"` // in_progress, hired, refused
	CreatedAt      time.Time `json:"created_at"`
}


func (Stage) TableName() string {
	return "hrd.stages"
}

func (Applicant) TableName() string {
	return "hrd.applicants"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Stage{}, &Applicant{})
}
