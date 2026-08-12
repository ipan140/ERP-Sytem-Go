package storage

import (
	"net/http"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// UploadFileHandler godoc
// @Summary Upload a file
// @Description Upload a file for storage (attachments)
// @Tags storage
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/storage/upload [post]
// @Security BearerAuth
func UploadFileHandler(c echo.Context) error {
	userID := c.Get("user_id")
	return utils.SendSuccess(c, http.StatusOK, "File uploaded successfully (mock)", map[string]interface{}{
		"uploaded_by": userID,
	})
}
