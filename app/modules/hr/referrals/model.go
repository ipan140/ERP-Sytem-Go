package referrals

import (
	"ERP-System/config"
	"time"
)

type ReferralReward struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"` // e.g. iPhone 15 Pro
	Cost      int       `gorm:"default:0" json:"cost"`                  // Cost in points
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type ReferralPoint struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"`  // Who referred
	ApplicantID uint      `json:"applicant_id"` // Who was referred
	Points      int       `json:"points"`
	Reason      string    `gorm:"type:varchar(255)" json:"reason"` // e.g. Candidate Hired
	CreatedAt   time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &ReferralReward{}, &ReferralPoint{})
}
