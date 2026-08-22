package mailer

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateEmailLog godoc
// @Summary Create a new EmailLog
// @Description Create a new EmailLog in the system
// @Tags core-mailer
// @Accept json
// @Produce json
// @Success 201 {object} EmailLog
// @Param request body EmailLog true "Payload"
// @Router /api/core/mailer [post]
// @Security BearerAuth
func CreateEmailLogHandler(c echo.Context) error {
	var data EmailLog
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateEmailLogService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllEmailLog godoc
// @Summary Get all EmailLog
// @Description Retrieve a list of all EmailLog
// @Tags core-mailer
// @Produce json
// @Success 200 {object} []EmailLog
// @Router /api/core/mailer [get]
// @Security BearerAuth
func GetAllEmailLogHandler(c echo.Context) error {
	data, err := GetAllEmailLogService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetEmailLogByID godoc
// @Summary Get a EmailLog by ID
// @Description Retrieve a specific EmailLog by its ID
// @Tags core-mailer
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} EmailLog
// @Router /api/core/mailer/{id} [get]
// @Security BearerAuth
func GetEmailLogByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmailLogByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateEmailLog godoc
// @Summary Update a EmailLog
// @Description Update an existing EmailLog
// @Tags core-mailer
// @Accept json
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/mailer/{id} [put]
// @Security BearerAuth
func UpdateEmailLogHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmailLogByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateEmailLogService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteEmailLog godoc
// @Summary Delete a EmailLog
// @Description Delete a EmailLog by ID
// @Tags core-mailer
// @Produce json
// @Param id path int true "EmailLog ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/core/mailer/{id} [delete]
// @Security BearerAuth
func DeleteEmailLogHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmailLogService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}


