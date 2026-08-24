package user_roles

import "ERP-System/app/modules/auth"

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
