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

func CreateJobPositionHandler(c echo.Context) error {
	var data JobPosition
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateJobPositionService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllJobPositionHandler(c echo.Context) error {
	data, err := GetAllJobPositionService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetJobPositionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobPositionByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobPositionByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateJobPositionService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJobPositionService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateWorkingScheduleHandler(c echo.Context) error {
	var data WorkingSchedule
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateWorkingScheduleService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllWorkingScheduleHandler(c echo.Context) error {
	data, err := GetAllWorkingScheduleService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetWorkingScheduleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkingScheduleByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkingScheduleByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateWorkingScheduleService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWorkingScheduleService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateContractHandler(c echo.Context) error {
	var data Contract
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateContractService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllContractHandler(c echo.Context) error {
	data, err := GetAllContractService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetContractByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetContractByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetContractByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateContractService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteContractService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateSkillHandler(c echo.Context) error {
	var data Skill
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSkillService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllSkillHandler(c echo.Context) error {
	data, err := GetAllSkillService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSkillByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSkillService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateSkillLevelHandler(c echo.Context) error {
	var data SkillLevel
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSkillLevelService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllSkillLevelHandler(c echo.Context) error {
	data, err := GetAllSkillLevelService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetSkillLevelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillLevelByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillLevelByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSkillLevelService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillLevelService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateEmployeeSkillHandler(c echo.Context) error {
	var data EmployeeSkill
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateEmployeeSkillService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllEmployeeSkillHandler(c echo.Context) error {
	data, err := GetAllEmployeeSkillService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetEmployeeSkillByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateEmployeeSkillService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeSkillService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

func CreateResumeLineHandler(c echo.Context) error {
	var data ResumeLine
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateResumeLineService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
func GetAllResumeLineHandler(c echo.Context) error {
	data, err := GetAllResumeLineService()
	if err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func GetResumeLineByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResumeLineByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Retrieved successfully", data)
}
func UpdateResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResumeLineByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateResumeLineService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Updated successfully", data)
}
func DeleteResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteResumeLineService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
