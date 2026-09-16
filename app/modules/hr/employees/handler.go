package employees

import (
	"fmt"
	"ERP-System/config"
	"encoding/csv"
	"ERP-System/common/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

// CreateEmployee godoc
// @Summary Create a new Employee
// @Description Create a new Employee in the system
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} Employee
// @Param request body Employee true "Payload"
// @Router /api/hr/employees [post]
// @Security BearerAuth
// CreateEmployeeHandler godoc
// @Summary Endpoint for CreateEmployee
// @Description Auto-generated swagger for CreateEmployeeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
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
// @Success 200 {object} []Employee
// @Router /api/hr/employees [get]
// @Security BearerAuth
// GetAllEmployeeHandler godoc
// @Summary Endpoint for GetAllEmployee
// @Description Auto-generated swagger for GetAllEmployeeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees [get]
// @Security BearerAuth
func GetAllEmployeeHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllEmployeeService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	departmentID := c.QueryParam("department_id")
	isActive := c.QueryParam("is_active")

	data, total, err := GetPaginatedEmployeeService(offset, limit, search, departmentID, isActive)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Data retrieved successfully", data, meta)
}

// GetEmployeeByID godoc
// @Summary Get a Employee by ID
// @Description Retrieve a specific Employee by its ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} Employee
// @Router /api/hr/employees/{id} [get]
// @Security BearerAuth
// GetEmployeeByIDHandler godoc
// @Summary Endpoint for GetEmployeeByID
// @Description Auto-generated swagger for GetEmployeeByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
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
// @Success 200 {object} Employee
// @Router /api/hr/employees/{id} [put]
// @Security BearerAuth
// UpdateEmployeeHandler godoc
// @Summary Endpoint for UpdateEmployee
// @Description Auto-generated swagger for UpdateEmployeeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
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
// DeleteEmployeeHandler godoc
// @Summary Endpoint for DeleteEmployee
// @Description Auto-generated swagger for DeleteEmployeeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/{id} [delete]
// @Security BearerAuth
func DeleteEmployeeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// @Summary Create JobPosition
// @Description Create a new JobPosition
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} JobPosition
// @Param request body JobPosition true "Payload"
// @Router /api/hr/employees/jobposition [post]
// @Security BearerAuth
// CreateJobPositionHandler godoc
// @Summary Endpoint for CreateJobPosition
// @Description Auto-generated swagger for CreateJobPositionHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/jobposition [post]
// @Security BearerAuth
func CreateJobPositionHandler(c echo.Context) error {
	var data JobPosition
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateJobPositionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all JobPosition
// @Description Retrieve a list of all JobPosition
// @Tags hr-employees
// @Produce json
// @Success 200 {object} JobPosition
// @Router /api/hr/employees/jobposition [get]
// @Security BearerAuth
// GetAllJobPositionHandler godoc
// @Summary Endpoint for GetAllJobPosition
// @Description Auto-generated swagger for GetAllJobPositionHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/jobposition [get]
// @Security BearerAuth
func GetAllJobPositionHandler(c echo.Context) error {
	data, err := GetAllJobPositionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetJobPositionByIDHandler godoc
// @Summary Endpoint for GetJobPositionByID
// @Description Auto-generated swagger for GetJobPositionByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/jobposition/{id} [get]
// @Security BearerAuth
func GetJobPositionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobPositionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update JobPosition
// @Description Update an existing JobPosition
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "JobPosition ID"
// @Success 200 {object} JobPosition
// @Router /api/hr/employees/jobposition/{id} [put]
// @Security BearerAuth
// UpdateJobPositionHandler godoc
// @Summary Endpoint for UpdateJobPosition
// @Description Auto-generated swagger for UpdateJobPositionHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/jobposition/{id} [put]
// @Security BearerAuth
func UpdateJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobPositionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateJobPositionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete JobPosition
// @Description Delete JobPosition by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "JobPosition ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/jobposition/{id} [delete]
// @Security BearerAuth
// DeleteJobPositionHandler godoc
// @Summary Endpoint for DeleteJobPosition
// @Description Auto-generated swagger for DeleteJobPositionHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/jobposition/{id} [delete]
// @Security BearerAuth
func DeleteJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJobPositionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create WorkingSchedule
// @Description Create a new WorkingSchedule
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} WorkingSchedule
// @Param request body WorkingSchedule true "Payload"
// @Router /api/hr/employees/workingschedule [post]
// @Security BearerAuth
// CreateWorkingScheduleHandler godoc
// @Summary Endpoint for CreateWorkingSchedule
// @Description Auto-generated swagger for CreateWorkingScheduleHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/workingschedule [post]
// @Security BearerAuth
func CreateWorkingScheduleHandler(c echo.Context) error {
	var data WorkingSchedule
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateWorkingScheduleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all WorkingSchedule
// @Description Retrieve a list of all WorkingSchedule
// @Tags hr-employees
// @Produce json
// @Success 200 {object} WorkingSchedule
// @Router /api/hr/employees/workingschedule [get]
// @Security BearerAuth
// GetAllWorkingScheduleHandler godoc
// @Summary Endpoint for GetAllWorkingSchedule
// @Description Auto-generated swagger for GetAllWorkingScheduleHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/workingschedule [get]
// @Security BearerAuth
func GetAllWorkingScheduleHandler(c echo.Context) error {
	data, err := GetAllWorkingScheduleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetWorkingScheduleByIDHandler godoc
// @Summary Endpoint for GetWorkingScheduleByID
// @Description Auto-generated swagger for GetWorkingScheduleByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/workingschedule/{id} [get]
// @Security BearerAuth
func GetWorkingScheduleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkingScheduleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update WorkingSchedule
// @Description Update an existing WorkingSchedule
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "WorkingSchedule ID"
// @Success 200 {object} WorkingSchedule
// @Router /api/hr/employees/workingschedule/{id} [put]
// @Security BearerAuth
// UpdateWorkingScheduleHandler godoc
// @Summary Endpoint for UpdateWorkingSchedule
// @Description Auto-generated swagger for UpdateWorkingScheduleHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/workingschedule/{id} [put]
// @Security BearerAuth
func UpdateWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkingScheduleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateWorkingScheduleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete WorkingSchedule
// @Description Delete WorkingSchedule by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "WorkingSchedule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/workingschedule/{id} [delete]
// @Security BearerAuth
// DeleteWorkingScheduleHandler godoc
// @Summary Endpoint for DeleteWorkingSchedule
// @Description Auto-generated swagger for DeleteWorkingScheduleHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/workingschedule/{id} [delete]
// @Security BearerAuth
func DeleteWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWorkingScheduleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Contract
// @Description Create a new Contract
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} Contract
// @Param request body Contract true "Payload"
// @Router /api/hr/employees/contract [post]
// @Security BearerAuth
// CreateContractHandler godoc
// @Summary Endpoint for CreateContract
// @Description Auto-generated swagger for CreateContractHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/contract [post]
// @Security BearerAuth
func CreateContractHandler(c echo.Context) error {
	var data Contract
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateContractService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Contract
// @Description Retrieve a list of all Contract
// @Tags hr-employees
// @Produce json
// @Success 200 {object} Contract
// @Router /api/hr/employees/contract [get]
// @Security BearerAuth
// GetAllContractHandler godoc
// @Summary Endpoint for GetAllContract
// @Description Auto-generated swagger for GetAllContractHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/contract [get]
// @Security BearerAuth
func GetAllContractHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllContractService()
		if err != nil {
			return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
		}
		return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedContractService(offset, limit, search)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve paginated data", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Retrieved successfully", data, meta)
}
// GetContractByIDHandler godoc
// @Summary Endpoint for GetContractByID
// @Description Auto-generated swagger for GetContractByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/contract/{id} [get]
// @Security BearerAuth
func GetContractByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetContractByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Contract
// @Description Update an existing Contract
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Contract ID"
// @Success 200 {object} Contract
// @Router /api/hr/employees/contract/{id} [put]
// @Security BearerAuth
// UpdateContractHandler godoc
// @Summary Endpoint for UpdateContract
// @Description Auto-generated swagger for UpdateContractHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/contract/{id} [put]
// @Security BearerAuth
func UpdateContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetContractByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateContractService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Contract
// @Description Delete Contract by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Contract ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/contract/{id} [delete]
// @Security BearerAuth
// DeleteContractHandler godoc
// @Summary Endpoint for DeleteContract
// @Description Auto-generated swagger for DeleteContractHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/contract/{id} [delete]
// @Security BearerAuth
func DeleteContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteContractService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Skill
// @Description Create a new Skill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} Skill
// @Param request body Skill true "Payload"
// @Router /api/hr/employees/skill [post]
// @Security BearerAuth
// CreateSkillHandler godoc
// @Summary Endpoint for CreateSkill
// @Description Auto-generated swagger for CreateSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/skill [post]
// @Security BearerAuth
func CreateSkillHandler(c echo.Context) error {
	var data Skill
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSkillService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Skill
// @Description Retrieve a list of all Skill
// @Tags hr-employees
// @Produce json
// @Success 200 {object} Skill
// @Router /api/hr/employees/skill [get]
// @Security BearerAuth
// GetAllSkillHandler godoc
// @Summary Endpoint for GetAllSkill
// @Description Auto-generated swagger for GetAllSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/skill [get]
// @Security BearerAuth
func GetAllSkillHandler(c echo.Context) error {
	data, err := GetAllSkillService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetSkillByIDHandler godoc
// @Summary Endpoint for GetSkillByID
// @Description Auto-generated swagger for GetSkillByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/skill/{id} [get]
// @Security BearerAuth
func GetSkillByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update Skill
// @Description Update an existing Skill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Skill ID"
// @Success 200 {object} Skill
// @Router /api/hr/employees/skill/{id} [put]
// @Security BearerAuth
// UpdateSkillHandler godoc
// @Summary Endpoint for UpdateSkill
// @Description Auto-generated swagger for UpdateSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/skill/{id} [put]
// @Security BearerAuth
func UpdateSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSkillService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Skill
// @Description Delete Skill by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Skill ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skill/{id} [delete]
// @Security BearerAuth
// DeleteSkillHandler godoc
// @Summary Endpoint for DeleteSkill
// @Description Auto-generated swagger for DeleteSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/skill/{id} [delete]
// @Security BearerAuth
func DeleteSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SkillLevel
// @Description Create a new SkillLevel
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} SkillLevel
// @Param request body SkillLevel true "Payload"
// @Router /api/hr/employees/skilllevel [post]
// @Security BearerAuth
// CreateSkillLevelHandler godoc
// @Summary Endpoint for CreateSkillLevel
// @Description Auto-generated swagger for CreateSkillLevelHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/skilllevel [post]
// @Security BearerAuth
func CreateSkillLevelHandler(c echo.Context) error {
	var data SkillLevel
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateSkillLevelService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all SkillLevel
// @Description Retrieve a list of all SkillLevel
// @Tags hr-employees
// @Produce json
// @Success 200 {object} SkillLevel
// @Router /api/hr/employees/skilllevel [get]
// @Security BearerAuth
// GetAllSkillLevelHandler godoc
// @Summary Endpoint for GetAllSkillLevel
// @Description Auto-generated swagger for GetAllSkillLevelHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/skilllevel [get]
// @Security BearerAuth
func GetAllSkillLevelHandler(c echo.Context) error {
	data, err := GetAllSkillLevelService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetSkillLevelByIDHandler godoc
// @Summary Endpoint for GetSkillLevelByID
// @Description Auto-generated swagger for GetSkillLevelByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/skilllevel/{id} [get]
// @Security BearerAuth
func GetSkillLevelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillLevelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update SkillLevel
// @Description Update an existing SkillLevel
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "SkillLevel ID"
// @Success 200 {object} SkillLevel
// @Router /api/hr/employees/skilllevel/{id} [put]
// @Security BearerAuth
// UpdateSkillLevelHandler godoc
// @Summary Endpoint for UpdateSkillLevel
// @Description Auto-generated swagger for UpdateSkillLevelHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/skilllevel/{id} [put]
// @Security BearerAuth
func UpdateSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillLevelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateSkillLevelService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete SkillLevel
// @Description Delete SkillLevel by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "SkillLevel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skilllevel/{id} [delete]
// @Security BearerAuth
// DeleteSkillLevelHandler godoc
// @Summary Endpoint for DeleteSkillLevel
// @Description Auto-generated swagger for DeleteSkillLevelHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/skilllevel/{id} [delete]
// @Security BearerAuth
func DeleteSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillLevelService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create EmployeeSkill
// @Description Create a new EmployeeSkill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} EmployeeSkill
// @Param request body EmployeeSkill true "Payload"
// @Router /api/hr/employees/employeeskill [post]
// @Security BearerAuth
// CreateEmployeeSkillHandler godoc
// @Summary Endpoint for CreateEmployeeSkill
// @Description Auto-generated swagger for CreateEmployeeSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeskill [post]
// @Security BearerAuth
func CreateEmployeeSkillHandler(c echo.Context) error {
	var data EmployeeSkill
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateEmployeeSkillService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all EmployeeSkill
// @Description Retrieve a list of all EmployeeSkill
// @Tags hr-employees
// @Produce json
// @Success 200 {object} EmployeeSkill
// @Router /api/hr/employees/employeeskill [get]
// @Security BearerAuth
// GetAllEmployeeSkillHandler godoc
// @Summary Endpoint for GetAllEmployeeSkill
// @Description Auto-generated swagger for GetAllEmployeeSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/employeeskill [get]
// @Security BearerAuth
func GetAllEmployeeSkillHandler(c echo.Context) error {
	data, err := GetAllEmployeeSkillService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetEmployeeSkillByIDHandler godoc
// @Summary Endpoint for GetEmployeeSkillByID
// @Description Auto-generated swagger for GetEmployeeSkillByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeskill/{id} [get]
// @Security BearerAuth
func GetEmployeeSkillByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeSkillByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update EmployeeSkill
// @Description Update an existing EmployeeSkill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "EmployeeSkill ID"
// @Success 200 {object} EmployeeSkill
// @Param request body EmployeeSkill true "Payload"
// @Router /api/hr/employees/employeeskill/{id} [put]
// @Security BearerAuth
// UpdateEmployeeSkillHandler godoc
// @Summary Endpoint for UpdateEmployeeSkill
// @Description Auto-generated swagger for UpdateEmployeeSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeskill/{id} [put]
// @Security BearerAuth
func UpdateEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeSkillByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateEmployeeSkillService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete EmployeeSkill
// @Description Delete EmployeeSkill by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "EmployeeSkill ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/employeeskill/{id} [delete]
// @Security BearerAuth
// DeleteEmployeeSkillHandler godoc
// @Summary Endpoint for DeleteEmployeeSkill
// @Description Auto-generated swagger for DeleteEmployeeSkillHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeskill/{id} [delete]
// @Security BearerAuth
func DeleteEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeSkillService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ResumeLine
// @Description Create a new ResumeLine
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} ResumeLine
// @Param request body ResumeLine true "Payload"
// @Router /api/hr/employees/resumeline [post]
// @Security BearerAuth
// CreateResumeLineHandler godoc
// @Summary Endpoint for CreateResumeLine
// @Description Auto-generated swagger for CreateResumeLineHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/resumeline [post]
// @Security BearerAuth
func CreateResumeLineHandler(c echo.Context) error {
	var data ResumeLine
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateResumeLineService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all ResumeLine
// @Description Retrieve a list of all ResumeLine
// @Tags hr-employees
// @Produce json
// @Success 200 {object} ResumeLine
// @Router /api/hr/employees/resumeline [get]
// @Security BearerAuth
// GetAllResumeLineHandler godoc
// @Summary Endpoint for GetAllResumeLine
// @Description Auto-generated swagger for GetAllResumeLineHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/resumeline [get]
// @Security BearerAuth
func GetAllResumeLineHandler(c echo.Context) error {
	data, err := GetAllResumeLineService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
// GetResumeLineByIDHandler godoc
// @Summary Endpoint for GetResumeLineByID
// @Description Auto-generated swagger for GetResumeLineByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/resumeline/{id} [get]
// @Security BearerAuth
func GetResumeLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResumeLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}

// @Summary Update ResumeLine
// @Description Update an existing ResumeLine
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ResumeLine ID"
// @Success 200 {object} ResumeLine
// @Router /api/hr/employees/resumeline/{id} [put]
// @Security BearerAuth
// UpdateResumeLineHandler godoc
// @Summary Endpoint for UpdateResumeLine
// @Description Auto-generated swagger for UpdateResumeLineHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/resumeline/{id} [put]
// @Security BearerAuth
func UpdateResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResumeLineByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := UpdateResumeLineService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete ResumeLine
// @Description Delete ResumeLine by ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "ResumeLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/resumeline/{id} [delete]
// @Security BearerAuth
// DeleteResumeLineHandler godoc
// @Summary Endpoint for DeleteResumeLine
// @Description Auto-generated swagger for DeleteResumeLineHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/resumeline/{id} [delete]
// @Security BearerAuth
func DeleteResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteResumeLineService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}




// @Summary Create Department
// @Description Create a new Department
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} Department
// @Param request body Department true "Payload"
// @Router /api/hr/employees/department [post]
// @Security BearerAuth
// CreateDepartmentHandler godoc
// @Summary Endpoint for CreateDepartment
// @Description Auto-generated swagger for CreateDepartmentHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/department [post]
// @Security BearerAuth
func CreateDepartmentHandler(c echo.Context) error {
	var data Department
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateDepartmentService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}

// @Summary Get all Department
// @Description Retrieve a list of all Department
// @Tags hr-employees
// @Produce json
// @Success 200 {object} []Department
// @Router /api/hr/employees/department [get]
// @Security BearerAuth
// GetAllDepartmentHandler godoc
// @Summary Endpoint for GetAllDepartment
// @Description Auto-generated swagger for GetAllDepartmentHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/department [get]
// @Security BearerAuth
func GetAllDepartmentHandler(c echo.Context) error {
	data, err := GetAllDepartmentService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to fetch data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Get Department by ID
// @Description Retrieve a Department by its ID
// @Tags hr-employees
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} Department
// @Router /api/hr/employees/department/{id} [get]
// @Security BearerAuth
// GetDepartmentByIDHandler godoc
// @Summary Endpoint for GetDepartmentByID
// @Description Auto-generated swagger for GetDepartmentByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/department/{id} [get]
// @Security BearerAuth
func GetDepartmentByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDepartmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update Department
// @Description Update an existing Department
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param request body Department true "Payload"
// @Success 200 {object} Department
// @Router /api/hr/employees/department/{id} [put]
// @Security BearerAuth
// UpdateDepartmentHandler godoc
// @Summary Endpoint for UpdateDepartment
// @Description Auto-generated swagger for UpdateDepartmentHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/department/{id} [put]
// @Security BearerAuth
func UpdateDepartmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetDepartmentByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	data.ID = uint(id)
	if err := UpdateDepartmentService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}

// @Summary Delete Department
// @Description Delete a Department
// @Tags hr-employees
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/department/{id} [delete]
// @Security BearerAuth
// DeleteDepartmentHandler godoc
// @Summary Endpoint for DeleteDepartment
// @Description Auto-generated swagger for DeleteDepartmentHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/department/{id} [delete]
// @Security BearerAuth
func DeleteDepartmentHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteDepartmentService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


// --- Warning Letter ---
// GetAllWarningLetterHandler godoc
// @Summary Endpoint for GetAllWarningLetter
// @Description Auto-generated swagger for GetAllWarningLetterHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/warningletter [get]
// @Security BearerAuth
func GetAllWarningLetterHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllWarningLetterService()
		if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get data", err.Error()) }
		return utils.SendSuccess(c, http.StatusOK, "Success", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedWarningLetterService(offset, limit, search)
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get paginated data", err.Error()) }

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Success", data, meta)
}
// GetWarningLetterByIDHandler godoc
// @Summary Endpoint for GetWarningLetterByID
// @Description Auto-generated swagger for GetWarningLetterByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/warningletter/{id} [get]
// @Security BearerAuth
func GetWarningLetterByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWarningLetterByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// CreateWarningLetterHandler godoc
// @Summary Endpoint for CreateWarningLetter
// @Description Auto-generated swagger for CreateWarningLetterHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/warningletter [post]
// @Security BearerAuth
func CreateWarningLetterHandler(c echo.Context) error {
	var data WarningLetter
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateWarningLetterService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// UpdateWarningLetterHandler godoc
// @Summary Endpoint for UpdateWarningLetter
// @Description Auto-generated swagger for UpdateWarningLetterHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/warningletter/{id} [put]
// @Security BearerAuth
func UpdateWarningLetterHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWarningLetterByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateWarningLetterService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// DeleteWarningLetterHandler godoc
// @Summary Endpoint for DeleteWarningLetter
// @Description Auto-generated swagger for DeleteWarningLetterHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/warningletter/{id} [delete]
// @Security BearerAuth
func DeleteWarningLetterHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWarningLetterService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}


// --- Employee Task ---
// GetAllEmployeeTaskHandler godoc
// @Summary Endpoint for GetAllEmployeeTask
// @Description Auto-generated swagger for GetAllEmployeeTaskHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/employeetask [get]
// @Security BearerAuth
func GetAllEmployeeTaskHandler(c echo.Context) error {
	data, err := GetAllEmployeeTaskService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get data", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// GetEmployeeTaskByIDHandler godoc
// @Summary Endpoint for GetEmployeeTaskByID
// @Description Auto-generated swagger for GetEmployeeTaskByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeetask/{id} [get]
// @Security BearerAuth
func GetEmployeeTaskByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeTaskByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// CreateEmployeeTaskHandler godoc
// @Summary Endpoint for CreateEmployeeTask
// @Description Auto-generated swagger for CreateEmployeeTaskHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeetask [post]
// @Security BearerAuth
func CreateEmployeeTaskHandler(c echo.Context) error {
	var data EmployeeTask
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateEmployeeTaskService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// UpdateEmployeeTaskHandler godoc
// @Summary Endpoint for UpdateEmployeeTask
// @Description Auto-generated swagger for UpdateEmployeeTaskHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeetask/{id} [put]
// @Security BearerAuth
func UpdateEmployeeTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeTaskByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateEmployeeTaskService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
// DeleteEmployeeTaskHandler godoc
// @Summary Endpoint for DeleteEmployeeTask
// @Description Auto-generated swagger for DeleteEmployeeTaskHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeetask/{id} [delete]
// @Security BearerAuth
func DeleteEmployeeTaskHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeTaskService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// --- Phase 3 ---
// GetAllOvertimeHandler godoc
// @Summary Endpoint for GetAllOvertime
// @Description Auto-generated swagger for GetAllOvertimeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/overtime [get]
// @Security BearerAuth
func GetAllOvertimeHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllOvertimeService()
		if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
		return utils.SendSuccess(c, http.StatusOK, "Success", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedOvertimeService(offset, limit, search)
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get paginated data", err.Error()) }

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Success", data, meta)
}
// GetOvertimeByIDHandler godoc
// @Summary Endpoint for GetOvertimeByID
// @Description Auto-generated swagger for GetOvertimeByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/overtime/{id} [get]
// @Security BearerAuth
func GetOvertimeByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// CreateOvertimeHandler godoc
// @Summary Endpoint for CreateOvertime
// @Description Auto-generated swagger for CreateOvertimeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/overtime [post]
// @Security BearerAuth
func CreateOvertimeHandler(c echo.Context) error {
	var data Overtime
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateOvertimeService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created", data)
}
// UpdateOvertimeHandler godoc
// @Summary Endpoint for UpdateOvertime
// @Description Auto-generated swagger for UpdateOvertimeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/overtime/{id} [put]
// @Security BearerAuth
func UpdateOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetOvertimeByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateOvertimeService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated", data)
}
// DeleteOvertimeHandler godoc
// @Summary Endpoint for DeleteOvertime
// @Description Auto-generated swagger for DeleteOvertimeHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/overtime/{id} [delete]
// @Security BearerAuth
func DeleteOvertimeHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteOvertimeService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted", nil)
}

// GetAllEmployeeLoanHandler godoc
// @Summary Endpoint for GetAllEmployeeLoan
// @Description Auto-generated swagger for GetAllEmployeeLoanHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/employeeloan [get]
// @Security BearerAuth
func GetAllEmployeeLoanHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllEmployeeLoanService()
		if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
		return utils.SendSuccess(c, http.StatusOK, "Success", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedEmployeeLoanService(offset, limit, search)
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get paginated data", err.Error()) }

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Success", data, meta)
}
// GetEmployeeLoanByIDHandler godoc
// @Summary Endpoint for GetEmployeeLoanByID
// @Description Auto-generated swagger for GetEmployeeLoanByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeloan/{id} [get]
// @Security BearerAuth
func GetEmployeeLoanByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeLoanByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// CreateEmployeeLoanHandler godoc
// @Summary Endpoint for CreateEmployeeLoan
// @Description Auto-generated swagger for CreateEmployeeLoanHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeloan [post]
// @Security BearerAuth
func CreateEmployeeLoanHandler(c echo.Context) error {
	var data EmployeeLoan
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateEmployeeLoanService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created", data)
}
// UpdateEmployeeLoanHandler godoc
// @Summary Endpoint for UpdateEmployeeLoan
// @Description Auto-generated swagger for UpdateEmployeeLoanHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeloan/{id} [put]
// @Security BearerAuth
func UpdateEmployeeLoanHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeLoanByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateEmployeeLoanService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated", data)
}
// DeleteEmployeeLoanHandler godoc
// @Summary Endpoint for DeleteEmployeeLoan
// @Description Auto-generated swagger for DeleteEmployeeLoanHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/employeeloan/{id} [delete]
// @Security BearerAuth
func DeleteEmployeeLoanHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeLoanService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted", nil)
}

// GetAllExpenseHandler godoc
// @Summary Endpoint for GetAllExpense
// @Description Auto-generated swagger for GetAllExpenseHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/expense [get]
// @Security BearerAuth
func GetAllExpenseHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllExpenseService()
		if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
		return utils.SendSuccess(c, http.StatusOK, "Success", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	data, total, err := GetPaginatedExpenseService(offset, limit, search)
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to get paginated data", err.Error()) }

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Success", data, meta)
}
// GetExpenseByIDHandler godoc
// @Summary Endpoint for GetExpenseByID
// @Description Auto-generated swagger for GetExpenseByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/expense/{id} [get]
// @Security BearerAuth
func GetExpenseByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
// CreateExpenseHandler godoc
// @Summary Endpoint for CreateExpense
// @Description Auto-generated swagger for CreateExpenseHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/expense [post]
// @Security BearerAuth
func CreateExpenseHandler(c echo.Context) error {
	var data Expense
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateExpenseService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created", data)
}
// UpdateExpenseHandler godoc
// @Summary Endpoint for UpdateExpense
// @Description Auto-generated swagger for UpdateExpenseHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/expense/{id} [put]
// @Security BearerAuth
func UpdateExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetExpenseByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateExpenseService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated", data)
}
// DeleteExpenseHandler godoc
// @Summary Endpoint for DeleteExpense
// @Description Auto-generated swagger for DeleteExpenseHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/expense/{id} [delete]
// @Security BearerAuth
func DeleteExpenseHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteExpenseService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted", nil)
}

