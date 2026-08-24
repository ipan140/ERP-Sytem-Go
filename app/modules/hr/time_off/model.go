package time_off

import (
	"ERP-System/app/modules/hr/employees"
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
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	LeaveTypeID  uint      `json:"leave_type_id"`
	LeaveType *LeaveType `gorm:"foreignKey:LeaveTypeID"` // Auto-added relation
	NumberOfDays float64   `gorm:"type:numeric(10,2);default:0" json:"number_of_days"`
	State        string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, approved
	CreatedAt    time.Time `json:"created_at"`
}

type LeaveRequest struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	EmployeeID       uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	LeaveTypeID      uint      `json:"leave_type_id"`
	LeaveType *LeaveType `gorm:"foreignKey:LeaveTypeID"` // Auto-added relation
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	NumberOfDays     float64   `gorm:"type:numeric(10,2);default:0" json:"number_of_days"`
	ManagerApproveID *uint     `json:"manager_approve_id"`
	ManagerApprove *employees.Employee `gorm:"foreignKey:ManagerApproveID" json:"managerapprove,omitempty"` // Odoo relation mapped
	HRApproveID      *uint     `json:"hr_approve_id"`
	HRApprove *employees.Employee `gorm:"foreignKey:HRApproveID" json:"hrapprove,omitempty"` // Odoo relation mapped
	Status           string    `gorm:"type:varchar(20);default:'draft'" json:"status"` // draft, approved, rejected
	CreatedAt        time.Time `json:"created_at"`
}


func (LeaveType) TableName() string {
	return "hrd.leave_types"
}

func (LeaveAllocation) TableName() string {
	return "hrd.leave_allocations"
}

func (LeaveRequest) TableName() string {
	return "hrd.leave_requests"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &LeaveType{}, &LeaveAllocation{}, &LeaveRequest{})
}
