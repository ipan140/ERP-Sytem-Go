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

func RegisterService(name, email, password string) (*User, *errors.AppError) {
	existing, _ := FindUserByEmail(email)
	if existing != nil {
		return nil, errors.NewBadRequest("Email sudah digunakan")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.NewInternalServer("Gagal memproses password")
	}

	// Cari company default, jika tidak ada, buat baru
	company, err := GetFirstCompany()
	if err != nil {
		return nil, errors.NewInternalServer("Kesalahan database")
	}

	if company == nil {
		company = &Company{Name: "My Company"}
		if err := CreateCompany(company); err != nil {
			return nil, errors.NewInternalServer("Gagal membuat company default")
		}
	}

	user := &User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      "staff",
		CompanyID: company.ID,
	}

	if err := CreateUser(user); err != nil {
		return nil, errors.NewInternalServer("Gagal mendaftarkan pengguna")
	}

	// [RabbitMQ] - Fase 2: Publish Event untuk Mengirim Welcome Email secara asinkron
	emailSubject := "Selamat Datang di ERP System!"
	emailBody := "Halo " + user.Name + ",\n\nTerima kasih telah mendaftar di sistem kami."
	_ = PublishEmailEvent(user.Email, emailSubject, emailBody)

	return user, nil
}
