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

// @Summary Create JobPosition
// @Description Create a new JobPosition
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/jobposition [post]
// @Security BearerAuth
func CreateJobPositionHandler(c echo.Context) error {
	var data JobPosition
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateJobPositionService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all JobPosition
// @Description Retrieve a list of all JobPosition
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/jobposition [get]
// @Security BearerAuth
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
// @Summary Update JobPosition
// @Description Update an existing JobPosition
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "JobPosition ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/jobposition/{id} [put]
// @Security BearerAuth
func UpdateJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJobPositionByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateJobPositionService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteJobPositionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJobPositionService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create WorkingSchedule
// @Description Create a new WorkingSchedule
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/workingschedule [post]
// @Security BearerAuth
func CreateWorkingScheduleHandler(c echo.Context) error {
	var data WorkingSchedule
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateWorkingScheduleService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all WorkingSchedule
// @Description Retrieve a list of all WorkingSchedule
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/workingschedule [get]
// @Security BearerAuth
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
// @Summary Update WorkingSchedule
// @Description Update an existing WorkingSchedule
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "WorkingSchedule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/workingschedule/{id} [put]
// @Security BearerAuth
func UpdateWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetWorkingScheduleByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateWorkingScheduleService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteWorkingScheduleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteWorkingScheduleService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Contract
// @Description Create a new Contract
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/contract [post]
// @Security BearerAuth
func CreateContractHandler(c echo.Context) error {
	var data Contract
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateContractService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all Contract
// @Description Retrieve a list of all Contract
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/contract [get]
// @Security BearerAuth
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
// @Summary Update Contract
// @Description Update an existing Contract
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Contract ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/contract/{id} [put]
// @Security BearerAuth
func UpdateContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetContractByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateContractService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteContractHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteContractService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create Skill
// @Description Create a new Skill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/skill [post]
// @Security BearerAuth
func CreateSkillHandler(c echo.Context) error {
	var data Skill
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSkillService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all Skill
// @Description Retrieve a list of all Skill
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skill [get]
// @Security BearerAuth
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
// @Summary Update Skill
// @Description Update an existing Skill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "Skill ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skill/{id} [put]
// @Security BearerAuth
func UpdateSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSkillService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create SkillLevel
// @Description Create a new SkillLevel
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/skilllevel [post]
// @Security BearerAuth
func CreateSkillLevelHandler(c echo.Context) error {
	var data SkillLevel
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateSkillLevelService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all SkillLevel
// @Description Retrieve a list of all SkillLevel
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skilllevel [get]
// @Security BearerAuth
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
// @Summary Update SkillLevel
// @Description Update an existing SkillLevel
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "SkillLevel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/skilllevel/{id} [put]
// @Security BearerAuth
func UpdateSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetSkillLevelByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateSkillLevelService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteSkillLevelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteSkillLevelService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create EmployeeSkill
// @Description Create a new EmployeeSkill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/employeeskill [post]
// @Security BearerAuth
func CreateEmployeeSkillHandler(c echo.Context) error {
	var data EmployeeSkill
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateEmployeeSkillService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all EmployeeSkill
// @Description Retrieve a list of all EmployeeSkill
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/employeeskill [get]
// @Security BearerAuth
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
// @Summary Update EmployeeSkill
// @Description Update an existing EmployeeSkill
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "EmployeeSkill ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/employeeskill/{id} [put]
// @Security BearerAuth
func UpdateEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetEmployeeSkillByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateEmployeeSkillService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteEmployeeSkillHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteEmployeeSkillService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}

// @Summary Create ResumeLine
// @Description Create a new ResumeLine
// @Tags hr-employees
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /api/hr/employees/resumeline [post]
// @Security BearerAuth
func CreateResumeLineHandler(c echo.Context) error {
	var data ResumeLine
	if err := c.Bind(&data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := CreateResumeLineService(&data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to create", err.Error()) }
	return utils.SendSuccess(c, http.StatusCreated, "Created successfully", data)
}
// @Summary Get all ResumeLine
// @Description Retrieve a list of all ResumeLine
// @Tags hr-employees
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/resumeline [get]
// @Security BearerAuth
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
// @Summary Update ResumeLine
// @Description Update an existing ResumeLine
// @Tags hr-employees
// @Accept json
// @Produce json
// @Param id path int true "ResumeLine ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/hr/employees/resumeline/{id} [put]
// @Security BearerAuth
func UpdateResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetResumeLineByIDService(uint(id))
	if err != nil { return utils.SendError(c, http.StatusNotFound, "Not found", err.Error()) }
	if err := c.Bind(data); err != nil { return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error()) }
	if err := UpdateResumeLineService(data); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to update", err.Error()) }
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
func DeleteResumeLineHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteResumeLineService(uint(id)); err != nil { return utils.SendError(c, http.StatusInternalServerError, "Failed to delete", err.Error()) }
	return utils.SendSuccess(c, http.StatusOK, "Deleted successfully", nil)
}
