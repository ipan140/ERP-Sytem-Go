package payroll

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreatePayslip godoc
// @Summary Create a new Payslip
// @Description Create a new Payslip in the system
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/payroll [post]
// @Security BearerAuth
func CreatePayslipHandler(c echo.Context) error {
	var data Payslip
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreatePayslipService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllPayslip godoc
// @Summary Get all Payslip
// @Description Retrieve a list of all Payslip
// @Tags hr-payroll
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll [get]
// @Security BearerAuth
func GetAllPayslipHandler(c echo.Context) error {
	data, err := GetAllPayslipService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetPayslipByID godoc
// @Summary Get a Payslip by ID
// @Description Retrieve a specific Payslip by its ID
// @Tags hr-payroll
// @Produce json
// @Param id path int true "Payslip ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/{id} [get]
// @Security BearerAuth
func GetPayslipByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPayslipByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdatePayslip godoc
// @Summary Update a Payslip
// @Description Update an existing Payslip
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Param id path int true "Payslip ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/{id} [put]
// @Security BearerAuth
func UpdatePayslipHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPayslipByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdatePayslipService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeletePayslip godoc
// @Summary Delete a Payslip
// @Description Delete a Payslip by ID
// @Tags hr-payroll
// @Produce json
// @Param id path int true "Payslip ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/{id} [delete]
// @Security BearerAuth
func DeletePayslipHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePayslipService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