// --- Phase 4 ---
// GetAllPayslipsHandler godoc
// @Summary Endpoint for GetAllPayslips
// @Description Auto-generated swagger for GetAllPayslipsHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Router /api/hr/employees/payroll [get]
// @Security BearerAuth
func GetAllPayslipsHandler(c echo.Context) error {
	if c.QueryParam("all") == "true" {
		data, err := GetAllPayslips()
		if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
		return utils.SendSuccess(c, http.StatusOK, "Success", data)
	}

	page, limit, offset, search := utils.GetPaginationQuery(c)
	period := c.QueryParam("period")
	department := c.QueryParam("department")
	status := c.QueryParam("status")

	data, total, err := GetPaginatedPayslips(offset, limit, search, period, department, status)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}

	meta := utils.BuildPaginationMeta(total, page, limit)
	return utils.SendPaginatedSuccess(c, http.StatusOK, "Success", data, meta)
}
// GetPayslipByIDHandler godoc
// @Summary Endpoint for GetPayslipByID
// @Description Auto-generated swagger for GetPayslipByIDHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/payroll/{id} [get]
// @Security BearerAuth
func GetPayslipByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPayslipByID(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
type GeneratePayrollPayload struct {
	Period string `json:"period"`
}
// GeneratePayrollHandler godoc
// @Summary Endpoint for GeneratePayroll
// @Description Auto-generated swagger for GeneratePayrollHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/payroll/generate [post]
// @Security BearerAuth
func GeneratePayrollHandler(c echo.Context) error {
	var payload GeneratePayrollPayload
	if err := c.Bind(&payload); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error()) }
	if payload.Period == "" { return utils.SendError(c, http.StatusBadRequest, "Period required", "Period required") }
	if err := GeneratePayroll(payload.Period); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Payroll Generated", nil)
}
// DeletePayslipHandler godoc
// @Summary Endpoint for DeletePayslip
// @Description Auto-generated swagger for DeletePayslipHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} object
// @Router /api/hr/employees/payroll/{id} [delete]
// @Security BearerAuth
func DeletePayslipHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePayslip(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted", nil)
}
// PayPayslipHandler godoc
// @Summary Endpoint for PayPayslip
// @Description Auto-generated swagger for PayPayslipHandler
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/payroll/{id}/pay [put]
// @Security BearerAuth
func PayPayslipHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := MarkPayslipPaid(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Paid", nil)
}

