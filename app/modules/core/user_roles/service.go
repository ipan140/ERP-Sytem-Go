package user_roles

import (
	"ERP-System/config"
)

func GetAllUsersRolesService() ([]UserRoleResponse, error) {
	var users []UserRoleResponse
	// Menggunakan Raw SQL via GORM agar tidak perlu memanggil struct User dari modul auth
	err := config.DB.Table("users").Select("id, name, email, roles").Find(&users).Error
	return users, err
}

func AssignRoleService(userID uint, roles string) error {
	// Update langsung ke kolom 'roles' di tabel users
	err := config.DB.Table("users").Where("id = ?", userID).Update("roles", roles).Error
	return err
}
