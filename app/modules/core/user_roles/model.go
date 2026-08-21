package user_roles

type AssignRoleRequest struct {
	UserID uint     `json:"user_id" validate:"required"`
	Roles  []string `json:"roles" validate:"required"`
}

type UserRoleResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Roles string `json:"roles"`
}