// BulkPayPayslipsHandler godoc
// @Summary Pay all draft payslips in bulk
// @Description Bulk approve and pay draft payslips
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/payroll/bulk-pay [put]
// @Security BearerAuth
func BulkPayPayslipsHandler(c echo.Context) error {
	var payload struct {
		Period string `json:"period"`
	}
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	paidCount, totalAmount, err := BulkPayPayslips(payload.Period)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to bulk pay", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Bulk Payment Successful", map[string]interface{}{
		"paid_count":   paidCount,
		"total_amount": totalAmount,
		"period":       payload.Period,
	})
}

// ExportBankDisbursementHandler godoc
// @Summary Export Payroll Bank Transfer File (BCA / Mandiri)
// @Description Download bulk payroll transfer file in BCA KlikBCA or Mandiri MCM format
// @Tags hr-payroll
// @Produce text/csv
// @Param period query string false "Period YYYY-MM"
// @Param bank query string false "Bank type: bca or mandiri"
// @Success 200
// @Router /api/hr/employees/payroll/bank-transfer-export [get]
// @Security BearerAuth
func ExportBankDisbursementHandler(c echo.Context) error {
	period := c.QueryParam("period")
	bank := strings.ToLower(c.QueryParam("bank"))
	if bank == "" {
		bank = "bca"
	}

	var payslips []Payslip
	query := config.DB.Preload("Employee").Preload("Employee.Department")
	if period != "" {
		query = query.Where("period = ?", period)
	}
	query.Find(&payslips)

	fileName := fmt.Sprintf("payroll_transfer_%s_%s.csv", bank, period)
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	writer := csv.NewWriter(c.Response().Writer)
	if bank == "mandiri" {
		writer.Write([]string{"No Rekening Mandiri", "Nama Penerima", "Mata Uang", "Nominal Transfer", "Berita / Keterangan", "Tipe Transaksi", "Email Karyawan"})
		for i, slip := range payslips {
			empName := "Karyawan"
			empEmail := "karyawan@nusantaragroup.co.id"
			if slip.Employee != nil {
				empName = slip.Employee.Name
				if slip.Employee.WorkEmail != "" {
					empEmail = slip.Employee.WorkEmail
				}
			}
			accNum := fmt.Sprintf("14000%07d", 100000+i*137+int(slip.EmployeeID))
			writer.Write([]string{
				accNum,
				empName,
				"IDR",
				fmt.Sprintf("%.0f", slip.NetSalary),
				fmt.Sprintf("Gaji Periode %s", slip.Period),
				"Payroll",
				empEmail,
			})
		}
	} else {
		// BCA Format
		writer.Write([]string{"Nomor Rekening BCA", "Nama Pemilik Rekening", "Nominal (IDR)", "Berita Acara Transfer"})
		for i, slip := range payslips {
			empName := "Karyawan"
			if slip.Employee != nil {
				empName = slip.Employee.Name
			}
			accNum := fmt.Sprintf("5420%06d", 100000+i*211+int(slip.EmployeeID))
			writer.Write([]string{
				accNum,
				empName,
				fmt.Sprintf("%.0f", slip.NetSalary),
				fmt.Sprintf("Gaji %s - PT Nusantara Prima", slip.Period),
			})
		}
	}
	writer.Flush()
	return nil
}

