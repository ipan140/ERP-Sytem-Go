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
