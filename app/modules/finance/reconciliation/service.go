package reconciliation

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
	"time"
)

func GetAllBankStatementsService() ([]BankStatementItem, error) {
	return GetAllBankStatementsRepo()
}

func CreateBankStatementService(item *BankStatementItem) error {
	if item.Date.IsZero() {
		item.Date = time.Now()
	}
	return CreateBankStatementRepo(item)
}

func UpdateBankStatementService(item *BankStatementItem) error {
	return UpdateBankStatementRepo(item)
}

func DeleteBankStatementService(id uint) error {
	return DeleteBankStatementRepo(id)
}

func AutoReconcileService() (int, error) {
	unreconciled, err := GetUnreconciledStatementsRepo()
	if err != nil {
		return 0, err
	}

	matchedCount := 0
	for _, item := range unreconciled {
		if item.Credit > 0 && item.MatchedInvoice == "" {
			item.IsReconciled = true
			item.MatchedInvoice = "INV/AUTO/" + item.RefNumber
			_ = UpdateBankStatementRepo(&item)
			matchedCount++
		}
	}
	return matchedCount, nil
}

func ManualReconcileService(id uint, matchedInvoice string) (*BankStatementItem, error) {
	item, err := GetBankStatementByIDRepo(id)
	if err != nil {
		return nil, err
	}
	item.IsReconciled = true
	item.MatchedInvoice = matchedInvoice
	err = UpdateBankStatementRepo(item)
	return item, err
}

// Import Bank CSV (KlikBCA / Mandiri MCM Standard)
// Header expected: Tanggal, Keterangan, No. Ref, Debit, Kredit, Bank
func ImportBankStatementCsvService(reader io.Reader) (int, error) {
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = -1 // Allow variable fields
	csvReader.TrimLeadingSpace = true

	rows, err := csvReader.ReadAll()
	if err != nil {
		return 0, err
	}

	importedCount := 0
	for i, row := range rows {
		// Skip header row
		if i == 0 && (strings.Contains(strings.ToLower(row[0]), "tanggal") || strings.Contains(strings.ToLower(row[0]), "date")) {
			continue
		}
		if len(row) < 3 {
			continue
		}

		// Parse date (support YYYY-MM-DD or DD/MM/YYYY)
		dateStr := strings.TrimSpace(row[0])
		txDate, parseErr := time.Parse("2006-01-02", dateStr)
		if parseErr != nil {
			txDate, parseErr = time.Parse("02/01/2006", dateStr)
			if parseErr != nil {
				txDate = time.Now()
			}
		}

		desc := strings.TrimSpace(row[1])
		ref := ""
		if len(row) > 2 {
			ref = strings.TrimSpace(row[2])
		}

		var debit, credit float64
		if len(row) > 3 {
			dStr := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(row[3]), ",", ""), " ", "")
			debit, _ = strconv.ParseFloat(dStr, 64)
		}
		if len(row) > 4 {
			cStr := strings.ReplaceAll(strings.ReplaceAll(strings.TrimSpace(row[4]), ",", ""), " ", "")
			credit, _ = strconv.ParseFloat(cStr, 64)
		}

		bankName := "BCA Giro Operasional"
		if len(row) > 5 && strings.TrimSpace(row[5]) != "" {
			bankName = strings.TrimSpace(row[5])
		}

		item := BankStatementItem{
			Date:         txDate,
			Description:  desc,
			RefNumber:    ref,
			Debit:        debit,
			Credit:       credit,
			BankName:     bankName,
			IsReconciled: false,
		}

		if err := CreateBankStatementRepo(&item); err == nil {
			importedCount++
		}
	}

	return importedCount, nil
}
