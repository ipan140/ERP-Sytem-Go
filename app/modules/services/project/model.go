package project

import (

	"ERP-System/app/modules/auth"
	"ERP-System/app/modules/core/base"
	"ERP-System/app/modules/hr/employees"
	"ERP-System/config"
	"time"
)

type Project struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	ManagerID   uint                `json:"manager_id"` // Employee ID
	Manager     *employees.Employee `json:"manager,omitempty"` // Belongs To Employee
	CustomerID  uint      `json:"customer_id"`                                    // Partner ID
	Customer *base.Partner `gorm:"foreignKey:CustomerID" json:"customer,omitempty"` // Cross-module relation
	SaleOrderID *uint     `json:"sale_order_id"`                                  // 1. Integrasi Sales (Auto-Create)

	State       string    `gorm:"type:varchar(50);default:'active'" json:"state"` // active, done, cancelled
	CreatedAt   time.Time `json:"created_at"`
}

type Task struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	ProjectID         uint      `json:"project_id"`
	Project *Project `gorm:"foreignKey:ProjectID"` // Auto-added relation
	Name              string    `gorm:"type:varchar(255);not null" json:"name"`
	AssigneeID        uint      `json:"assignee_id"`       // Employee ID
	Assignee *auth.User `gorm:"foreignKey:AssigneeID" json:"assignee,omitempty"` // Odoo relation mapped
	RequiredSkillID   *uint     `json:"required_skill_id"` // 6. Skill Routing Teknisi
	RequiredSkill *employees.Skill `gorm:"foreignKey:RequiredSkillID" json:"requiredskill,omitempty"` // Odoo relation mapped
	Deadline          time.Time `json:"deadline"`
	Stage             string    `gorm:"type:varchar(50);default:'todo'" json:"stage"` // todo, in_progress, done
	CustomerSignature string    `gorm:"type:text" json:"customer_signature"`          // 5. Digital Signature Lapangan
	CreatedAt         time.Time `json:"created_at"`
}

// 8. Milestone Billing
type ProjectMilestone struct {
	ID                uint    `gorm:"primaryKey" json:"id"`
	ProjectID         uint    `json:"project_id"`
	Project *Project `gorm:"foreignKey:ProjectID"` // Auto-added relation
	Name              string  `gorm:"type:varchar(100);not null" json:"name"` // e.g., Pondasi Selesai
	IsReached         bool    `gorm:"default:false" json:"is_reached"`
	InvoicePercentage float64 `gorm:"type:numeric(5,2)" json:"invoice_percentage"` // Berapa % yang ditagihkan saat ini selesai
}

// 3. Task Dependencies (Gantt Chart)
type TaskDependency struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	TaskID       uint `json:"task_id"`        // Tugas ini
	Task *Task `gorm:"foreignKey:TaskID"` // Auto-added relation
	BlocksTaskID uint `json:"blocks_task_id"` // Menghalangi tugas ini
	BlocksTask *Task `gorm:"foreignKey:BlocksTaskID" json:"blockstask,omitempty"` // Odoo relation mapped
}

// 9. Resource Forecasting
type ResourceForecast struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	EmployeeID   uint      `json:"employee_id"`
	Employee *employees.Employee `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"` // Cross-module relation
	ProjectID    uint      `json:"project_id"`
	Project *Project `gorm:"foreignKey:ProjectID"` // Auto-added relation
	HoursPerWeek float64   `gorm:"type:numeric(5,2)" json:"hours_per_week"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
}


func (Project) TableName() string {
	return "services.projects"
}

func (Task) TableName() string {
	return "services.tasks"
}

func (ProjectMilestone) TableName() string {
	return "services.project_milestones"
}

func (TaskDependency) TableName() string {
	return "services.task_dependencies"
}

func (ResourceForecast) TableName() string {
	return "services.resource_forecasts"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Project{}, &Task{}, &ProjectMilestone{}, &TaskDependency{}, &ResourceForecast{})
}