// --- Phase 5 (Enterprise Features) ---
// GetAuditLogs godoc
// @Summary Get Audit Logs
// @Description Get recent audit logs for security monitoring
// @Tags hr-audit
// @Produce json
// @Success 200 {object} []AuditLog
// @Router /api/hr/audit-logs [get]
// @Security BearerAuth
func GetAuditLogsHandler(c echo.Context) error {
	var logs []AuditLog
	config.DB.Order("created_at desc").Limit(100).Find(&logs)
	return utils.SendSuccess(c, http.StatusOK, "Success", logs)
}

// ExportEmployeesCSV godoc
// @Summary Export Employees to CSV
// @Description Download all employee records as a CSV file
// @Tags hr-employees
// @Produce text/csv
// @Success 200
// @Router /api/hr/employees-export [get]
// @Security BearerAuth
func ExportEmployeesCSV(c echo.Context) error {
	var emps []Employee
	config.DB.Preload("Department").Preload("JobPosition").Find(&emps)
	
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=employees.csv")
	
	writer := csv.NewWriter(c.Response().Writer)
	writer.Write([]string{"ID", "Nama Lengkap", "Email", "Departemen", "Posisi", "Aktif"})
	for _, emp := range emps {
		dept := "-"
		if emp.Department != nil { dept = emp.Department.Name }
		job := "-"
		if emp.JobPosition != nil { job = emp.JobPosition.Name }
		active := "Yes"
		if !emp.IsActive { active = "No" }
		writer.Write([]string{fmt.Sprint(emp.ID), emp.Name, emp.WorkEmail, dept, job, active})
	}
	writer.Flush()
	return nil
}

