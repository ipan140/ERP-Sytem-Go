package accounting

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateJournalEntry godoc
// @Summary Create a new JournalEntry
// @Description Create a new JournalEntry in the system
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/accounting [post]
// @Security BearerAuth
func CreateJournalEntryHandler(c echo.Context) error {
	var data JournalEntry
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateJournalEntryService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllJournalEntry godoc
// @Summary Get all JournalEntry
// @Description Retrieve a list of all JournalEntry
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting [get]
// @Security BearerAuth
func GetAllJournalEntryHandler(c echo.Context) error {
	data, err := GetAllJournalEntryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetJournalEntryByID godoc
// @Summary Get a JournalEntry by ID
// @Description Retrieve a specific JournalEntry by its ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [get]
// @Security BearerAuth
func GetJournalEntryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateJournalEntry godoc
// @Summary Update a JournalEntry
// @Description Update an existing JournalEntry
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [put]
// @Security BearerAuth
func UpdateJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateJournalEntryService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteJournalEntry godoc
// @Summary Delete a JournalEntry
// @Description Delete a JournalEntry by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [delete]
// @Security BearerAuth
func DeleteJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJournalEntryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
