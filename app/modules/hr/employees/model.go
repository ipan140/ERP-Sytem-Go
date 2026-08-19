package employees

import (
	"ERP-System/config"
	"time"
)

// --- 1. Employees & Organization ---

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID *uint     `json:"manager_id"` // Pointer to allow null (Employee ID)
	ParentID  *uint     `json:"parent_id"`  // Pointer to allow null (Self-referential)
	CreatedAt time.Time `json:"created_at"`
}

type JobPosition struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. Software Engineer
	DepartmentID *uint     `json:"department_id"`
	State        string    `gorm:"type:varchar(50);default:'recruit'" json:"state"` // recruit, open
	CreatedAt    time.Time `json:"created_at"`
}

type Employee struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	Name             string    `gorm:"type:varchar(255);not null" json:"name"`
	UserID           *uint     `json:"user_id"`       // Links to auth users
	DepartmentID     *uint     `json:"department_id"` // Links to Department
	JobPositionID    *uint     `json:"job_position_id"`
	ManagerID        *uint     `json:"manager_id"` // Self-referential
	WorkEmail        string    `gorm:"type:varchar(100)" json:"work_email"`
	WorkPhone        string    `gorm:"type:varchar(50)" json:"work_phone"`
	EmergencyContact string    `gorm:"type:varchar(100)" json:"emergency_contact"`
	EmergencyPhone   string    `gorm:"type:varchar(50)" json:"emergency_phone"`
	CreatedAt        time.Time `json:"created_at"`
}

// --- 2. Contracts ---

type WorkingSchedule struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"` // e.g. Senin - Jumat, 40 Jam/Minggu
	HoursPerWeek float64   `json:"hours_per_week"`
	CreatedAt    time.Time `json:"created_at"`
}

type Contract struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	EmployeeID        uint      `json:"employee_id"`
	JobPositionID     uint      `json:"job_position_id"`
	Wage              float64   `gorm:"type:numeric(15,2);default:0" json:"wage"` // Basic Salary
	StartDate         time.Time `json:"start_date"`
	EndDate           *time.Time`json:"end_date"`
	WorkingScheduleID *uint     `json:"working_schedule_id"`
	State             string    `gorm:"type:varchar(50);default:'draft'" json:"state"` // draft, probation, open, close
	CreatedAt         time.Time `json:"created_at"`
}

// --- 3. Skills & Resumes ---

type Skill struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"` // e.g. Golang
}

type SkillLevel struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	SkillID uint   `json:"skill_id"`
	Name    string `gorm:"type:varchar(50);not null" json:"name"` // e.g. Expert, Beginner
}

type EmployeeSkill struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	EmployeeID   uint `json:"employee_id"`
	SkillID      uint `json:"skill_id"`
	SkillLevelID uint `json:"skill_level_id"`
}

type ResumeLine struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EmployeeID  uint      `json:"employee_id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Type        string    `gorm:"type:varchar(50)" json:"type"` // experience, education
	DateStart   time.Time `json:"date_start"`
	DateEnd     *time.Time`json:"date_end"`
	Description string    `gorm:"type:text" json:"description"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, 
		&Department{}, &JobPosition{}, &Employee{},
		&WorkingSchedule{}, &Contract{},
		&Skill{}, &SkillLevel{}, &EmployeeSkill{}, &ResumeLine{},
	)
}