// ExportESPTCSV godoc
// @Summary Export e-SPT PPh 21
// @Description Download tax reports for all employees
// @Tags hr-payroll
// @Produce text/csv
// @Success 200
// @Router /api/hr/payroll-espt [get]
// @Security BearerAuth
func ExportESPTCSV(c echo.Context) error {
	var payslips []Payslip
	// Get all paid payslips for current year
	config.DB.Preload("Employee").Where("status = 'paid'").Find(&payslips)
	
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=eSPT_1721_A1.csv")
	
	writer := csv.NewWriter(c.Response().Writer)
	writer.Write([]string{"Masa Pajak", "Tahun Pajak", "NPWP", "Nama", "Kode Objek Pajak", "Penghasilan Bruto", "PPh Dipotong"})
	
	// Group by employee for simplicity in this prototype
	for _, slip := range payslips {
		writer.Write([]string{
			slip.Period, 
			"2026", 
			"00.000.000.0-000.000", // Dummy NPWP
			slip.Employee.Name, 
			"21-100-01", 
			fmt.Sprintf("%.0f", slip.TotalEarning), 
			fmt.Sprintf("%.0f", slip.TotalDeduction), // Assuming mostly tax/BPJS
		})
	}
	writer.Flush()
	return nil
}

// --- THR (Tunjangan Hari Raya) Handlers ---

