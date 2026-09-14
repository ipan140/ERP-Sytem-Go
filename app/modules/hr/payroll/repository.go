package payroll

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreatePayslip(data *Payslip) error {
	return config.DB.Create(data).Error
}

func GetAllPayslip() ([]Payslip, error) {
	var list []Payslip
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedPayslips(offset, limit int, search, employeeID, state string) ([]Payslip, int64, error) {
	var list []Payslip
	var total int64

	query := config.DB.Model(&Payslip{})

	if employeeID != "" && employeeID != "all" && employeeID != "0" {
		query = query.Where("hrd.payslips.employee_id = ?", employeeID)
	}

	if state != "" && state != "all" {
		query = query.Where("hrd.payslips.state = ?", state)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = hrd.payslips.employee_id").
			Where("hrd.payslips.name ILIKE ? OR hrd.employees.name ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Preload(clause.Associations).Order("hrd.payslips.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetPayslipByID(id uint) (*Payslip, error) {
	var data Payslip
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdatePayslip(data *Payslip) error {
	return config.DB.Save(data).Error
}

func DeletePayslip(id uint) error {
	return config.DB.Delete(&Payslip{}, id).Error
}

func CreatePayslipLine(data *PayslipLine) error { return config.DB.Create(data).Error }
func GetAllPayslipLine() ([]PayslipLine, error) {
	var list []PayslipLine
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetPayslipLineByID(id uint) (*PayslipLine, error) {
	var data PayslipLine
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdatePayslipLine(data *PayslipLine) error { return config.DB.Save(data).Error }
func DeletePayslipLine(id uint) error           { return config.DB.Delete(&PayslipLine{}, id).Error }

func CreateSalaryRule(data *SalaryRule) error { return config.DB.Create(data).Error }
func GetAllSalaryRule() ([]SalaryRule, error) {
	var list []SalaryRule
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
func GetSalaryRuleByID(id uint) (*SalaryRule, error) {
	var data SalaryRule
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateSalaryRule(data *SalaryRule) error { return config.DB.Save(data).Error }
func DeleteSalaryRule(id uint) error          { return config.DB.Delete(&SalaryRule{}, id).Error }
