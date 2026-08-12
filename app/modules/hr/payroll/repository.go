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
