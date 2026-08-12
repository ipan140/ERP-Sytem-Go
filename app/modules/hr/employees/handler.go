package employees

import (
	"net/http"
	"strconv"
	"ERP-System/common/utils"
	"github.com/labstack/echo/v4"
)

// CreateEmployee godoc
// @Summary Create a new Employee
// @Description Create a new Employee in the system
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees [post]
// @Security BearerAuth
func CreateEmployeeHandler(c echo.Context) error {
	var data Employee
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateEmployeeService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllEmployee godoc
// @Summary Get all Employee
// @Description Retrieve a list of all Employee
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees [get]
// @Security BearerAuth
func GetAllEmployeeHandler(c echo.Context) error {
	data, err := GetAllEmployeeService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetEmployeeByID godoc
// @Summary Get a Employee by ID
// @Description Retrieve a specific Employee by its ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/{id} [get]
// @Security BearerAuth
func GetEmployeeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateEmployee godoc
// @Summary Update a Employee
// @Description Update an existing Employee
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/{id} [put]
// @Security BearerAuth
func UpdateEmployeeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateEmployeeService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteEmployee godoc
// @Summary Delete a Employee
// @Description Delete a Employee by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/{id} [delete]
// @Security BearerAuth
func DeleteEmployeeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}
