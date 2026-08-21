package middleware

import (
	"net/http"
	"strings"

	"ERP-System/common/constants"
	"ERP-System/common/utils"

	"github.com/labstack/echo/v4"
)

// RequireRoles adalah Satpam Utama.
// Fungsi ini menerima daftar role apa saja yang diperbolehkan masuk ke rute tertentu.
func RequireRoles(allowedRoles ...constants.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Ambil data role user dari Token (JWT) yang sedang login.
			// Biasanya saat login, role disimpan di JWT Claims.
			// Untuk contoh ini, kita asumsikan Anda menyimpannya dalam header atau konteks.

			// Misal Anda mengambilnya dari context yang di-set oleh JWT Middleware:
			// userClaims := c.Get("user").(*jwt.Token).Claims.(*JwtCustomClaims)
			// userRoles := userClaims.Roles

			// SIMULASI SEMENTARA: Kita ambil dari Header "X-User-Roles"
			// (Ganti bagian ini nanti dengan penarikan JWT sungguhan Anda)
			roleString := c.Request().Header.Get("X-User-Roles")
			if roleString == "" {
				return utils.SendError(c, http.StatusUnauthorized, "Akses Ditolak", "Anda belum login atau tidak memiliki tiket akses.")
			}

			// Anggap satu user bisa punya banyak role (dipisah koma)
			userRoles := strings.Split(roleString, ",")

			// 2. Jika dia adalah SUPERADMIN, selalu izinkan masuk tanpa periksa lagi!
			for _, userRole := range userRoles {
				if userRole == string(constants.RoleSuperadmin) {
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
