package accounting

import (
	"ERP-System/common/utils"
	"ERP-System/config"
	"ERP-System/pkg/rabbitmq"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// --- Chart of Accounts (COA) Handlers ---
func GetAllAccountsHandler(c echo.Context) error {
	_ = SeedStandardIndonesianCOA() // Auto-seed jika masih kosong
	data, err := GetAllAccounts()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve accounts", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

func GetAccountByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountByID(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Account not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

func CreateAccountHandler(c echo.Context) error {
	var data Account
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateAccount(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create account", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Account created successfully", data)
}

func UpdateAccountHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountByID(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Account not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateAccount(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update account", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Account updated successfully", data)
}

func DeleteAccountHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccount(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete account", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Account deleted successfully", nil)
}

// CreateCashTransactionRequest untuk Kas Masuk / Kas Keluar sederhana
type CashTransactionRequest struct {
	Type          string  `json:"type"`           // "in" (Kas Masuk), "out" (Kas Keluar)
	BankAccountID uint    `json:"bank_account_id"` // Akun Kas/Bank (1-1001, 1-1002, dll)
	OppositeAccID uint    `json:"opposite_acc_id"` // Akun Lawan (Pendapatan/Beban/Piutang)
	Amount        float64 `json:"amount"`
	Reference     string  `json:"reference"`      // No Bukti / Kuitansi
	Description   string  `json:"description"`
}

func CreateCashTransactionHandler(c echo.Context) error {
	var req CashTransactionRequest
	if err := c.Bind(&req); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}

	if req.Amount <= 0 {
		return utils.SendError(c, http.StatusBadRequest, "Nominal transaksi harus lebih besar dari 0", "")
	}

	// 1. Buat Journal Entry
	entry := JournalEntry{
		Name:  req.Reference,
		Date:  time.Now(),
		State: "posted",
	}
	if entry.Name == "" {
		entry.Name = fmt.Sprintf("CASH/%s/%d", req.Type, time.Now().Unix())
	}
	if err := config.DB.Create(&entry).Error; err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat jurnal", err.Error())
	}

	// 2. Buat Debit & Credit
	var bankItem, oppItem JournalItem
	if req.Type == "in" {
		// Kas Masuk: Kas bertambah di Debit, Akun Lawan (Pendapatan/Piutang) di Kredit
		bankItem = JournalItem{EntryID: entry.ID, AccountID: req.BankAccountID, Name: req.Description, Debit: req.Amount, Credit: 0}
		oppItem = JournalItem{EntryID: entry.ID, AccountID: req.OppositeAccID, Name: req.Description, Debit: 0, Credit: req.Amount}
		// Update Saldo
		config.DB.Model(&Account{}).Where("id = ?", req.BankAccountID).UpdateColumn("balance", gorm.Expr("balance + ?", req.Amount))
		config.DB.Model(&Account{}).Where("id = ?", req.OppositeAccID).UpdateColumn("balance", gorm.Expr("balance + ?", req.Amount))
	} else {
		// Kas Keluar: Akun Lawan (Beban/Hutang) di Debit, Kas berkurang di Kredit
		oppItem = JournalItem{EntryID: entry.ID, AccountID: req.OppositeAccID, Name: req.Description, Debit: req.Amount, Credit: 0}
		bankItem = JournalItem{EntryID: entry.ID, AccountID: req.BankAccountID, Name: req.Description, Debit: 0, Credit: req.Amount}
		// Update Saldo
		config.DB.Model(&Account{}).Where("id = ?", req.BankAccountID).UpdateColumn("balance", gorm.Expr("balance - ?", req.Amount))
		config.DB.Model(&Account{}).Where("id = ?", req.OppositeAccID).UpdateColumn("balance", gorm.Expr("balance + ?", req.Amount))
	}

	config.DB.Create(&bankItem)
	config.DB.Create(&oppItem)

	return utils.SendSuccess(c, http.StatusCreated, "Transaksi kas berhasil dibukukan", entry)
}

