package payroll

import (
	"ERP-System/config"
)

func CreatePayslip(data *Payslip) error {
	return config.DB.Create(data).Error
}

func GetAllPayslip() ([]Payslip, error) {
	var list []Payslip
	err := config.DB.Find(&list).Error
	return list, err
}

func GetPayslipByID(id uint) (*Payslip, error) {
	var data Payslip
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdatePayslip(data *Payslip) error {
	return config.DB.Save(data).Error
}

func DeletePayslip(id uint) error {
	return config.DB.Delete(&Payslip{}, id).Error
}

func CreatePayslipLine(data *PayslipLine) error { return config.DB.Create(data).Error }
func GetAllPayslipLine() ([]PayslipLine, error) { var list []PayslipLine; err := config.DB.Find(&list).Error; return list, err }
func GetPayslipLineByID(id uint) (*PayslipLine, error) { var data PayslipLine; err := config.DB.First(&data, id).Error; return &data, err }
func UpdatePayslipLine(data *PayslipLine) error { return config.DB.Save(data).Error }
func DeletePayslipLine(id uint) error { return config.DB.Delete(&PayslipLine{}, id).Error }

func CreateSalaryRule(data *SalaryRule) error { return config.DB.Create(data).Error }
func GetAllSalaryRule() ([]SalaryRule, error) { var list []SalaryRule; err := config.DB.Find(&list).Error; return list, err }
func GetSalaryRuleByID(id uint) (*SalaryRule, error) { var data SalaryRule; err := config.DB.First(&data, id).Error; return &data, err }
func UpdateSalaryRule(data *SalaryRule) error { return config.DB.Save(data).Error }
func DeleteSalaryRule(id uint) error { return config.DB.Delete(&SalaryRule{}, id).Error }
