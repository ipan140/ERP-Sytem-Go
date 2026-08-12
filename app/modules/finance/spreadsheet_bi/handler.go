package spreadsheet_bi

import (
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

// CreateSpreadsheet godoc
// @Summary Create a new Spreadsheet
// @Description Create a new Spreadsheet in the system
// @Tags finance-spreadsheet_bi
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/spreadsheet_bi [post]
// @Security BearerAuth
func CreateSpreadsheetHandler(c echo.Context) error {
	var data Spreadsheet
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateSpreadsheetService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllSpreadsheet godoc
// @Summary Get all Spreadsheet
// @Description Retrieve a list of all Spreadsheet
// @Tags finance-spreadsheet_bi
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/spreadsheet_bi [get]
// @Security BearerAuth
func GetAllSpreadsheetHandler(c echo.Context) error {
	data, err := GetAllSpreadsheetService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetSpreadsheetByID godoc
// @Summary Get a Spreadsheet by ID
// @Description Retrieve a specific Spreadsheet by its ID
// @Tags finance-spreadsheet_bi
// @Produce json
// @Param id path int true "Spreadsheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/spreadsheet_bi/{id} [get]
// @Security BearerAuth
func GetSpreadsheetByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSpreadsheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateSpreadsheet godoc
// @Summary Update a Spreadsheet
// @Description Update an existing Spreadsheet
// @Tags finance-spreadsheet_bi
// @Accept json
// @Produce json
// @Param id path int true "Spreadsheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/spreadsheet_bi/{id} [put]
// @Security BearerAuth
func UpdateSpreadsheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSpreadsheetByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateSpreadsheetService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteSpreadsheet godoc
// @Summary Delete a Spreadsheet
// @Description Delete a Spreadsheet by ID
// @Tags finance-spreadsheet_bi
// @Produce json
// @Param id path int true "Spreadsheet ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/spreadsheet_bi/{id} [delete]
// @Security BearerAuth
func DeleteSpreadsheetHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSpreadsheetService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
