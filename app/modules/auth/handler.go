package auth

import (
	"net/http"

	"ERP-System/common/utils"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginHandler godoc
// @Summary Login User
// @Description Authenticate user and get JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Router /api/auth/login [post]
func LoginHandler(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format JSON tidak valid", err.Error())
	}

	// Panggil validasi (menggunakan package validator yang sudah kita buat)
	val := utils.NewValidator()
	if err := val.Validate(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Data tidak lengkap", err.Error())
	}

	token, appErr := LoginService(req.Email, req.Password)
	if appErr != nil {
		return utils.SendError(c, appErr.Code, appErr.Message, "")
	}

	return utils.SendSuccess(c, http.StatusOK, "Login berhasil!", map[string]string{
		"token": token,
	})
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/auth/register [post]
func RegisterHandler(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Format JSON tidak valid", err.Error())
	}

	val := utils.NewValidator()
	if err := val.Validate(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Data tidak lengkap atau tidak valid", err.Error())
	}

	user, appErr := RegisterService(req.Name, req.Email, req.Password)
	if appErr != nil {
		return utils.SendError(c, appErr.Code, appErr.Message, "")
	}

	return utils.SendSuccess(c, http.StatusCreated, "Registrasi berhasil!", user)
}


