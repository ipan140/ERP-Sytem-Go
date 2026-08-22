package payroll

import (
	"encoding/json"
	"log"

	"ERP-System/pkg/rabbitmq"
)

type PayrollMassPayload struct {
	BatchID string `json:"batch_id"`
	Format  string `json:"format"` // "pdf" (Payslip) or "excel" (Bank transfer)
	UserID  uint   `json:"user_id"`
}

func GenerateMassPayslipService(batchID string, format string, userID uint) error {
	if rabbitmq.Channel != nil {
		req := PayrollMassPayload{
			BatchID: batchID,
			Format:  format,
			UserID:  userID,
		}
		body, _ := json.Marshal(req)
		err := rabbitmq.PublishEvent(rabbitmq.Channel, "hr_payroll_generate", body)
		log.Printf("💸 Event RabbitMQ: Generate Mass Payroll %s (%s) dikirim ke antrean!", batchID, format)
		return err
	}
	return nil
}

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