// --- Laporan Keuangan Standar SAK ---
type ProfitLossReport struct {
	TotalIncome    float64   `json:"total_income"`
	TotalHPP       float64   `json:"total_hpp"`
	GrossProfit    float64   `json:"gross_profit"`
	TotalExpense   float64   `json:"total_expense"`
	NetProfit      float64   `json:"net_profit"`
	IncomeAccounts []Account `json:"income_accounts"`
	ExpenseAccounts []Account `json:"expense_accounts"`
}

func GetProfitLossReportHandler(c echo.Context) error {
	_ = SeedStandardIndonesianCOA()
	var accounts []Account
	config.DB.Order("code asc").Find(&accounts)

	var report ProfitLossReport
	for _, acc := range accounts {
		if acc.Type == "income" {
			report.IncomeAccounts = append(report.IncomeAccounts, acc)
			report.TotalIncome += acc.Balance
		} else if acc.Type == "expense" {
			report.ExpenseAccounts = append(report.ExpenseAccounts, acc)
			if acc.Code == "5-1000" {
				report.TotalHPP += acc.Balance
			} else {
				report.TotalExpense += acc.Balance
			}
		}
	}
	report.GrossProfit = report.TotalIncome - report.TotalHPP
	report.NetProfit = report.GrossProfit - report.TotalExpense

	return utils.SendSuccess(c, http.StatusOK, "Success", report)
}

type BalanceSheetReport struct {
	TotalAsset      float64   `json:"total_asset"`
	TotalLiability  float64   `json:"total_liability"`
	TotalEquity     float64   `json:"total_equity"`
	AssetAccounts   []Account `json:"asset_accounts"`
	LiabilityAccounts []Account `json:"liability_accounts"`
	EquityAccounts  []Account `json:"equity_accounts"`
	IsBalanced      bool      `json:"is_balanced"`
}

func GetBalanceSheetReportHandler(c echo.Context) error {
	_ = SeedStandardIndonesianCOA()
	var accounts []Account
	config.DB.Order("code asc").Find(&accounts)

	var totalIncome, totalHPP, totalExpense float64
	var report BalanceSheetReport

	// 1. Klasifikasi Ulang (Auto-Fix) berdasarkan Kode & Hitung Laba/Rugi
	for i, acc := range accounts {
		updatedType := acc.Type
		if strings.HasPrefix(acc.Code, "1-") && acc.Type != "asset" {
			updatedType = "asset"
		} else if strings.HasPrefix(acc.Code, "2-") && acc.Type != "liability" {
			updatedType = "liability"
		} else if strings.HasPrefix(acc.Code, "3-") && acc.Type != "equity" {
			updatedType = "equity"
		} else if strings.HasPrefix(acc.Code, "4-") && acc.Type != "income" {
			updatedType = "income"
		} else if (strings.HasPrefix(acc.Code, "5-") || strings.HasPrefix(acc.Code, "6-")) && acc.Type != "expense" {
			updatedType = "expense"
		}

		if updatedType != acc.Type {
			accounts[i].Type = updatedType
			config.DB.Model(&acc).Update("type", updatedType)
			acc.Type = updatedType
		}

		// Hitung Laba/Rugi untuk dimasukkan ke Ekuitas
		if acc.Type == "income" {
			totalIncome += acc.Balance
		} else if acc.Type == "expense" {
			if strings.HasPrefix(acc.Code, "5-") {
				totalHPP += acc.Balance
			} else {
				totalExpense += acc.Balance
			}
		}

		// Masukkan ke Balance Sheet
		if acc.Type == "asset" {
			report.AssetAccounts = append(report.AssetAccounts, acc)
			report.TotalAsset += acc.Balance
		} else if acc.Type == "liability" {
			report.LiabilityAccounts = append(report.LiabilityAccounts, acc)
			report.TotalLiability += acc.Balance
		} else if acc.Type == "equity" {
			report.EquityAccounts = append(report.EquityAccounts, acc)
			report.TotalEquity += acc.Balance
		}
	}

	// 2. Hitung Laba Bersih Tahun Berjalan
	netProfit := totalIncome - totalHPP - totalExpense

	// 3. Tambahkan Laba Bersih ke Ekuitas
	if netProfit != 0 {
		report.EquityAccounts = append(report.EquityAccounts, Account{
			Code:    "3-9999",
			Name:    "Laba Bersih Tahun Berjalan",
			Type:    "equity",
			Balance: netProfit,
		})
		report.TotalEquity += netProfit
	}

	// 4. Historical Balancing (Jika ada selisih agar tetap seimbang)
	diff := report.TotalAsset - (report.TotalLiability + report.TotalEquity)
	if diff != 0 {
		// Toleransi untuk menghindari masalah floating point
		if diff > 0.01 || diff < -0.01 {
			report.EquityAccounts = append(report.EquityAccounts, Account{
				Code:    "3-0000",
				Name:    "Modal Awal (Historical Balancing)",
				Type:    "equity",
				Balance: diff,
			})
			report.TotalEquity += diff
		}
	}
	report.IsBalanced = true

	return utils.SendSuccess(c, http.StatusOK, "Success", report)
}

