package time_off

import (
	"ERP-System/config"
	"time"
)

type LeaveType struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. Sick, Annual
	RequiresApproval bool      `gorm:"default:true" json:"requires_approval"`
	CreatedAt        time.Time `json:"created_at"`
}

type LeaveAllocation struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	EmployeeID   uint      `json:"employee_id"`
	LeaveTypeID  uint      `json:"leave_type_id"`
	NumberOfDays float64   `gorm:"type:numeric(10,2);default:0" json:"number_of_days"`
	State        string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, approved
	CreatedAt    time.Time `json:"created_at"`
}

type LeaveRequest struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	EmployeeID       uint      `json:"employee_id"`
	LeaveTypeID      uint      `json:"leave_type_id"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	NumberOfDays     float64   `gorm:"type:numeric(10,2);default:0" json:"number_of_days"`
	ManagerApproveID *uint     `json:"manager_approve_id"`
	HRApproveID      *uint     `json:"hr_approve_id"`
	Status           string    `gorm:"type:varchar(20);default:'draft'" json:"status"` // draft, approved, rejected
	CreatedAt        time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &LeaveType{}, &LeaveAllocation{}, &LeaveRequest{})
}
