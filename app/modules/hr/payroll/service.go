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

func CreatePayslipLineService(data *PayslipLine) error        { return CreatePayslipLine(data) }
func GetAllPayslipLineService() ([]PayslipLine, error)        { return GetAllPayslipLine() }
func GetPayslipLineByIDService(id uint) (*PayslipLine, error) { return GetPayslipLineByID(id) }
func UpdatePayslipLineService(data *PayslipLine) error        { return UpdatePayslipLine(data) }
func DeletePayslipLineService(id uint) error                  { return DeletePayslipLine(id) }

func CreateSalaryRuleService(data *SalaryRule) error        { return CreateSalaryRule(data) }
func GetAllSalaryRuleService() ([]SalaryRule, error)        { return GetAllSalaryRule() }
func GetSalaryRuleByIDService(id uint) (*SalaryRule, error) { return GetSalaryRuleByID(id) }
func UpdateSalaryRuleService(data *SalaryRule) error        { return UpdateSalaryRule(data) }
func DeleteSalaryRuleService(id uint) error                 { return DeleteSalaryRule(id) }