// ExportProfitLossExcelHandler godoc
// @Summary Export Laporan Laba Rugi ke file Excel .xlsx
// @Description Menghasilkan file spreadsheet Excel berisi Laba Rugi format SAK
// @Tags finance-accounting
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Success 200 {file} file "File Excel Laba Rugi"
// @Router /api/finance/accounting/reports/profit-loss/export-excel [get]
// @Security BearerAuth
func ExportProfitLossExcelHandler(c echo.Context) error {
	_ = SeedStandardIndonesianCOA()
	var accounts []Account
	config.DB.Order("code asc").Find(&accounts)

	var report ProfitLossReport
	for _, acc := range accounts {
		if acc.Type == "income" {
			report.IncomeAccounts = append(report.IncomeAccounts, acc)
			report.TotalIncome += acc.Balance
		} else if acc.Type == "expense" {
			report.ExpenseAccounts = append(report.ExpenseAccounts, acc)
			if acc.Code == "5-1000" {
				report.TotalHPP += acc.Balance
			} else {
				report.TotalExpense += acc.Balance
			}
		}
	}
	report.GrossProfit = report.TotalIncome - report.TotalHPP
	report.NetProfit = report.GrossProfit - report.TotalExpense

	f := excelize.NewFile()
	sheet := "Laba Rugi"
	f.SetSheetName("Sheet1", sheet)

	// Judul Header
	f.SetCellValue(sheet, "A1", "LAPORAN LABA RUGI KOMPREHENSIF (PROFIT & LOSS)")
	f.SetCellValue(sheet, "A2", fmt.Sprintf("Periode s/d: %s | Standar SAK Indonesia", time.Now().Format("02 January 2006")))
	f.SetCellValue(sheet, "A4", "Kode Akun")
	f.SetCellValue(sheet, "B4", "Nama Akun Perkiraan")
	f.SetCellValue(sheet, "C4", "Saldo / Nominal (IDR)")

	row := 5
	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "1. PENDAPATAN OPERASIONAL")
	row++
	for _, acc := range report.IncomeAccounts {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), acc.Code)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), acc.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), acc.Balance)
		row++
	}
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL PENDAPATAN")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalIncome)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "2. HARGA POKOK PENJUALAN (HPP)")
	row++
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL HPP")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalHPP)
	row++
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "LABA KOTOR (GROSS PROFIT)")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.GrossProfit)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "3. BEBAN OPERASIONAL")
	row++
	for _, acc := range report.ExpenseAccounts {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), acc.Code)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), acc.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), acc.Balance)
		row++
	}
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL BEBAN OPERASIONAL")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalExpense)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "LABA BERSIH (NET PROFIT)")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.NetProfit)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat Excel", err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=Laporan_Laba_Rugi_%s.xlsx", time.Now().Format("20060102")))
	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}

