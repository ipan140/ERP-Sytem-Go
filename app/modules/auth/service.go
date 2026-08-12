package auth

import (
	"ERP-System/common/errors"
	"ERP-System/common/utils"
)

func LoginService(email, password string) (string, *errors.AppError) {
	user, err := FindUserByEmail(email)
	if err != nil {
		return "", errors.NewInternalServer("Terjadi kesalahan pada database")
	}

	if user == nil {
		return "", errors.NewUnauthorized("Email atau password salah")
	}

	// Validasi hash password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.NewUnauthorized("Email atau password salah")
	}

	// Generate Token yang mengandung Role
	token, err := utils.GenerateToken(user.ID, user.CompanyID, user.Role)
	if err != nil {
		return "", errors.NewInternalServer("Gagal membuat token autentikasi")
	}

	return token, nil
}
