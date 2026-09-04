package expenses

import (
	"ERP-System/app/modules/finance/budget"
	"ERP-System/config"
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

type EmployeeDeptInfo struct {
	DepartmentName string
}

func getEmployeeDepartmentName(employeeID uint) string {
	var res struct {
		Name string
	}
	err := config.DB.Table("hr.departments").
		Joins("JOIN hr.employees ON hr.employees.department_id = hr.departments.id").
		Where("hr.employees.id = ?", employeeID).
		Select("hr.departments.name as name").
		Scan(&res).Error
	if err != nil || res.Name == "" {
		return ""
	}
	return res.Name
}

func CreateExpenseService(data *Expense) error {
	// Integrasi Kontrol Anggaran (Budgeting Guard)
	if data.EmployeeID > 0 && data.TotalAmount > 0 {
		deptName := getEmployeeDepartmentName(data.EmployeeID)
		if deptName != "" {
			var bg budget.DepartmentBudget
			err := config.DB.Where("department_name ILIKE ?", "%"+deptName+"%").First(&bg).Error
			if err == nil {
				// Cek sisa anggaran
				remaining := bg.AllocatedLimit - bg.RealizedSpent
				if data.TotalAmount > remaining {
					return fmt.Errorf("pengajuan ditolak: Melebihi plafon anggaran divisi %s! Sisa anggaran: Rp %.0f, Pengajuan: Rp %.0f", deptName, remaining, data.TotalAmount)
				}
				// Tambahkan ke realisasi anggaran
				bg.RealizedSpent += data.TotalAmount
				_ = budget.UpdateBudgetService(&bg)
			}
		}
	}

	return CreateExpense(data)
}

func GetAllExpenseService() ([]Expense, error) {
	return GetAllExpense()
}

func GetExpenseByIDService(id uint) (*Expense, error) {
	return GetExpenseByID(id)
}

func UpdateExpenseService(data *Expense) error {
	return UpdateExpense(data)
}

func DeleteExpenseService(id uint) error {
	exp, err := GetExpenseByID(id)
	if err == nil && exp != nil && exp.EmployeeID > 0 && exp.TotalAmount > 0 {
		// Mengembalikan kuota realisasi anggaran jika klaim dihapus
		deptName := getEmployeeDepartmentName(exp.EmployeeID)
		if deptName != "" {
			var bg budget.DepartmentBudget
			err := config.DB.Where("department_name ILIKE ?", "%"+deptName+"%").First(&bg).Error
			if err == nil {
				bg.RealizedSpent -= exp.TotalAmount
				if bg.RealizedSpent < 0 {
					bg.RealizedSpent = 0
				}
				_ = budget.UpdateBudgetService(&bg)
			}
		}
	}
	return DeleteExpense(id)
}

func CreateExpenseSheetService(data *ExpenseSheet) error        { return CreateExpenseSheet(data) }
func GetAllExpenseSheetService() ([]ExpenseSheet, error)        { return GetAllExpenseSheet() }
func GetExpenseSheetByIDService(id uint) (*ExpenseSheet, error) { return GetExpenseSheetByID(id) }
func UpdateExpenseSheetService(data *ExpenseSheet) error        { return UpdateExpenseSheet(data) }
func DeleteExpenseSheetService(id uint) error                   { return DeleteExpenseSheet(id) }

func ExportExpensesExcelService(period string, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"action": "export_expenses", "period": period, "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("💸 Event RabbitMQ: Export Excel Klaim Biaya %s dikirim ke antrean!", period)
		return nil
	}
	return errors.New("rabbitmq channel is unavailable")
}

// --- PETTY CASH (KAS KECIL - SISTEM IMPREST) SERVICES ---

func GetOrCreatePettyCashFundService() (*PettyCashFund, []PettyCashTransaction, error) {
	var fund PettyCashFund
	err := config.DB.First(&fund).Error
	if err != nil {
		// Inisialisasi kas kecil default plafon Rp 5.000.000
		fund = PettyCashFund{
			Name:           "Kas Kecil Operasional Kantor",
			Custodian:      "Staff Keuangan & Kasir",
			PlafondLimit:   5000000,
			CurrentBalance: 5000000,
		}
		_ = config.DB.Create(&fund).Error
	}

	var txs []PettyCashTransaction
	_ = config.DB.Where("fund_id = ?", fund.ID).Order("id desc").Limit(50).Find(&txs).Error
	return &fund, txs, nil
}

func RecordPettyCashExpenseService(tx *PettyCashTransaction) error {
	var fund PettyCashFund
	if err := config.DB.First(&fund, tx.FundID).Error; err != nil {
		return fmt.Errorf("kas kecil tidak ditemukan: %v", err)
	}

	if tx.Amount > fund.CurrentBalance {
		return fmt.Errorf("saldo kas kecil tidak mencukupi! Sisa saldo fisik: Rp %.0f, pengeluaran: Rp %.0f. Harap lakukan penggantian (replenishment)", fund.CurrentBalance, tx.Amount)
	}

	tx.TxType = "expense"
	tx.CreatedAt = time.Now()
	if err := config.DB.Create(tx).Error; err != nil {
		return err
	}

	// Kurangi saldo berjalan
	fund.CurrentBalance -= tx.Amount
	fund.UpdatedAt = time.Now()
	_ = config.DB.Save(&fund).Error

	// Update juga saldo akun COA Kas Kecil (1-1001) jika ada
	config.DB.Table("accounts").Where("code = ?", "1-1001").
		UpdateColumn("balance", config.DB.Raw("balance - ?", tx.Amount))

	return nil
}

func ReplenishPettyCashService(fundID uint, recordedBy string) (float64, error) {
	var fund PettyCashFund
	if err := config.DB.First(&fund, fundID).Error; err != nil {
		return 0, fmt.Errorf("kas kecil tidak ditemukan: %v", err)
	}

	neededAmount := fund.PlafondLimit - fund.CurrentBalance
	if neededAmount <= 0 {
		return 0, fmt.Errorf("saldo kas kecil masih penuh pada batas plafon Rp %.0f", fund.PlafondLimit)
	}

	tx := PettyCashTransaction{
		FundID:      fund.ID,
		TxType:      "replenish",
		Amount:      neededAmount,
		Description: fmt.Sprintf("Penggantian Kas Kecil (Imprest Replenishment) ke batas plafon Rp %.0f", fund.PlafondLimit),
		ReceiptRef:  fmt.Sprintf("REPL/%s/%d", time.Now().Format("20060102"), time.Now().Unix()%1000),
		RecordedBy:  recordedBy,
		CreatedAt:   time.Now(),
	}
	if err := config.DB.Create(&tx).Error; err != nil {
		return 0, err
	}

	// Kembalikan saldo ke batas plafon
	fund.CurrentBalance = fund.PlafondLimit
	fund.UpdatedAt = time.Now()
	_ = config.DB.Save(&fund).Error

	// Update saldo akun Kas Kecil (1-1001) di GL
	config.DB.Table("accounts").Where("code = ?", "1-1001").
		UpdateColumn("balance", fund.PlafondLimit)

	return neededAmount, nil
}
