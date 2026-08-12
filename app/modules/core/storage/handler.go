package storage

import (
	"net/http"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

func UploadFileHandler(c echo.Context) error {
	// TODO: Implementasi logika multipart/form-data untuk menyimpan file ke disk atau S3
	
	// Mengambil informasi siapa yang sedang login (dari token JWT)
	userID := c.Get("user_id")

	return utils.SendSuccess(c, http.StatusOK, "Fitur upload file belum diimplementasi secara penuh", map[string]interface{}{
		"uploaded_by": userID,
	})
}
