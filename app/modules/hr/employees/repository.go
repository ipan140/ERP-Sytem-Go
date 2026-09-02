package employees

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateEmployee(data *Employee) error {
	return config.DB.Create(data).Error
}

func GetAllEmployee() ([]Employee, error) {
	var list []Employee
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetEmployeeByID(id uint) (*Employee, error) {
	var data Employee
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateEmployee(data *Employee) error {
	return config.DB.Save(data).Error
}

func DeleteEmployee(id uint) error {
	return config.DB.Delete(&Employee{}, id).Error
}

func CreateJobPosition(data *JobPosition) error { return config.DB.Create(data).Error }
func GetAllJobPosition() ([]JobPosition, error) {
	var list []JobPosition
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetJobPositionByID(id uint) (*JobPosition, error) {
	var data JobPosition
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateJobPosition(data *JobPosition) error { return config.DB.Save(data).Error }
func DeleteJobPosition(id uint) error           { return config.DB.Delete(&JobPosition{}, id).Error }

func CreateWorkingSchedule(data *WorkingSchedule) error { return config.DB.Create(data).Error }
func GetAllWorkingSchedule() ([]WorkingSchedule, error) {
	var list []WorkingSchedule
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetWorkingScheduleByID(id uint) (*WorkingSchedule, error) {
	var data WorkingSchedule
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateWorkingSchedule(data *WorkingSchedule) error { return config.DB.Save(data).Error }
func DeleteWorkingSchedule(id uint) error               { return config.DB.Delete(&WorkingSchedule{}, id).Error }

func CreateContract(data *Contract) error { return config.DB.Create(data).Error }
func GetAllContract() ([]Contract, error) {
	var list []Contract
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetContractByID(id uint) (*Contract, error) {
	var data Contract
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateContract(data *Contract) error { return config.DB.Save(data).Error }
func DeleteContract(id uint) error        { return config.DB.Delete(&Contract{}, id).Error }

func CreateSkill(data *Skill) error { return config.DB.Create(data).Error }
func GetAllSkill() ([]Skill, error) {
	var list []Skill
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSkillByID(id uint) (*Skill, error) {
	var data Skill
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSkill(data *Skill) error { return config.DB.Save(data).Error }
func DeleteSkill(id uint) error     { return config.DB.Delete(&Skill{}, id).Error }

func CreateSkillLevel(data *SkillLevel) error { return config.DB.Create(data).Error }
func GetAllSkillLevel() ([]SkillLevel, error) {
	var list []SkillLevel
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSkillLevelByID(id uint) (*SkillLevel, error) {
	var data SkillLevel
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSkillLevel(data *SkillLevel) error { return config.DB.Save(data).Error }
func DeleteSkillLevel(id uint) error          { return config.DB.Delete(&SkillLevel{}, id).Error }

func CreateEmployeeSkill(data *EmployeeSkill) error { return config.DB.Create(data).Error }
func GetAllEmployeeSkill() ([]EmployeeSkill, error) {
	var list []EmployeeSkill
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetEmployeeSkillByID(id uint) (*EmployeeSkill, error) {
	var data EmployeeSkill
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateEmployeeSkill(data *EmployeeSkill) error { return config.DB.Save(data).Error }
func DeleteEmployeeSkill(id uint) error             { return config.DB.Delete(&EmployeeSkill{}, id).Error }

func CreateResumeLine(data *ResumeLine) error { return config.DB.Create(data).Error }
func GetAllResumeLine() ([]ResumeLine, error) {
	var list []ResumeLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetResumeLineByID(id uint) (*ResumeLine, error) {
	var data ResumeLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateResumeLine(data *ResumeLine) error { return config.DB.Save(data).Error }
func DeleteResumeLine(id uint) error          { return config.DB.Delete(&ResumeLine{}, id).Error }

func CreateDepartment(data *Department) error { return config.DB.Create(data).Error }
func GetAllDepartment() ([]Department, error) {
	var list []Department
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetDepartmentByID(id uint) (*Department, error) {
	var data Department
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateDepartment(data *Department) error { return config.DB.Save(data).Error }
func DeleteDepartment(id uint) error          { return config.DB.Delete(&Department{}, id).Error }


// --- Warning Letter ---
func GetAllWarningLetter() ([]WarningLetter, error) {
	var list []WarningLetter
	err := config.DB.Preload("Employee").Find(&list).Error
	return list, err
}
func GetWarningLetterByID(id uint) (*WarningLetter, error) {
	var data WarningLetter
	err := config.DB.Preload("Employee").First(&data, id).Error
	return &data, err
}
func CreateWarningLetter(data *WarningLetter) error {
	return config.DB.Create(data).Error
}
func UpdateWarningLetter(data *WarningLetter) error {
	return config.DB.Save(data).Error
}
func DeleteWarningLetter(id uint) error {
	return config.DB.Delete(&WarningLetter{}, id).Error
}


// --- Employee Task ---
func GetAllEmployeeTask() ([]EmployeeTask, error) {
	var list []EmployeeTask
	err := config.DB.Preload("Employee").Find(&list).Error
	return list, err
}
func GetEmployeeTaskByID(id uint) (*EmployeeTask, error) {
	var data EmployeeTask
	err := config.DB.Preload("Employee").First(&data, id).Error
	return &data, err
}
func CreateEmployeeTask(data *EmployeeTask) error {
	return config.DB.Create(data).Error
}
func UpdateEmployeeTask(data *EmployeeTask) error {
	return config.DB.Save(data).Error
}
func DeleteEmployeeTask(id uint) error {
	return config.DB.Delete(&EmployeeTask{}, id).Error
}

// --- Phase 3 ---
func GetAllOvertime() ([]Overtime, error) {
	var list []Overtime
	err := config.DB.Preload("Employee").Find(&list).Error
	return list, err
}
func GetOvertimeByID(id uint) (*Overtime, error) {
	var data Overtime
	err := config.DB.Preload("Employee").First(&data, id).Error
	return &data, err
}
func CreateOvertime(data *Overtime) error { return config.DB.Create(data).Error }
func UpdateOvertime(data *Overtime) error { return config.DB.Save(data).Error }
func DeleteOvertime(id uint) error { return config.DB.Delete(&Overtime{}, id).Error }

func GetAllEmployeeLoan() ([]EmployeeLoan, error) {
	var list []EmployeeLoan
	err := config.DB.Preload("Employee").Find(&list).Error
	return list, err
}
func GetEmployeeLoanByID(id uint) (*EmployeeLoan, error) {
	var data EmployeeLoan
	err := config.DB.Preload("Employee").First(&data, id).Error
	return &data, err
}
func CreateEmployeeLoan(data *EmployeeLoan) error { return config.DB.Create(data).Error }
func UpdateEmployeeLoan(data *EmployeeLoan) error { return config.DB.Save(data).Error }
func DeleteEmployeeLoan(id uint) error { return config.DB.Delete(&EmployeeLoan{}, id).Error }

func GetAllExpense() ([]Expense, error) {
	var list []Expense
	err := config.DB.Preload("Employee").Find(&list).Error
	return list, err
}
func GetExpenseByID(id uint) (*Expense, error) {
	var data Expense
	err := config.DB.Preload("Employee").First(&data, id).Error
	return &data, err
}
func CreateExpense(data *Expense) error { return config.DB.Create(data).Error }
func UpdateExpense(data *Expense) error { return config.DB.Save(data).Error }
func DeleteExpense(id uint) error { return config.DB.Delete(&Expense{}, id).Error }

// --- Phase 4: Payroll ---
func GetAllPayslips() ([]Payslip, error) {
	var list []Payslip
	err := config.DB.Preload("Employee").Preload("PayslipLines").Order("created_at desc").Find(&list).Error
	return list, err
}
func GetPayslipByID(id uint) (*Payslip, error) {
	var data Payslip
	err := config.DB.Preload("Employee.Department").Preload("Employee.JobPosition").Preload("PayslipLines").First(&data, id).Error
	return &data, err
}
func GeneratePayroll(period string) error {
	var employees []Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		return err
	}

	for _, emp := range employees {
		// 1. Get Basic Salary from first contract
		var basicSalary float64 = 0
		var contract Contract
		if err := config.DB.Where("employee_id = ?", emp.ID).Order("created_at desc").First(&contract).Error; err == nil {
			basicSalary = contract.Wage
		}
		if basicSalary == 0 {
			continue // Skip employees without a contract/salary
		}

		// Check if payslip already exists
		var existingCount int64
		config.DB.Model(&Payslip{}).Where("employee_id = ? AND period = ?", emp.ID, period).Count(&existingCount)
		if existingCount > 0 {
			continue // Skip if already generated
		}

		payslip := Payslip{
			EmployeeID:  emp.ID,
			Period:      period,
			BasicSalary: basicSalary,
			Status:      "draft",
		}

		var lines []PayslipLine
		var totalEarning = basicSalary
		var totalDeduction float64 = 0

		// Basic Salary Line
		lines = append(lines, PayslipLine{Category: "earning", Name: "Gaji Pokok", Amount: basicSalary})

		// 2. Overtime (Mock logic: assume 50k per hour for approved overtime)
		var overtimes []Overtime
		config.DB.Where("employee_id = ? AND status = 'approved'", emp.ID).Find(&overtimes)
		var otAmount float64 = 0
		for _, ot := range overtimes {
			otAmount += ot.Hours * 50000 // 50k / hour rate
		}
		if otAmount > 0 {
			lines = append(lines, PayslipLine{Category: "earning", Name: "Uang Lembur (Overtime)", Amount: otAmount})
			totalEarning += otAmount
		}

		// 3. Employee Loans (Deduct monthly installment)
		var loans []EmployeeLoan
		config.DB.Where("employee_id = ? AND status = 'approved'", emp.ID).Find(&loans)
		for _, loan := range loans {
			if loan.MonthlyInstallment > 0 {
				lines = append(lines, PayslipLine{Category: "deduction", Name: "Potongan Kasbon / Pinjaman", Amount: loan.MonthlyInstallment})
				totalDeduction += loan.MonthlyInstallment
			}
		}

		// 4. BPJS & Pajak
		bpjsKes := basicSalary * 0.01 // 1%
		bpjsTk := basicSalary * 0.03  // 3% (JHT 2% + JP 1%)
		lines = append(lines, PayslipLine{Category: "deduction", Name: "BPJS Kesehatan (1%)", Amount: bpjsKes})
		lines = append(lines, PayslipLine{Category: "deduction", Name: "BPJS Ketenagakerjaan (3%)", Amount: bpjsTk})
		totalDeduction += (bpjsKes + bpjsTk)

		// Simple PPh 21 TER (flat 5% on earnings above 5jt)
		taxable := totalEarning - 5000000
		if taxable > 0 {
			pph21 := taxable * 0.05
			lines = append(lines, PayslipLine{Category: "deduction", Name: "PPh 21 (Pajak Penghasilan)", Amount: pph21})
			totalDeduction += pph21
		}

		payslip.TotalEarning = totalEarning
		payslip.TotalDeduction = totalDeduction
		payslip.NetSalary = totalEarning - totalDeduction
		payslip.PayslipLines = lines

		config.DB.Create(&payslip)
	}
	return nil
}
func DeletePayslip(id uint) error { return config.DB.Delete(&Payslip{}, id).Error }
func MarkPayslipPaid(id uint) error { return config.DB.Model(&Payslip{}).Where("id = ?", id).Update("status", "paid").Error }
