package payroll

func CreatePayslipService(data *Payslip) error {
	return CreatePayslip(data)
}

func GetAllPayslipService() ([]Payslip, error) {
	return GetAllPayslip()
}

func GetPayslipByIDService(id uint) (*Payslip, error) {
	return GetPayslipByID(id)
}

func UpdatePayslipService(data *Payslip) error {
	return UpdatePayslip(data)
}

func DeletePayslipService(id uint) error {
	return DeletePayslip(id)
}