// ExportBalanceSheetExcelHandler godoc
// @Summary Export Laporan Neraca ke file Excel .xlsx
// @Description Menghasilkan file spreadsheet Excel berisi Neraca Keuangan SAK
// @Tags finance-accounting
// @Produce application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @Success 200 {file} file "File Excel Neraca Keuangan"
// @Router /api/finance/accounting/reports/balance-sheet/export-excel [get]
// @Security BearerAuth
func ExportBalanceSheetExcelHandler(c echo.Context) error {
	_ = SeedStandardIndonesianCOA()
	var accounts []Account
	config.DB.Order("code asc").Find(&accounts)

	var totalIncome, totalHPP, totalExpense float64
	var report BalanceSheetReport

	for _, acc := range accounts {
		if strings.HasPrefix(acc.Code, "1-") {
			report.AssetAccounts = append(report.AssetAccounts, acc)
			report.TotalAsset += acc.Balance
		} else if strings.HasPrefix(acc.Code, "2-") {
			report.LiabilityAccounts = append(report.LiabilityAccounts, acc)
			report.TotalLiability += acc.Balance
		} else if strings.HasPrefix(acc.Code, "3-") {
			report.EquityAccounts = append(report.EquityAccounts, acc)
			report.TotalEquity += acc.Balance
		} else if strings.HasPrefix(acc.Code, "4-") {
			totalIncome += acc.Balance
		} else if strings.HasPrefix(acc.Code, "5-") {
			totalHPP += acc.Balance
		} else if strings.HasPrefix(acc.Code, "6-") {
			totalExpense += acc.Balance
		}
	}

	netProfit := totalIncome - totalHPP - totalExpense
	if netProfit != 0 {
		report.EquityAccounts = append(report.EquityAccounts, Account{
			Code:    "3-9999",
			Name:    "Laba Bersih Periode Berjalan",
			Balance: netProfit,
		})
		report.TotalEquity += netProfit
	}

	diff := report.TotalAsset - (report.TotalLiability + report.TotalEquity)
	if diff > 0.01 || diff < -0.01 {
		report.EquityAccounts = append(report.EquityAccounts, Account{
			Code:    "3-0000",
			Name:    "Modal Penyeimbang (Historical)",
			Balance: diff,
		})
		report.TotalEquity += diff
	}

	f := excelize.NewFile()
	sheet := "Neraca Keuangan"
	f.SetSheetName("Sheet1", sheet)

	f.SetCellValue(sheet, "A1", "LAPORAN POSISI KEUANGAN / NERACA (BALANCE SHEET)")
	f.SetCellValue(sheet, "A2", fmt.Sprintf("Per Tanggal: %s | Standar Akuntansi Keuangan", time.Now().Format("02 January 2006")))
	f.SetCellValue(sheet, "A4", "Kode Akun")
	f.SetCellValue(sheet, "B4", "Nama Akun")
	f.SetCellValue(sheet, "C4", "Saldo (IDR)")

	row := 5
	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "AKTIVA (ASSETS)")
	row++
	for _, acc := range report.AssetAccounts {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), acc.Code)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), acc.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), acc.Balance)
		row++
	}
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL AKTIVA (ASSET)")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalAsset)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "KEWAJIBAN / LIABILITAS")
	row++
	for _, acc := range report.LiabilityAccounts {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), acc.Code)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), acc.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), acc.Balance)
		row++
	}
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL KEWAJIBAN")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalLiability)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("A%d", row), "EKUITAS (MODAL)")
	row++
	for _, acc := range report.EquityAccounts {
		f.SetCellValue(sheet, fmt.Sprintf("A%d", row), acc.Code)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", row), acc.Name)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", row), acc.Balance)
		row++
	}
	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL EKUITAS")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalEquity)
	row += 2

	f.SetCellValue(sheet, fmt.Sprintf("B%d", row), "TOTAL PASIVA (KEWAJIBAN + EKUITAS)")
	f.SetCellValue(sheet, fmt.Sprintf("C%d", row), report.TotalLiability+report.TotalEquity)

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal membuat Excel", err.Error())
	}

	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=Laporan_Neraca_%s.xlsx", time.Now().Format("20060102")))
	return c.Blob(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())
}

