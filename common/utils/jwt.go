package utils

import (
	"ERP-System/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID    uint   `json:"user_id"`
	CompanyID uint   `json:"company_id"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, companyID uint, role string) (string, error) {
	claims := &JWTCustomClaims{
		UserID:    userID,
		CompanyID: companyID,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.AppConfig.JWTSecret
	if secret == "" {
		secret = "default_secret" // fallback
	}

	return token.SignedString([]byte(secret))
}
