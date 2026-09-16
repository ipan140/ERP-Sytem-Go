package employees

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

func CreateEmployeeService(data *Employee) error {
	return CreateEmployee(data)
}

func GetAllEmployeeService() ([]Employee, error) {
	return GetAllEmployee()
}

func GetPaginatedEmployeeService(offset int, limit int, search string, departmentID string, isActive string) ([]Employee, int64, error) {
	return GetPaginatedEmployees(offset, limit, search, departmentID, isActive)
}

func GetEmployeeByIDService(id uint) (*Employee, error) {
	return GetEmployeeByID(id)
}

func UpdateEmployeeService(data *Employee) error {
	return UpdateEmployee(data)
}

func DeleteEmployeeService(id uint) error {
	return DeleteEmployee(id)
}

func CreateJobPositionService(data *JobPosition) error        { return CreateJobPosition(data) }
func GetAllJobPositionService() ([]JobPosition, error)        { return GetAllJobPosition() }
func GetJobPositionByIDService(id uint) (*JobPosition, error) { return GetJobPositionByID(id) }
func UpdateJobPositionService(data *JobPosition) error        { return UpdateJobPosition(data) }
func DeleteJobPositionService(id uint) error                  { return DeleteJobPosition(id) }

func CreateWorkingScheduleService(data *WorkingSchedule) error { return CreateWorkingSchedule(data) }
func GetAllWorkingScheduleService() ([]WorkingSchedule, error) { return GetAllWorkingSchedule() }
func GetWorkingScheduleByIDService(id uint) (*WorkingSchedule, error) {
	return GetWorkingScheduleByID(id)
}
func UpdateWorkingScheduleService(data *WorkingSchedule) error { return UpdateWorkingSchedule(data) }
func DeleteWorkingScheduleService(id uint) error               { return DeleteWorkingSchedule(id) }

func CreateContractService(data *Contract) error        { return CreateContract(data) }
func GetAllContractService() ([]Contract, error)        { return GetAllContract() }
func GetPaginatedContractService(offset, limit int, search string) ([]Contract, int64, error) {
	return GetPaginatedContract(offset, limit, search)
}
func GetContractByIDService(id uint) (*Contract, error) { return GetContractByID(id) }
func UpdateContractService(data *Contract) error        { return UpdateContract(data) }
func DeleteContractService(id uint) error               { return DeleteContract(id) }

func CreateSkillService(data *Skill) error        { return CreateSkill(data) }
func GetAllSkillService() ([]Skill, error)        { return GetAllSkill() }
func GetSkillByIDService(id uint) (*Skill, error) { return GetSkillByID(id) }
func UpdateSkillService(data *Skill) error        { return UpdateSkill(data) }
func DeleteSkillService(id uint) error            { return DeleteSkill(id) }

func CreateSkillLevelService(data *SkillLevel) error        { return CreateSkillLevel(data) }
func GetAllSkillLevelService() ([]SkillLevel, error)        { return GetAllSkillLevel() }
func GetSkillLevelByIDService(id uint) (*SkillLevel, error) { return GetSkillLevelByID(id) }
func UpdateSkillLevelService(data *SkillLevel) error        { return UpdateSkillLevel(data) }
func DeleteSkillLevelService(id uint) error                 { return DeleteSkillLevel(id) }

func CreateEmployeeSkillService(data *EmployeeSkill) error        { return CreateEmployeeSkill(data) }
func GetAllEmployeeSkillService() ([]EmployeeSkill, error)        { return GetAllEmployeeSkill() }
func GetEmployeeSkillByIDService(id uint) (*EmployeeSkill, error) { return GetEmployeeSkillByID(id) }
func UpdateEmployeeSkillService(data *EmployeeSkill) error        { return UpdateEmployeeSkill(data) }
func DeleteEmployeeSkillService(id uint) error                    { return DeleteEmployeeSkill(id) }

func CreateResumeLineService(data *ResumeLine) error        { return CreateResumeLine(data) }
func GetAllResumeLineService() ([]ResumeLine, error)        { return GetAllResumeLine() }
func GetResumeLineByIDService(id uint) (*ResumeLine, error) { return GetResumeLineByID(id) }
func UpdateResumeLineService(data *ResumeLine) error        { return UpdateResumeLine(data) }
func DeleteResumeLineService(id uint) error                 { return DeleteResumeLine(id) }