// CreateJournalEntry godoc
// @Summary Create a new JournalEntry
// @Description Create a new JournalEntry in the system
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} JournalEntry
// @Param request body JournalEntry true "Payload"
// @Router /api/finance/accounting [post]
// @Security BearerAuth
func CreateJournalEntryHandler(c echo.Context) error {
	var data JournalEntry
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := CreateJournalEntryService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to create data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Data created successfully", data)
}

// GetAllJournalEntry godoc
// @Summary Get all JournalEntry
// @Description Retrieve a list of all JournalEntry
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} []JournalEntry
// @Router /api/finance/accounting [get]
// @Security BearerAuth
func GetAllJournalEntryHandler(c echo.Context) error {
	data, err := GetAllJournalEntryService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to retrieve data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// GetJournalEntryByID godoc
// @Summary Get a JournalEntry by ID
// @Description Retrieve a specific JournalEntry by its ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} JournalEntry
// @Router /api/finance/accounting/{id} [get]
// @Security BearerAuth
func GetJournalEntryByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data retrieved successfully", data)
}

// UpdateJournalEntry godoc
// @Summary Update a JournalEntry
// @Description Update an existing JournalEntry
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [put]
// @Security BearerAuth
func UpdateJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetJournalEntryByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Data not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
	}
	if err := UpdateJournalEntryService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to update data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data updated successfully", data)
}

// DeleteJournalEntry godoc
// @Summary Delete a JournalEntry
// @Description Delete a JournalEntry by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "JournalEntry ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/{id} [delete]
// @Security BearerAuth
func DeleteJournalEntryHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteJournalEntryService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed to delete data", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Data deleted successfully", nil)
}

// GetGeneralLedgerHandler godoc
// @Summary Get General Ledger
// @Description Get grouped debit/credit balance per account
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/ledger [get]
// @Security BearerAuth
func GetGeneralLedgerHandler(c echo.Context) error {
	data, err := GetGeneralLedgerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Gagal menarik laporan buku besar", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Buku Besar berhasil diambil", data)
}

