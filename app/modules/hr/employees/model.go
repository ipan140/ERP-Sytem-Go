package employees

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"time"
)

type Department struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Name      string      `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID *uint       `json:"manager_id"`
	Manager   *Employee   `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	ParentID  *uint       `json:"parent_id"`
	Parent    *Department `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	CreatedAt time.Time   `json:"created_at"`
}

type JobPosition struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	Name         string      `gorm:"type:varchar(100);not null" json:"name"`
	DepartmentID *uint       `json:"department_id"`
	Department   *Department `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	State        string      `gorm:"type:varchar(50);default:'recruit'" json:"state"`
	CreatedAt    time.Time   `json:"created_at"`
}

type Employee struct {
	ID               uint         `gorm:"primaryKey" json:"id"`
	Name             string       `gorm:"type:varchar(255);not null" json:"name"`
	UserID           *uint        `json:"user_id"`
	User             *auth.User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	DepartmentID     *uint        `json:"department_id"`
	Department       *Department  `gorm:"foreignKey:DepartmentID" json:"department,omitempty"`
	JobPositionID    *uint        `json:"job_position_id"`
	JobPosition      *JobPosition `gorm:"foreignKey:JobPositionID" json:"job_position,omitempty"`
	ManagerID        *uint        `json:"manager_id"`
	Manager          *Employee    `gorm:"foreignKey:ManagerID" json:"manager,omitempty"`
	WorkEmail        string       `gorm:"type:varchar(100)" json:"work_email"`
	WorkPhone        string       `gorm:"type:varchar(50)" json:"work_phone"`
	EmergencyContact string       `gorm:"type:varchar(100)" json:"emergency_contact"`
	EmergencyPhone   string       `gorm:"type:varchar(50)" json:"emergency_phone"`
	CreatedAt        time.Time    `json:"created_at"`
}

// --- 2. Contracts ---

type WorkingSchedule struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(100);not null" json:"name"`
	HoursPerWeek float64   `json:"hours_per_week"`
	CreatedAt    time.Time `json:"created_at"`
}

type Contract struct {
	ID                uint             `gorm:"primaryKey" json:"id"`
	EmployeeID        uint             `json:"employee_id"`
	Employee          *Employee        `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	JobPositionID     uint             `json:"job_position_id"`
	JobPosition       *JobPosition     `gorm:"foreignKey:JobPositionID" json:"job_position,omitempty"`
	Wage              float64          `gorm:"type:numeric(15,2);default:0" json:"wage"`
	StartDate         time.Time        `json:"start_date"`
	EndDate           *time.Time       `json:"end_date"`
	WorkingScheduleID *uint            `json:"working_schedule_id"`
	WorkingSchedule   *WorkingSchedule `gorm:"foreignKey:WorkingScheduleID" json:"working_schedule,omitempty"`
	State             string           `gorm:"type:varchar(50);default:'draft'" json:"state"`
	CreatedAt         time.Time        `json:"created_at"`
}

// --- 3. Skills & Resumes ---

type Skill struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}

type SkillLevel struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	SkillID uint   `json:"skill_id"`
	Skill   *Skill `gorm:"foreignKey:SkillID" json:"skill,omitempty"`
	Name    string `gorm:"type:varchar(50);not null" json:"name"`
}

type EmployeeSkill struct {
	ID           uint        `gorm:"primaryKey" json:"id"`
	EmployeeID   uint        `json:"employee_id"`
	Employee     *Employee   `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	SkillID      uint        `json:"skill_id"`
	Skill        *Skill      `gorm:"foreignKey:SkillID" json:"skill,omitempty"`
	SkillLevelID uint        `json:"skill_level_id"`
	SkillLevel   *SkillLevel `gorm:"foreignKey:SkillLevelID" json:"skill_level,omitempty"`
}

type ResumeLine struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Employee    *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Title       string     `gorm:"type:varchar(255);not null" json:"title"`
	Type        string     `gorm:"type:varchar(50)" json:"type"`
	DateStart   time.Time  `json:"date_start"`
	DateEnd     *time.Time `json:"date_end"`
	Description string     `gorm:"type:text" json:"description"`
}

func (Department) TableName() string {
	return "hrd.departments"
}
func (JobPosition) TableName() string {
	return "hrd.job_positions"
}
func (Employee) TableName() string {
	return "hrd.employees"
}
func (WorkingSchedule) TableName() string {
	return "hrd.working_schedules"
}
func (Contract) TableName() string {
	return "hrd.contracts"
}
func (Skill) TableName() string {
	return "hrd.skills"
}
func (SkillLevel) TableName() string {
	return "hrd.skill_levels"
}
func (EmployeeSkill) TableName() string {
	return "hrd.employee_skills"
}
func (ResumeLine) TableName() string {
	return "hrd.resume_lines"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate,
		&Department{}, &JobPosition{}, &Employee{},
		&WorkingSchedule{}, &Contract{},
		&Skill{}, &SkillLevel{}, &EmployeeSkill{}, &ResumeLine{},
	)
}
