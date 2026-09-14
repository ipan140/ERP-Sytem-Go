package employees

import (
	"fmt"
	"time"
	"gorm.io/gorm"
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

func GetPaginatedEmployees(offset int, limit int, search string, departmentID string, isActive string) ([]Employee, int64, error) {
	var list []Employee
	var total int64

	query := config.DB.Model(&Employee{})

	if isActive == "true" {
		query = query.Where("is_active = ?", true)
	} else if isActive == "false" {
		query = query.Where("is_active = ?", false)
	}

	if departmentID != "" && departmentID != "all" && departmentID != "0" {
		query = query.Where("department_id = ?", departmentID)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("name ILIKE ? OR work_email ILIKE ? OR work_phone ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
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

func GetPaginatedPayslips(offset, limit int, search, period, department, status string) ([]Payslip, int64, error) {
	var list []Payslip
	var total int64

	query := config.DB.Model(&Payslip{}).
		Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.hr_payslips.employee_id").
		Joins("LEFT JOIN hrd.departments ON hrd.departments.id = hrd.employees.department_id")

	if period != "" && period != "all" {
		query = query.Where("hrd.hr_payslips.period = ?", period)
	}

	if status != "" && status != "all" {
		query = query.Where("hrd.hr_payslips.status = ?", status)
	}

	if department != "" && department != "all" {
		query = query.Where("hrd.departments.name ILIKE ?", "%"+department+"%")
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("hrd.employees.name ILIKE ? OR hrd.hr_payslips.period ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload("Employee.Department").Preload("Employee.JobPosition").Preload("PayslipLines").
		Order("hrd.hr_payslips.created_at DESC, hrd.hr_payslips.id DESC").
		Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
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

		// 4. BPJS & Pajak PPh 21 (TER 2024 Resmi: PP 58/2023 & PMK 168/2023)
		bpjsKes := basicSalary * 0.01 // 1%
		bpjsTk := basicSalary * 0.03  // 3% (JHT 2% + JP 1%)
		lines = append(lines, PayslipLine{Category: "deduction", Name: "BPJS Kesehatan (1%)", Amount: bpjsKes})
		lines = append(lines, PayslipLine{Category: "deduction", Name: "BPJS Ketenagakerjaan (3%)", Amount: bpjsTk})
		totalDeduction += (bpjsKes + bpjsTk)

		// Hitung PPh 21 TER 2024 resmi berdasarkan Penghasilan Bruto (totalEarning) dan Kategori PTKP
		ptkpStatus := emp.PTKPStatus
		if ptkpStatus == "" {
			ptkpStatus = "TK/0"
		}
		pph21, terCat, terRate := CalculatePPh21TER(totalEarning, ptkpStatus)
		if pph21 > 0 {
			lineName := fmt.Sprintf("PPh 21 TER 2024 (Kat. %s - %.2f%%)", terCat, terRate*100)
			lines = append(lines, PayslipLine{Category: "deduction", Name: lineName, Amount: pph21})
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
func MarkPayslipPaid(id uint) error {
	var payslip Payslip
	if err := config.DB.Preload("Employee").First(&payslip, id).Error; err != nil {
		return err
	}

	// 1. Update status ke paid
	if err := config.DB.Model(&payslip).Update("status", "paid").Error; err != nil {
		return err
	}

	// 2. Auto-Post ke Modul Finance (Journal Entry & Journal Items)
	// Cari akun kas/bank (1-1002 atau 1-1003) dan akun beban gaji (6-1000)
	var expenseAccountID, bankAccountID, taxAccountID uint
	config.DB.Table("finance.accounts").Select("id").Where("code = ?", "6-1000").Scan(&expenseAccountID)
	config.DB.Table("finance.accounts").Select("id").Where("code = ?", "1-1003").Scan(&bankAccountID) // Mandiri Payroll
	if bankAccountID == 0 {
		config.DB.Table("finance.accounts").Select("id").Where("code = ?", "1-1002").Scan(&bankAccountID) // BCA Operasional fallback
	}
	config.DB.Table("finance.accounts").Select("id").Where("code = ?", "2-1200").Scan(&taxAccountID) // Hutang PPh 21

	if expenseAccountID > 0 && bankAccountID > 0 {
		entryName := fmt.Sprintf("PAYROLL/%s/%s", payslip.Period, payslip.Employee.Name)
		
		type SimpleJournalEntry struct {
			ID        uint      `gorm:"primaryKey"`
			Name      string
			Date      time.Time
			State     string
			CreatedAt time.Time
		}
		type SimpleJournalItem struct {
			ID        uint    `gorm:"primaryKey"`
			EntryID   uint
			AccountID uint
			Name      string
			Debit     float64
			Credit    float64
		}

		entry := SimpleJournalEntry{
			Name:      entryName,
			Date:      time.Now(),
			State:     "posted",
			CreatedAt: time.Now(),
		}
		config.DB.Table("finance.journal_entrys").Create(&entry)

		// Debit Beban Gaji Total
		itemDebit := SimpleJournalItem{
			EntryID:   entry.ID,
			AccountID: expenseAccountID,
			Name:      fmt.Sprintf("Beban Gaji: %s (%s)", payslip.Employee.Name, payslip.Period),
			Debit:     payslip.TotalEarning,
			Credit:    0,
		}
		config.DB.Table("finance.journal_items").Create(&itemDebit)

		// Kredit Kas Bank (Take Home Pay)
		itemCreditBank := SimpleJournalItem{
			EntryID:   entry.ID,
			AccountID: bankAccountID,
			Name:      fmt.Sprintf("Pembayaran Gaji Net (THP) - %s", payslip.Employee.Name),
			Debit:     0,
			Credit:    payslip.NetSalary,
		}
		config.DB.Table("finance.journal_items").Create(&itemCreditBank)

		// Kredit Hutang PPh 21 / Potongan (Jika ada potongan)
		if payslip.TotalDeduction > 0 && taxAccountID > 0 {
			itemCreditTax := SimpleJournalItem{
				EntryID:   entry.ID,
				AccountID: taxAccountID,
				Name:      fmt.Sprintf("Titipan Potongan Pajak / PPh 21 - %s", payslip.Employee.Name),
				Debit:     0,
				Credit:    payslip.TotalDeduction,
			}
			config.DB.Table("finance.journal_items").Create(&itemCreditTax)
		}

		// Update saldo rekening bank operasional berkurang
		config.DB.Table("finance.accounts").Where("id = ?", bankAccountID).UpdateColumn("balance", gorm.Expr("balance - ?", payslip.NetSalary))
		// Update saldo beban bertambah
		config.DB.Table("finance.accounts").Where("id = ?", expenseAccountID).UpdateColumn("balance", gorm.Expr("balance + ?", payslip.TotalEarning))
	}

	return nil
}

func BulkPayPayslips(period string) (int, float64, error) {
	var payslips []Payslip
	query := config.DB.Preload("Employee").Where("status = 'draft'")
	if period != "" {
		query = query.Where("period = ?", period)
	}
	if err := query.Find(&payslips).Error; err != nil {
		return 0, 0, err
	}

	paidCount := 0
	totalAmount := 0.0
	for _, ps := range payslips {
		if err := MarkPayslipPaid(ps.ID); err == nil {
			paidCount++
			totalAmount += ps.NetSalary
		}
	}
	return paidCount, totalAmount, nil
}

// --- THR Repository ---
func GenerateTHR(year int, cutoffDate time.Time) error {
	var employees []Employee
	if err := config.DB.Find(&employees).Error; err != nil {
		return err
	}

	for _, emp := range employees {
		// Dapatkan gaji pokok dari kontrak aktif terbaru
		var contract Contract
		var wage float64 = 0
		var contractStart time.Time = emp.CreatedAt
		if err := config.DB.Where("employee_id = ?", emp.ID).Order("created_at desc").First(&contract).Error; err == nil {
			wage = contract.Wage
			if !contract.StartDate.IsZero() {
				contractStart = contract.StartDate
			}
		}

		if wage == 0 {
			continue
		}

		// Hitung masa kerja (JoinDate jika ada, atau ContractStart / CreatedAt)
		joinDate := contractStart
		if emp.JoinDate != nil && !emp.JoinDate.IsZero() {
			joinDate = *emp.JoinDate
		}

		// Hitung selisih bulan antara joinDate dan cutoffDate
		yearsDiff := cutoffDate.Year() - joinDate.Year()
		monthsDiff := int(cutoffDate.Month()) - int(joinDate.Month())
		totalMonths := yearsDiff*12 + monthsDiff
		// Sesuaikan bila hari belum mencapai
		if cutoffDate.Day() < joinDate.Day() {
			totalMonths--
		}

		if totalMonths < 1 {
			// Kurang dari 1 bulan tidak dapat THR sesuai Permenaker No. 6/2016
			continue
		}

		var thrAmount float64 = 0
		var calcType string = ""

		if totalMonths >= 12 {
			// Masa kerja 12 bulan atau lebih: 1 bulan upah penuh
			thrAmount = wage
			calcType = "Full 1 Bulan Upah (Masa Kerja >= 1 Tahun)"
		} else {
			// Masa kerja 1 bulan s/d < 12 bulan: Prorata N / 12 * Upah
			thrAmount = (float64(totalMonths) / 12.0) * wage
			calcType = fmt.Sprintf("Prorata %d/12 Bulan (Masa Kerja %d Bulan)", totalMonths, totalMonths)
		}

		// Cek jika sudah pernah di-generate untuk tahun ini
		var existing EmployeeTHR
		if err := config.DB.Where("employee_id = ? AND year = ?", emp.ID, year).First(&existing).Error; err == nil {
			// Update yang ada
			existing.BasicWage = wage
			existing.THRAmount = thrAmount
			existing.CutoffDate = cutoffDate
			existing.JoinDate = joinDate
			existing.TenureMonths = totalMonths
			existing.CalculationType = calcType
			config.DB.Save(&existing)
		} else {
			// Buat baru
			newTHR := EmployeeTHR{
				EmployeeID:      emp.ID,
				Year:            year,
				CutoffDate:      cutoffDate,
				JoinDate:        joinDate,
				TenureMonths:    totalMonths,
				BasicWage:       wage,
				THRAmount:       thrAmount,
				CalculationType: calcType,
				Status:          "draft",
			}
			config.DB.Create(&newTHR)
		}
	}
	return nil
}

func GetAllTHR(year int) ([]EmployeeTHR, error) {
	var list []EmployeeTHR
	query := config.DB.Preload("Employee")
	if year > 0 {
		query = query.Where("year = ?", year)
	}
	err := query.Order("id desc").Find(&list).Error
	return list, err
}

func UpdateTHRStatus(id uint, status string) error {
	return config.DB.Model(&EmployeeTHR{}).Where("id = ?", id).Update("status", status).Error
}
