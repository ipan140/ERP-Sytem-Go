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

func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

func CreateCompany(company *Company) error {
	return config.DB.Create(company).Error
}

func GetFirstCompany() (*Company, error) {
	var company Company
	err := config.DB.First(&company).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &company, nil
}