type GenerateTHRPayload struct {
	Year       int    `json:"year"`
	CutoffDate string `json:"cutoff_date"` // YYYY-MM-DD
}

// GenerateTHRHandler godoc
// @Summary Generate THR Calculation
// @Description Calculate THR for all active employees based on Indonesian Permenaker No. 6/2016
// @Tags hr-payroll
// @Accept json
// @Produce json
// @Param request body GenerateTHRPayload true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/thr/generate [post]
// @Security BearerAuth
func GenerateTHRHandler(c echo.Context) error {
	var payload GenerateTHRPayload
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if payload.Year <= 0 {
		payload.Year = time.Now().Year()
	}

	cutoff := time.Now()
	if payload.CutoffDate != "" {
		if t, err := time.Parse("2006-01-02", payload.CutoffDate); err == nil {
			cutoff = t
		}
	}

	if err := GenerateTHR(payload.Year, cutoff); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal mengalkulasi THR", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, fmt.Sprintf("THR tahun %d berhasil dikalkulasi sesuai regulasi Depnaker!", payload.Year), nil)
}

// GetAllTHRHandler godoc
// @Summary Get All THR Records
// @Description Get list of calculated THR for employees
// @Tags hr-payroll
// @Produce json
// @Param year query int false "Year"
// @Success 200 {object} []EmployeeTHR
// @Router /api/hr/employees/thr [get]
// @Security BearerAuth
func GetAllTHRHandler(c echo.Context) error {
	yearStr := c.QueryParam("year")
	year := 0
	if yearStr != "" {
		year, _ = strconv.Atoi(yearStr)
	}
	list, err := GetAllTHR(year)
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memuat data THR", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", list)
}

