package employees

import (
	"fmt"
	"ERP-System/app/modules/auth"
	"ERP-System/config"
	"time"
	"gorm.io/gorm"
)

type Department struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	Name      string      `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID *uint       `json:"manager_id"`
	Manager   *Employee   `json:"manager,omitempty"`
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
	EmergencyPhone   string         `gorm:"type:varchar(50)" json:"emergency_phone"`
	PTKPStatus       string         `gorm:"type:varchar(20);default:'TK/0'" json:"ptkp_status"` // TK/0, TK/1, K/0, K/1, K/2, K/3
	JoinDate         *time.Time     `json:"join_date"`
	IsActive         bool           `gorm:"default:true" json:"is_active"`
	CreatedAt        time.Time      `json:"created_at"`
	DeletedAt        *time.Time     `gorm:"index" json:"deleted_at,omitempty"`
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


type WarningLetter struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Employee    *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	WarningType string     `gorm:"type:varchar(50);not null" json:"warning_type"`
	IssueDate   time.Time  `json:"issue_date"`
	ExpiryDate  *time.Time `json:"expiry_date"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
}

type EmployeeTask struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	EmployeeID uint       `json:"employee_id"`
	Employee   *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	TaskName   string     `gorm:"type:varchar(255);not null" json:"task_name"`
	Type       string     `gorm:"type:varchar(50);default:'onboarding'" json:"type"`
	Status     string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
}

func (WarningLetter) TableName() string {
	return "hrd.warning_letters"
}
func (EmployeeTask) TableName() string {
	return "hrd.employee_tasks"
}


// --- 4. Transactions & Finance ---
type Overtime struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Employee    *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Date        time.Time  `json:"date"`
	Hours       float64    `json:"hours"`
	Description string     `gorm:"type:text" json:"description"`
	Status      string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
}

type EmployeeLoan struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	EmployeeID         uint       `json:"employee_id"`
	Employee           *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	PrincipalAmount    float64    `gorm:"type:numeric(15,2)" json:"principal_amount"`
	TenorMonths        int        `json:"tenor_months"`
	MonthlyInstallment float64    `gorm:"type:numeric(15,2)" json:"monthly_installment"`
	Status             string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	Date               time.Time  `json:"date"`
	CreatedAt          time.Time  `json:"created_at"`
}

type Expense struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	EmployeeID  uint       `json:"employee_id"`
	Employee    *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	ExpenseType string     `gorm:"type:varchar(100)" json:"expense_type"`
	Amount      float64    `gorm:"type:numeric(15,2)" json:"amount"`
	Description string     `gorm:"type:text" json:"description"`
	Status      string     `gorm:"type:varchar(50);default:'pending'" json:"status"`
	Date        time.Time  `json:"date"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (Overtime) TableName() string {
	return "hrd.overtimes"
}
func (EmployeeLoan) TableName() string {
	return "hrd.employee_loans"
}
func (Expense) TableName() string {
	return "hrd.expenses"
}


// --- 5. Payroll ---
type Payslip struct {
	ID              uint          `gorm:"primaryKey" json:"id"`
	EmployeeID      uint          `json:"employee_id"`
	Employee        *Employee     `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Period          string        `gorm:"type:varchar(20);not null" json:"period"`
	BasicSalary     float64       `gorm:"type:numeric(15,2)" json:"basic_salary"`
	TotalEarning    float64       `gorm:"type:numeric(15,2)" json:"total_earning"`
	TotalDeduction  float64       `gorm:"type:numeric(15,2)" json:"total_deduction"`
	NetSalary       float64       `gorm:"type:numeric(15,2)" json:"net_salary"`
	Status          string        `gorm:"type:varchar(50);default:'draft'" json:"status"`
	PayslipLines    []PayslipLine `gorm:"foreignKey:PayslipID" json:"payslip_lines,omitempty"`
	CreatedAt       time.Time     `json:"created_at"`
}

type PayslipLine struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PayslipID uint      `json:"payslip_id"`
	Category  string    `gorm:"type:varchar(50)" json:"category"` // earning, deduction
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Amount    float64   `gorm:"type:numeric(15,2)" json:"amount"`
}

func (Payslip) TableName() string { return "hrd.hr_payslips" }
func (PayslipLine) TableName() string { return "hrd.hr_payslip_lines" }

// --- 5b. Tunjangan Hari Raya (THR) ---
type EmployeeTHR struct {
	ID              uint       `gorm:"primaryKey" json:"id"`
	EmployeeID      uint       `json:"employee_id"`
	Employee        *Employee  `gorm:"foreignKey:EmployeeID" json:"employee,omitempty"`
	Year            int        `json:"year"`
	CutoffDate      time.Time  `json:"cutoff_date"`
	JoinDate        time.Time  `json:"join_date"`
	TenureMonths    int        `json:"tenure_months"`
	BasicWage       float64    `gorm:"type:numeric(15,2)" json:"basic_wage"`
	THRAmount       float64    `gorm:"type:numeric(15,2)" json:"thr_amount"`
	CalculationType string     `gorm:"type:varchar(50)" json:"calculation_type"` // Full (1 Month), Prorate (N/12)
	Status          string     `gorm:"type:varchar(50);default:'draft'" json:"status"` // draft, approved, paid
	CreatedAt       time.Time  `json:"created_at"`
}

func (EmployeeTHR) TableName() string { return "hrd.employee_thrs" }


// --- 6. Audit & System ---
type AuditLog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	TargetTable string    `gorm:"type:varchar(100)" json:"target_table"`
	Action    string    `gorm:"type:varchar(50)" json:"action"` // UPDATE, DELETE
	RecordID  uint      `json:"record_id"`
	OldData   string    `gorm:"type:text" json:"old_data"`
	NewData   string    `gorm:"type:text" json:"new_data"`
	UserID    uint      `json:"user_id"` // Admin who did it
	CreatedAt time.Time `json:"created_at"`
}

func (AuditLog) TableName() string { return "hrd.audit_logs" }

// Hooks
func (c *Contract) AfterUpdate(tx *gorm.DB) (err error) {
	// Simple hook to log changes (In real app, we diff old vs new)
	log := AuditLog{
		TargetTable: "hrd.contracts",
		Action:    "UPDATE",
		RecordID:  c.ID,
		OldData:   "Previous Wage (tracked by DB)",
		NewData:   fmt.Sprintf("New Wage: %f", c.Wage),
		UserID:    1, // Default Admin
	}
	tx.Create(&log)
	return
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
		&WarningLetter{}, &EmployeeTask{},
		&Overtime{}, &EmployeeLoan{}, &Expense{},
		&Payslip{}, &PayslipLine{}, &EmployeeTHR{}, &AuditLog{},
	)
}
