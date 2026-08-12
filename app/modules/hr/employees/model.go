package employees

import (
	"ERP-System/config"
	"time"
)

type Department struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	ManagerID *uint     `json:"manager_id"` // Pointer to allow null (Employee ID)
	ParentID  *uint     `json:"parent_id"`  // Pointer to allow null (Self-referential)
	CreatedAt time.Time `json:"created_at"`
}

type Employee struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"type:varchar(255);not null" json:"name"`
	UserID       *uint     `json:"user_id"`       // Links to auth users
	DepartmentID *uint     `json:"department_id"` // Links to Department
	JobTitle     string    `gorm:"type:varchar(100)" json:"job_title"`
	ManagerID    *uint     `json:"manager_id"` // Self-referential
	WorkEmail    string    `gorm:"type:varchar(100)" json:"work_email"`
	WorkPhone    string    `gorm:"type:varchar(50)" json:"work_phone"`
	CreatedAt    time.Time `json:"created_at"`
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Department{}, &Employee{})
}
