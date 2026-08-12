package time_off

import (
	"ERP-System/config"
	"time"
)

type LeaveType struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(100);not null" json:"name"`
	RequiresApproval bool      `gorm:"default:true" json:"requires_approval"`
	CreatedAt        time.Time `json:"created_at"`
}

type LeaveRequest struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	LeaveTypeID uint      `json:"leave_type_id"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	Status      string    `gorm:"type:varchar(20);default:'draft'" json:"status"` // draft, approved, rejected
	CreatedAt   time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &LeaveType{}, &LeaveRequest{})
}