func ImportEmployeesExcelService(filePath string, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"action": "import_employees", "file_path": filePath, "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "core_excel_import", body)
		log.Printf("👥 Event RabbitMQ: Import Excel Karyawan %s dikirim ke antrean!", filePath)
	}
	return nil
}

func CreateDepartmentService(data *Department) error { return CreateDepartment(data) }
func GetAllDepartmentService() ([]Department, error) { return GetAllDepartment() }
func GetDepartmentByIDService(id uint) (*Department, error) { return GetDepartmentByID(id) }
func UpdateDepartmentService(data *Department) error { return UpdateDepartment(data) }
func DeleteDepartmentService(id uint) error { return DeleteDepartment(id) }


// --- Warning Letter ---
func GetAllWarningLetterService() ([]WarningLetter, error) { return GetAllWarningLetter() }
func GetPaginatedWarningLetterService(offset, limit int, search string) ([]WarningLetter, int64, error) {
	return GetPaginatedWarningLetter(offset, limit, search)
}
func GetWarningLetterByIDService(id uint) (*WarningLetter, error) { return GetWarningLetterByID(id) }
func CreateWarningLetterService(data *WarningLetter) error { return CreateWarningLetter(data) }
func UpdateWarningLetterService(data *WarningLetter) error { return UpdateWarningLetter(data) }
func DeleteWarningLetterService(id uint) error { return DeleteWarningLetter(id) }


// --- Employee Task ---
func GetAllEmployeeTaskService() ([]EmployeeTask, error) { return GetAllEmployeeTask() }
func GetEmployeeTaskByIDService(id uint) (*EmployeeTask, error) { return GetEmployeeTaskByID(id) }
func CreateEmployeeTaskService(data *EmployeeTask) error { return CreateEmployeeTask(data) }
func UpdateEmployeeTaskService(data *EmployeeTask) error { return UpdateEmployeeTask(data) }
func DeleteEmployeeTaskService(id uint) error { return DeleteEmployeeTask(id) }

// --- Phase 3 ---
func GetAllOvertimeService() ([]Overtime, error) { return GetAllOvertime() }
func GetPaginatedOvertimeService(offset, limit int, search string) ([]Overtime, int64, error) {
	return GetPaginatedOvertime(offset, limit, search)
}
func GetOvertimeByIDService(id uint) (*Overtime, error) { return GetOvertimeByID(id) }
func CreateOvertimeService(data *Overtime) error { return CreateOvertime(data) }
func UpdateOvertimeService(data *Overtime) error { return UpdateOvertime(data) }
func DeleteOvertimeService(id uint) error { return DeleteOvertime(id) }

func GetAllEmployeeLoanService() ([]EmployeeLoan, error) { return GetAllEmployeeLoan() }
func GetPaginatedEmployeeLoanService(offset, limit int, search string) ([]EmployeeLoan, int64, error) {
	return GetPaginatedEmployeeLoan(offset, limit, search)
}
func GetEmployeeLoanByIDService(id uint) (*EmployeeLoan, error) { return GetEmployeeLoanByID(id) }
func CreateEmployeeLoanService(data *EmployeeLoan) error { return CreateEmployeeLoan(data) }
func UpdateEmployeeLoanService(data *EmployeeLoan) error { return UpdateEmployeeLoan(data) }
func DeleteEmployeeLoanService(id uint) error { return DeleteEmployeeLoan(id) }

func GetAllExpenseService() ([]Expense, error) { return GetAllExpense() }
func GetPaginatedExpenseService(offset, limit int, search string) ([]Expense, int64, error) {
	return GetPaginatedExpense(offset, limit, search)
}
func GetExpenseByIDService(id uint) (*Expense, error) { return GetExpenseByID(id) }
func CreateExpenseService(data *Expense) error { return CreateExpense(data) }
func UpdateExpenseService(data *Expense) error { return UpdateExpense(data) }
func DeleteExpenseService(id uint) error { return DeleteExpense(id) }
