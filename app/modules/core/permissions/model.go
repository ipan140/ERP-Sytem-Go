package permissions

// RolePermission mewakili satu baris toggle di database
type RolePermission struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	RoleName  string `json:"role_name" validate:"required"` // Contoh: "SALES_MANAGER"
	Module    string `json:"module" validate:"required"`    // Contoh: "inventory" atau "sales"
	CanRead   bool   `json:"can_read"`                      // Toggle Lihat
	CanWrite  bool   `json:"can_write"`                     // Toggle Buat/Edit
	CanDelete bool   `json:"can_delete"`                    // Toggle Hapus
}

type TogglePermissionRequest struct {
	RoleName string `json:"role_name" validate:"required"`
	Module   string `json:"module" validate:"required"`
	Action   string `json:"action" validate:"required"` // "read", "write", "delete"
	Value    bool   `json:"value"`                      // true (On) / false (Off)
}
