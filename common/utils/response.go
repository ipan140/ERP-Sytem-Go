package utils

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func SendSuccess(c echo.Context, statusCode int, message string, data interface{}) error {
	return c.JSON(statusCode, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func SendError(c echo.Context, statusCode int, message string, err string) error {
	return c.JSON(statusCode, ErrorResponse{
		Message: message,
		Error:   err,
	})
}