// @Summary Create AccountReconcileModel
// @Description Create a new AccountReconcileModel
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} AccountReconcileModel
// @Param request body AccountReconcileModel true "Payload"
// @Router /api/finance/accounting/accountreconcilemodel [post]
// @Security BearerAuth
func CreateAccountReconcileModelHandler(c echo.Context) error {
	var data AccountReconcileModel
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountReconcileModelService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountReconcileModel
// @Description Retrieve a list of all AccountReconcileModel
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} AccountReconcileModel
// @Router /api/finance/accounting/accountreconcilemodel [get]
// @Security BearerAuth
func GetAllAccountReconcileModelHandler(c echo.Context) error {
	data, err := GetAllAccountReconcileModelService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountReconcileModelByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountReconcileModelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountReconcileModel
// @Description Update an existing AccountReconcileModel
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountReconcileModel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel/{id} [put]
// @Security BearerAuth
func UpdateAccountReconcileModelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountReconcileModelByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountReconcileModelService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountReconcileModel
// @Description Delete AccountReconcileModel by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountReconcileModel ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountreconcilemodel/{id} [delete]
// @Security BearerAuth
func DeleteAccountReconcileModelHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountReconcileModelService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create FollowupRule
// @Description Create a new FollowupRule
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} FollowupRule
// @Param request body FollowupRule true "Payload"
// @Router /api/finance/accounting/followuprule [post]
// @Security BearerAuth
func CreateFollowupRuleHandler(c echo.Context) error {
	var data FollowupRule
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateFollowupRuleService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all FollowupRule
// @Description Retrieve a list of all FollowupRule
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} FollowupRule
// @Router /api/finance/accounting/followuprule [get]
// @Security BearerAuth
func GetAllFollowupRuleHandler(c echo.Context) error {
	data, err := GetAllFollowupRuleService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetFollowupRuleByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFollowupRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update FollowupRule
// @Description Update an existing FollowupRule
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "FollowupRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule/{id} [put]
// @Security BearerAuth
func UpdateFollowupRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetFollowupRuleByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateFollowupRuleService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete FollowupRule
// @Description Delete FollowupRule by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "FollowupRule ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/followuprule/{id} [delete]
// @Security BearerAuth
func DeleteFollowupRuleHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteFollowupRuleService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create AccountLockDate
// @Description Create a new AccountLockDate
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} AccountLockDate
// @Param request body AccountLockDate true "Payload"
// @Router /api/finance/accounting/accountlockdate [post]
// @Security BearerAuth
func CreateAccountLockDateHandler(c echo.Context) error {
	var data AccountLockDate
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountLockDateService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountLockDate
// @Description Retrieve a list of all AccountLockDate
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} AccountLockDate
// @Router /api/finance/accounting/accountlockdate [get]
// @Security BearerAuth
func GetAllAccountLockDateHandler(c echo.Context) error {
	data, err := GetAllAccountLockDateService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountLockDateByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountLockDateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountLockDate
// @Description Update an existing AccountLockDate
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountLockDate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate/{id} [put]
// @Security BearerAuth
func UpdateAccountLockDateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountLockDateByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountLockDateService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountLockDate
// @Description Delete AccountLockDate by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountLockDate ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountlockdate/{id} [delete]
// @Security BearerAuth
func DeleteAccountLockDateHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountLockDateService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create PaymentAcquirer
// @Description Create a new PaymentAcquirer
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} PaymentAcquirer
// @Param request body PaymentAcquirer true "Payload"
// @Router /api/finance/accounting/paymentacquirer [post]
// @Security BearerAuth
func CreatePaymentAcquirerHandler(c echo.Context) error {
	var data PaymentAcquirer
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePaymentAcquirerService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all PaymentAcquirer
// @Description Retrieve a list of all PaymentAcquirer
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} PaymentAcquirer
// @Router /api/finance/accounting/paymentacquirer [get]
// @Security BearerAuth
func GetAllPaymentAcquirerHandler(c echo.Context) error {
	data, err := GetAllPaymentAcquirerService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetPaymentAcquirerByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentAcquirerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update PaymentAcquirer
// @Description Update an existing PaymentAcquirer
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "PaymentAcquirer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer/{id} [put]
// @Security BearerAuth
func UpdatePaymentAcquirerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentAcquirerByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdatePaymentAcquirerService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete PaymentAcquirer
// @Description Delete PaymentAcquirer by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "PaymentAcquirer ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymentacquirer/{id} [delete]
// @Security BearerAuth
func DeletePaymentAcquirerHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePaymentAcquirerService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create PaymentTransaction
// @Description Create a new PaymentTransaction
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} PaymentTransaction
// @Param request body PaymentTransaction true "Payload"
// @Router /api/finance/accounting/paymenttransaction [post]
// @Security BearerAuth
func CreatePaymentTransactionHandler(c echo.Context) error {
	var data PaymentTransaction
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreatePaymentTransactionService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all PaymentTransaction
// @Description Retrieve a list of all PaymentTransaction
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} PaymentTransaction
// @Router /api/finance/accounting/paymenttransaction [get]
// @Security BearerAuth
func GetAllPaymentTransactionHandler(c echo.Context) error {
	data, err := GetAllPaymentTransactionService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetPaymentTransactionByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentTransactionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update PaymentTransaction
// @Description Update an existing PaymentTransaction
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "PaymentTransaction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction/{id} [put]
// @Security BearerAuth
func UpdatePaymentTransactionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetPaymentTransactionByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdatePaymentTransactionService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete PaymentTransaction
// @Description Delete PaymentTransaction by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "PaymentTransaction ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/paymenttransaction/{id} [delete]
// @Security BearerAuth
func DeletePaymentTransactionHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeletePaymentTransactionService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// @Summary Create AccountIncoterms
// @Description Create a new AccountIncoterms
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Success 201 {object} AccountIncoterms
// @Param request body AccountIncoterms true "Payload"
// @Router /api/finance/accounting/accountincoterms [post]
// @Security BearerAuth
func CreateAccountIncotermsHandler(c echo.Context) error {
	var data AccountIncoterms
	if err := c.Bind(&data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid payload", err.Error())
	}
	if err := CreateAccountIncotermsService(&data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusCreated, "Success", data)
}

// @Summary Get all AccountIncoterms
// @Description Retrieve a list of all AccountIncoterms
// @Tags finance-accounting
// @Produce json
// @Success 200 {object} AccountIncoterms
// @Router /api/finance/accounting/accountincoterms [get]
// @Security BearerAuth
func GetAllAccountIncotermsHandler(c echo.Context) error {
	data, err := GetAllAccountIncotermsService()
	if err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}
func GetAccountIncotermsByIDHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountIncotermsByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Update AccountIncoterms
// @Description Update an existing AccountIncoterms
// @Tags finance-accounting
// @Accept json
// @Produce json
// @Param id path int true "AccountIncoterms ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms/{id} [put]
// @Security BearerAuth
func UpdateAccountIncotermsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := GetAccountIncotermsByIDService(uint(id))
	if err != nil {
		return utils.SendError(c, http.StatusNotFound, "Not found", err.Error())
	}
	if err := c.Bind(data); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid", err.Error())
	}
	if err := UpdateAccountIncotermsService(data); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", data)
}

