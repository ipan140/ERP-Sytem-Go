package user_roles

import (
	"net/http"
	"strings"

	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// GetAllUsersRolesHandler godoc
// @Summary Get all users with their roles
// @Description Retrieve a list of all users and their assigned roles (Superadmin only)
// @Tags Core - User Roles
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /core/user_roles [get]
func GetAllUsersRolesHandler(c echo.Context) error {
	users, err := GetAllUsersRolesService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal Mengambil Data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Berhasil mengambil daftar hak akses pengguna", users)
}

// AssignRoleHandler godoc
// @Summary Assign roles to a user
// @Description Assign multiple roles to a specific user by their ID (Superadmin only)
// @Tags Core - User Roles
// @Accept json
// @Produce json
// @Param body body AssignRoleRequest true "Assign Role Request Payload"
// @Success 200 {object} map[string]interface{}
// @Security BearerAuth
// @Router /core/user_roles/assign [post]
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
