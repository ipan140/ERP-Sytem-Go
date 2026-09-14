package payroll

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"ERP-System/app/modules/finance/tax"
	"ERP-System/config"
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

func GetPaginatedPayslipsService(offset, limit int, search, employeeID, state string) ([]Payslip, int64, error) {
	return GetPaginatedPayslips(offset, limit, search, employeeID, state)
}

func GetPayslipByIDService(id uint) (*Payslip, error) {
	return GetPayslipByID(id)
}

func UpdatePayslipService(data *Payslip) error {
	err := UpdatePayslip(data)
	if err == nil && (data.State == "done" || data.State == "approved") {
		// Otomatis sinkronisasi PPh 21 TER ke modul Tax
		_ = SyncPPh21ToTaxModule(data.ID)
	}
	return err
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

// SyncPPh21ToTaxModule mengecek seluruh potongan PPh 21 dari payslip dan memposting ke tabel finance.tax_reports
func SyncPPh21ToTaxModule(payslipID uint) error {
	var ps Payslip
	if err := config.DB.Preload("Employee").First(&ps, payslipID).Error; err != nil {
		return err
	}

	// Cari baris gaji dengan kode/rule TAX / PPH21
	var lines []PayslipLine
	config.DB.Joins("JOIN hrd.salary_rules ON hrd.salary_rules.id = hrd.payslip_lines.salary_rule_id").
		Where("hrd.payslip_lines.payslip_id = ? AND (hrd.salary_rules.code ILIKE '%tax%' OR hrd.salary_rules.code ILIKE '%pph%')", payslipID).
		Find(&lines)

	var taxAmount float64
	for _, l := range lines {
		taxAmount += l.Amount
	}

	if taxAmount <= 0 {
		// Fallback simulasi TER (5% dari gaji bruto jika rule belum di-setup)
		var totalGross float64
		config.DB.Model(&PayslipLine{}).Where("payslip_id = ? AND amount > 0", payslipID).Select("COALESCE(SUM(amount), 0)").Scan(&totalGross)
		if totalGross > 5400000 {
			taxAmount = totalGross * 0.05
		}
	}

	periodStr := ps.DateFrom.Format("January 2006")
	if ps.DateFrom.IsZero() {
		periodStr = time.Now().Format("January 2006")
	}

	partnerName := "Karyawan Internal Kantor"
	if ps.Employee != nil && ps.Employee.Name != "" {
		partnerName = ps.Employee.Name
	}

	taxReport := tax.TaxReportSummary{
		TaxPeriod:   periodStr,
		TaxType:     "PPh 21 TER Bulanan (Gaji Karyawan)",
		TaxBase:     taxAmount / 0.05, // DPP Estimasi
		TaxRate:     5.0,
		TaxAmount:   taxAmount,
		PartnerName: fmt.Sprintf("%s (Slip: %s)", partnerName, ps.Name),
		NPWP:        "01.234.567.8-012.000",
		Status:      "Siap Lapor",
		CreatedAt:   time.Now(),
	}

	return tax.CreateTaxReportService(&taxReport)
}
