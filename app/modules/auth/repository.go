package auth

import (
	"ERP-System/config"
	"errors"

	"gorm.io/gorm"
)

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := config.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Data tidak ditemukan
		}
		return nil, err
	}
	return &user, nil
}
