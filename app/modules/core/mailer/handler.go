package mailer

import (
	"net/http"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// TestSendMailHandler godoc
// @Summary Test email sending
// @Description Queue a test email via RabbitMQ
// @Tags mailer
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/mailer/test-send [post]
// @Security BearerAuth
func TestSendMailHandler(c echo.Context) error {
	// Call existing QueueEmail function
	err := QueueEmail("test@example.com", "Test Swagger", "Hello from Swagger UI!")
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to queue email", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Email queued to RabbitMQ successfully", nil)
}
