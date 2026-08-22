package payroll

import (
	"ERP-System/common/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreatePayslip godoc
// @Summary Create a new Payslip
// @Description Create a new Payslip in the system
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Success 201 {object} Payslip
// @Param request body Payslip true "Payload"
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
// @Success 200 {object} []Payslip
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
// @Success 200 {object} Payslip
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

// @Summary Create PayslipLine
// @Description Create a new PayslipLine
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Success 201 {object} PayslipLine
// @Param request body PayslipLine true "Payload"
// @Router /api/hr/payroll/payslipline [post]
// @Security BearerAuth
func CreatePayslipLineHandler(c echo.Context) error {
	var data PayslipLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePayslipLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all PayslipLine
// @Description Retrieve a list of all PayslipLine
// @Tags hr-payroll
// @Produce json
// @Success 200 {object} PayslipLine
// @Router /api/hr/payroll/payslipline [get]
// @Security BearerAuth
func GetAllPayslipLineHandler(c echo.Context) error {
	data, err := GetAllPayslipLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetPayslipLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPayslipLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update PayslipLine
// @Description Update an existing PayslipLine
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Param id path int true "PayslipLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/payslipline/{id} [put]
// @Security BearerAuth
func UpdatePayslipLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPayslipLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdatePayslipLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete PayslipLine
// @Description Delete PayslipLine by ID
// @Tags hr-payroll
// @Produce json
// @Param id path int true "PayslipLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/payslipline/{id} [delete]
// @Security BearerAuth
func DeletePayslipLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePayslipLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SalaryRule
// @Description Create a new SalaryRule
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Success 201 {object} SalaryRule
// @Param request body SalaryRule true "Payload"
// @Router /api/hr/payroll/salaryrule [post]
// @Security BearerAuth
func CreateSalaryRuleHandler(c echo.Context) error {
	var data SalaryRule
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSalaryRuleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SalaryRule
// @Description Retrieve a list of all SalaryRule
// @Tags hr-payroll
// @Produce json
// @Success 200 {object} SalaryRule
// @Router /api/hr/payroll/salaryrule [get]
// @Security BearerAuth
func GetAllSalaryRuleHandler(c echo.Context) error {
	data, err := GetAllSalaryRuleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSalaryRuleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalaryRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SalaryRule
// @Description Update an existing SalaryRule
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Param id path int true "SalaryRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/salaryrule/{id} [put]
// @Security BearerAuth
func UpdateSalaryRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSalaryRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSalaryRuleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SalaryRule
// @Description Delete SalaryRule by ID
// @Tags hr-payroll
// @Produce json
// @Param id path int true "SalaryRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/payroll/salaryrule/{id} [delete]
// @Security BearerAuth
func DeleteSalaryRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSalaryRuleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


