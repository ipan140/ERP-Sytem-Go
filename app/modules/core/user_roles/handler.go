package user_roles

import (
	"net/http"
	"strings"

	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

func GetAllUsersRolesHandler(c echo.Context) error {
	users, err := GetAllUsersRolesService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal Mengambil Data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Berhasil mengambil daftar hak akses pengguna", users)
}

func AssignRoleHandler(c echo.Context) error {
	var req AssignRoleRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format Data Salah", err.Error())
	}

	// Gabungkan array role menjadi string yang dipisah koma (contoh: "SALES_MANAGER,EMPLOYEE")
	rolesString := strings.Join(req.Roles, ",")

	err := AssignRoleService(req.UserID, rolesString)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal Menyimpan Role", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Hak Akses (Role) berhasil diperbarui!", nil)
}