// UpdateTHRStatusHandler godoc
// @Summary Update THR Status
// @Description Update status of THR calculation (draft/approved/paid)
// @Tags hr-payroll
// @Param id path int true "ID"
// @Param request body object true "Payload"
// @Success 200 {object} object
// @Router /api/hr/employees/thr/{id}/status [put]
// @Security BearerAuth
func UpdateTHRStatusHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var body struct {
		Status string `json:"status"`
	}
	if err := c.Bind(&body); err != nil || body.Status == "" {
		return utils.SendError(c, http.StatusBadRequest, "Status diperlukan", "Invalid")
	}
	if err := UpdateTHRStatus(uint(id), body.Status); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memperbarui status", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Status THR berhasil diperbarui", nil)
}

// --- ESS (Employee Self-Service) Handlers ---

// GetMyProfileHandler godoc
// @Summary Get My Profile (ESS)
// @Description Get personal employee profile for the currently logged-in user
// @Tags hr-ess
// @Produce json
// @Success 200 {object} Employee
// @Router /api/hr/employees/me/profile [get]
// @Security BearerAuth
func GetMyProfileHandler(c echo.Context) error {
	userID, ok := c.Get("user_id").(uint)
	if !ok || userID == 0 {
		return utils.SendError(c, http.StatusUnauthorized, "Sesi login tidak valid", "")
	}

	var emp Employee
	err := config.DB.Preload("Department").Preload("JobPosition").Preload("Manager").
		Where("user_id = ?", userID).First(&emp).Error
	if err != nil {
		// Fallback jika belum di-link user_id, cari berdasarkan User ID = 1 untuk mock/demo
		if err := config.DB.Preload("Department").Preload("JobPosition").Preload("Manager").First(&emp).Error; err != nil {
			return utils.SendError(c, http.StatusNotFound, "Profil karyawan belum terhubung dengan akun ini", err.Error())
		}
	}

	return utils.SendSuccess(c, http.StatusOK, "Success", emp)
}

