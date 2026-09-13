package user_roles

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ERP-System/config"
	redisPkg "ERP-System/pkg/redis"
)

var ctx = context.Background()

func GetAllUsersRolesService() ([]UserRoleResponse, error) {
	// [Redis] 1. Coba ambil dari Cache dulu
	cacheKey := "core:user_roles:all"
	
	if redisPkg.Client != nil {
		redisRepo := NewUserRolesRedis(redisPkg.Client)
		cachedData, err := redisRepo.GetUserRole(ctx, cacheKey)
		if err == nil && cachedData != "" {
			var users []UserRoleResponse
			_ = json.Unmarshal([]byte(cachedData), &users)
			fmt.Println("🚀 [Redis Hit] Fetching User Roles from Cache!")
			return users, nil
		}
	}

	// [DB] 2. Jika di Cache tidak ada (Miss), ambil dari PostgreSQL
	fmt.Println("🐢 [DB Hit] Fetching User Roles from PostgreSQL...")
	var users []UserRoleResponse
	err := config.DB.Table("setting.users").Select("id, name, email, role as roles").Find(&users).Error
	if err != nil || len(users) == 0 {
		err = config.DB.Table("users").Select("id, name, email, roles").Find(&users).Error
	}
	
	if err == nil && redisPkg.Client != nil {
		// [Redis] 3. Simpan hasil query DB ke Redis agar request selanjutnya cepat
		bytes, _ := json.Marshal(users)
		_ = redisPkg.Client.Set(ctx, cacheKey, bytes, 1*time.Hour).Err()
	}

	return users, err
}

func AssignRoleService(userID uint, roles string) error {
	err := config.DB.Table("setting.users").Where("id = ?", userID).Update("role", roles).Error
	if err != nil {
		err = config.DB.Table("users").Where("id = ?", userID).Update("roles", roles).Error
	}
	
	if err == nil && redisPkg.Client != nil {
		// [Redis] Hapus cache (Invalidation) karena ada perubahan role
		_ = redisPkg.Client.Del(ctx, "core:user_roles:all").Err()
		fmt.Println("🧹 [Redis Clear] Invalidated 'core:user_roles:all' due to role assignment")
	}
	
	return err
}

// CRUD Roles Dinamis

func CreateRoleService(req Role) error {
	return config.DB.Create(&req).Error
}

func GetAllRolesService() ([]Role, error) {
	var roles []Role
	err := config.DB.Find(&roles).Error
	return roles, err
}

func UpdateRoleService(id uint, req Role) error {
	var role Role
	if err := config.DB.First(&role, id).Error; err != nil {
		return err
	}
	role.Name = req.Name
	role.Description = req.Description
	return config.DB.Save(&role).Error
}

func DeleteRoleService(id uint) error {
	return config.DB.Delete(&Role{}, id).Error
}

