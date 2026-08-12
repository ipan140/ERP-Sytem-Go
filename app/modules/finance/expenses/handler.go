package expenses

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateExpense godoc
// @Summary Create a new Expense
// @Description Create a new Expense in the system
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/finance/expenses [post]
// @Security BearerAuth
func CreateExpenseHandler(c echo.Context) error {
	var data Expense
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateExpenseService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllExpense godoc
// @Summary Get all Expense
// @Description Retrieve a list of all Expense
// @Tags finance-expenses
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses [get]
// @Security BearerAuth
func GetAllExpenseHandler(c echo.Context) error {
	data, err := GetAllExpenseService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetExpenseByID godoc
// @Summary Get a Expense by ID
// @Description Retrieve a specific Expense by its ID
// @Tags finance-expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/{id} [get]
// @Security BearerAuth
func GetExpenseByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateExpense godoc
// @Summary Update a Expense
// @Description Update an existing Expense
// @Tags finance-expenses
// @Accept json
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/{id} [put]
// @Security BearerAuth
func UpdateExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateExpenseService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteExpense godoc
// @Summary Delete a Expense
// @Description Delete a Expense by ID
// @Tags finance-expenses
// @Produce json
// @Param id path int true "Expense ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/expenses/{id} [delete]
// @Security BearerAuth
func DeleteExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteExpenseService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
