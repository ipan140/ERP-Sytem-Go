package user_roles

import (
	"ERP-System/app/modules/auth"
	"ERP-System/config"
)

type AssignRoleRequest struct {
	UserID uint       `json:"user_id" validate:"required"`
	User   *auth.User `gorm:"foreignKey:UserID" json:"user,omitempty"` // Cross-module relation
	Roles  []string   `json:"roles" validate:"required"`
}

type UserRoleResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}

func (AssignRoleRequest) TableName() string {
	return "setting.assign_role_requests"
}

func (UserRoleResponse) TableName() string {
	return "setting.user_role_responses"
}

// Model baru untuk CRUD Role Dinamis
type Role struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"type:varchar(50);unique;not null" json:"name"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

func (Role) TableName() string {
	return "setting.roles"
}

func init() {
	config.ModelsToMigrate = append(config.ModelsToMigrate, &Role{})
}
