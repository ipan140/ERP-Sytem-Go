package middleware

import (
	"net/http"
	"strings"

	"ERP-System/common/constants"
	"ERP-System/common/utils"

	"github.com/labstack/echo/v4"
)

// Fungsi ini menerima daftar role apa saja yang diperbolehkan masuk ke rute tertentu.
func RequireRoles(allowedRoles ...constants.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Ambil data role user dari Context (diset otomatis oleh middleware Auth JWT)
			roleVal := c.Get("role")
			if roleVal == nil {
				return utils.SendError(c, http.StatusUnauthorized, "Akses Ditolak", "Anda belum login atau sesi telah habis.")
			}

			roleString, ok := roleVal.(string)
			if !ok || roleString == "" {
				return utils.SendError(c, http.StatusUnauthorized, "Akses Ditolak", "Role tidak valid atau kosong.")
			}

			// Anggap satu user bisa punya banyak role (dipisah koma)
			userRoles := strings.Split(roleString, ",")

			// 2. Jika dia adalah SUPERADMIN, selalu izinkan masuk tanpa periksa lagi!
			for _, userRole := range userRoles {
				if strings.TrimSpace(userRole) == string(constants.RoleSuperadmin) {
					return next(c)
				}
			}

			// 3. Cocokkan role user dengan role yang diizinkan (allowedRoles)
			hasAccess := false
			for _, userRole := range userRoles {
				for _, allowedRole := range allowedRoles {
					if strings.TrimSpace(userRole) == string(allowedRole) {
						hasAccess = true
						break
					}
				}
				if hasAccess {
					break
				}
			}

			// 4. Jika cocok, silakan masuk. Jika tidak, tendang keluar!
			if hasAccess {
				return next(c)
			}

			return utils.SendError(c, http.StatusForbidden, "Akses Terlarang", "Anda tidak memiliki hak akses (Role) untuk membuka menu ini.")
		}
	}
}
