package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"ERP-System/common/utils"
	"ERP-System/config"
	redisPkg "ERP-System/pkg/redis"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var ctx = context.Background()

func Auth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return utils.SendError(c, http.StatusUnauthorized, "Missing Authorization header", "")
			}

			parts := strings.Split(authHeader, " ")
			var tokenString string

			if len(parts) == 2 && parts[0] == "Bearer" {
				tokenString = parts[1]
			} else if len(parts) == 1 {
				// Fitur Toleransi: Jika user lupa mengetik "Bearer ", kita anggap seluruh teks adalah token
				tokenString = parts[0]
			} else {
				return utils.SendError(c, http.StatusUnauthorized, "Format Authorization tidak valid", "")
			}

			// [Keamanan Pilar 2] Cek JWT Blacklist di Redis
			if redisPkg.Client != nil {
				isBlacklisted, _ := redisPkg.Client.Exists(ctx, "jwt_blacklist:"+tokenString).Result()
				if isBlacklisted > 0 {
					return utils.SendError(c, http.StatusUnauthorized, "Session telah diakhiri (Token Blacklisted)", "")
				}
			}

			secret := config.AppConfig.JWTSecret

			token, err := jwt.ParseWithClaims(tokenString, &utils.JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return utils.SendError(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			}

			if claims, ok := token.Claims.(*utils.JWTCustomClaims); ok && token.Valid {
				// [Keamanan Pilar Tambahan] Pengecekan Single Active Session
				if redisPkg.Client != nil {
					activeToken, _ := redisPkg.Client.Get(ctx, fmt.Sprintf("active_token:%d", claims.UserID)).Result()
					if activeToken != "" && activeToken != tokenString {
						return utils.SendError(c, http.StatusUnauthorized, "Sesi Anda telah berakhir karena akun ini baru saja login di perangkat lain.", "")
					}
				}

				c.Set("user_id", claims.UserID)
				c.Set("company_id", claims.CompanyID)
				c.Set("role", claims.Role)
			}

			return next(c)
		}
	}
}

// RequireRole adalah middleware tingkat lanjut untuk mengecek apakah user memiliki hak akses (RBAC).
// Contoh penggunaan: e.GET("/admin", Handler, middleware.Auth(), middleware.RequireRole("superadmin", "admin"))
func RequireRole(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			userRole, ok := c.Get("role").(string)
			if !ok || userRole == "" {
				return utils.SendError(c, http.StatusForbidden, "Akses ditolak: Role tidak ditemukan", "")
			}

			// Cek apakah role user ada di dalam daftar role yang diizinkan
			isAllowed := false
			for _, role := range allowedRoles {
				if userRole == role {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				return utils.SendError(c, http.StatusForbidden, "Akses ditolak: Anda tidak memiliki hak akses (Forbidden)", "")
			}

			return next(c)
		}
	}
}