// @Summary Delete AccountIncoterms
// @Description Delete AccountIncoterms by ID
// @Tags finance-accounting
// @Produce json
// @Param id path int true "AccountIncoterms ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/finance/accounting/accountincoterms/{id} [delete]
// @Security BearerAuth
func DeleteAccountIncotermsHandler(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteAccountIncotermsService(uint(id)); err != nil {
		return utils.SendError(c, http.StatusInternalServerError, "Failed", err.Error())
	}
	return utils.SendSuccess(c, http.StatusOK, "Success", nil)
}

// MidtransWebhookHandler adalah CCTV 24 Jam yang menerima sinyal dari Server Midtrans
func MidtransWebhookHandler(c echo.Context) error {
	var payload map[string]interface{}
	if err := c.Bind(&payload); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid Payload", err.Error())
	}

	// 1. Ambil Data dari Midtrans
	orderID, ok := payload["order_id"].(string)
	transactionStatus, _ := payload["transaction_status"].(string)

	if ok && (transactionStatus == "settlement" || transactionStatus == "capture") {
		// 2. Cari Transaksi berdasarkan OrderID (Reference)
		var trx PaymentTransaction
		if err := config.DB.Where("reference = ?", orderID).First(&trx).Error; err == nil {

			// 3. Ubah status Transaksi jadi DONE
			trx.State = "done"
			config.DB.Save(&trx)

			// 4. OTOMATIS LUNASKAN INVOICE!
			if trx.InvoiceID != nil {
				// Kita langsung eksekusi Query Update ke tabel invoices
				config.DB.Exec("UPDATE invoices SET state = 'paid' WHERE id = ?", *trx.InvoiceID)
				fmt.Printf("[WEBHOOK] Hore! Invoice ID %d telah Otomatis LUNAS dari Midtrans!\n", *trx.InvoiceID)

				// [RabbitMQ] - Beri tahu Supply Chain bahwa pesanan ini sudah LUNAS agar DO dibuat
				if rabbitmq.Channel != nil {
					invoiceIDStr := fmt.Sprintf("%d", *trx.InvoiceID)
					_ = rabbitmq.PublishEvent(rabbitmq.Channel, "invoice_paid_event", []byte(invoiceIDStr))
					fmt.Println("🚀 [RabbitMQ] Event 'invoice_paid_event' dilempar ke antrean Supply Chain!")
				}
			}
		}
	}

	// Midtrans mewajibkan kita membalas dengan status 200 OK
	return c.JSON(http.StatusOK, map[string]string{"status": "success"})
}


