package base

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateCurrency godoc
// @Summary Create a new Currency
// @Description Create a new Currency in the system
// @Tags base
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/base [post]
// @Security BearerAuth
func CreateCurrencyHandler(c echo.Context) error {
	var data Currency
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateCurrencyService(&data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllCurrency godoc
// @Summary Get all Currency
// @Description Retrieve a list of all Currency
// @Tags base
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/base [get]
// @Security BearerAuth
func GetAllCurrencyHandler(c echo.Context) error {
	data, err := GetAllCurrencyService()
	if err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetCurrencyByID godoc
// @Summary Get a Currency by ID
// @Description Retrieve a specific Currency by its ID
// @Tags base
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/base/{id} [get]
// @Security BearerAuth
func GetCurrencyByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCurrencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateCurrency godoc
// @Summary Update a Currency
// @Description Update an existing Currency
// @Tags base
// @Accept json
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/base/{id} [put]
// @Security BearerAuth
func UpdateCurrencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetCurrencyByIDService(uint(id))
	if err != nil {
		return utils.SendError(c,  http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c,  http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateCurrencyService(data); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteCurrency godoc
// @Summary Delete a Currency
// @Description Delete a Currency by ID
// @Tags base
// @Produce json
// @Param id path int true "Currency ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/base/{id} [delete]
// @Security BearerAuth
func DeleteCurrencyHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteCurrencyService(uint(id)); err != nil {
		return utils.SendError(c,  http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
