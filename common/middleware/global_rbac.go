package middleware

import (
	"net/http"
	"strings"

	"ERP-System/common/constants"
	"ERP-System/config"

	"github.com/labstack/echo/v4"
)

// GlobalAutoRBAC dijalankan SETELAH Auth() membaca token.
func GlobalAutoRBAC() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			path := c.Request().URL.Path

			segments := strings.Split(path, "/")
			if len(segments) < 3 {
				return next(c)
			}
			moduleName := segments[2]
			if len(segments) > 3 {
				moduleName = segments[2] + "/" + segments[3]
			}

			action := "read"
			method := c.Request().Method
			if method == "POST" || method == "PUT" || method == "PATCH" {
				action = "write"
			} else if method == "DELETE" {
				action = "delete"
			}

			// Terima bersih dari Auth()
			roleVal := c.Get("role")
			if roleVal == nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Akses ditolak: Anda belum login atau sesi habis (RBAC)"})
			}
			roleString := roleVal.(string)
			userRoles := strings.Split(roleString, ",")

			for _, r := range userRoles {
				if strings.TrimSpace(r) == string(constants.RoleSuperadmin) {
					return next(c)
				}
			}

			isAllowed := false
			for _, r := range userRoles {
				roleName := strings.TrimSpace(r)
				var perm struct {
					CanRead   bool
					CanWrite  bool
					CanDelete bool
				}
				err := config.DB.Table("setting.role_permissions").Where("role_name = ? AND module = ?", roleName, moduleName).Scan(&perm).Error
				if err == nil {
					if action == "read" && perm.CanRead { isAllowed = true; break }
					if action == "write" && perm.CanWrite { isAllowed = true; break }
					if action == "delete" && perm.CanDelete { isAllowed = true; break }
				}
			}

			if isAllowed {
				return next(c)
			}

			return c.JSON(http.StatusForbidden, map[string]string{
				"error": "Akses Ditolak. Superadmin telah menonaktifkan fitur '" + moduleName + "' untuk peran Anda.",
			})
		}
	}
}
