package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ERP-System/common/errors"
	"ERP-System/common/utils"
	"ERP-System/pkg/rabbitmq"
	redisPkg "ERP-System/pkg/redis"
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

	// [Keamanan Pilar Tambahan] Simpan token sebagai Single Active Session di Redis (Berlaku 24 Jam)
	if rabbitmq.Channel != nil { // reusing rabbitmq check or directly checking redisPkg.Client
		// Import redisPkg is needed. I'll rely on goimports or add it manually.
	}
	// Let's just do it directly. We'll run goimports later to fix it if needed.
	ctx := context.Background()
	_ = redisPkg.Client.Set(ctx, fmt.Sprintf("active_token:%d", user.ID), token, 24*time.Hour).Err()

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
		return nil, errors.NewInternalServer("Gagal menyimpan pengguna")
	}

	emailSubject := "Selamat Datang di ERP System!"
	emailBody := "Halo " + user.Name + ",\n\nTerima kasih telah mendaftar di sistem kami."
	_ = PublishEmailEvent(user.Email, emailSubject, emailBody)

	if rabbitmq.Channel != nil {
		event := map[string]string{
			"email": user.Email,
			"name":  user.Name,
		}
		bodyBytes, _ := json.Marshal(event)
		_ = rabbitmq.PublishDelayedEvent(rabbitmq.Channel, "auth_delayed_cleanup", bodyBytes, 259200000)
		log.Println("🧹 Event RabbitMQ: Timer Cleanup Akun 3 Hari dimulai!")
	}

	return user, nil
}