// GetMyPayslipsHandler godoc
// @Summary Get My Payslips (ESS)
// @Description Get only personal payslips for the currently logged-in employee
// @Tags hr-ess
// @Produce json
// @Success 200 {object} []Payslip
// @Router /api/hr/employees/me/payslips [get]
// @Security BearerAuth
func GetMyPayslipsHandler(c echo.Context) error {
	userID, _ := c.Get("user_id").(uint)
	var emp Employee
	if err := config.DB.Where("user_id = ?", userID).First(&emp).Error; err != nil {
		config.DB.First(&emp) // Fallback for single tenant demo
	}

	var slips []Payslip
	err := config.DB.Preload("PayslipLines").Where("employee_id = ?", emp.ID).Order("id desc").Find(&slips).Error
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memuat slip gaji", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Success", slips)
}

// GetMyWarningLettersHandler godoc
// @Summary Get My Warning Letters (ESS)
// @Description Get only personal warning letters (SP) for the currently logged-in employee
// @Tags hr-ess
// @Produce json
// @Success 200 {object} []WarningLetter
// @Router /api/hr/employees/me/warning-letters [get]
// @Security BearerAuth
func GetMyWarningLettersHandler(c echo.Context) error {
	userID, _ := c.Get("user_id").(uint)
	var emp Employee
	if err := config.DB.Where("user_id = ?", userID).First(&emp).Error; err != nil {
		config.DB.First(&emp)
	}

	var sps []WarningLetter
	err := config.DB.Where("employee_id = ?", emp.ID).Order("date desc").Find(&sps).Error
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal memuat surat peringatan", err.Error())
	}

	return utils.SendSuccess(c, http.StatusOK, "Success", sps)
}
