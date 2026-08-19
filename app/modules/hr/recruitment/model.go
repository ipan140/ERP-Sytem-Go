package recruitment

import (
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
	StageID        uint      `json:"stage_id"`
	ExpectedSalary float64   `gorm:"type:numeric(15,2);default:0" json:"expected_salary"`
	State          string    `gorm:"type:varchar(50);default:'in_progress'" json:"state"` // in_progress, hired, refused
	CreatedAt      time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Stage{}, &Applicant{})
}
