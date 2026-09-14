package consolidation

import (
	"ERP-System/config"
)

func GenerateConsolidationService(name, period string) error {
	report := &ConsolidationReport{
		Name:   name,
		Period: period,
	}
	if err := config.DB.Create(report).Error; err != nil {
		return err
	}

	// Simulasi agregasi data dari seluruh cabang (Company)
	// Kita akan group saldo dari seluruh account_id yang ada di journal_items
	type Result struct {
		AccountID uint
		Balance   float64
	}
	var results []Result

	config.DB.Table("journal_items").
		Select("account_id, (sum(debit) - sum(credit)) as balance").
		Group("account_id").
		Scan(&results)

	for _, res := range results {
		cAcc := &ConsolidatedAccount{
			ReportID:  report.ID,
			AccountID: res.AccountID,
			Balance:   res.Balance,
		}
		config.DB.Create(cAcc)
	}
	return nil
}

func GetAllConsolidationReportsService() ([]ConsolidationReport, error) {
	return GetAllConsolidationReports()
}

func GetPaginatedConsolidationReportsService(offset, limit int, search string) ([]ConsolidationReport, int64, error) {
	return GetPaginatedConsolidationReports(offset, limit, search)
}

func GetConsolidationReportByIDService(id uint) (*ConsolidationReport, error) {
	return GetConsolidationReportByID(id)
}

func CreateConsolidationReportService(data *ConsolidationReport) error {
	if err := CreateConsolidationReport(data); err != nil {
		return err
	}
	// Buat simulasi consolidated accounts
	type Result struct {
		AccountID uint
		Balance   float64
	}
	var results []Result

	config.DB.Table("journal_items").
		Select("account_id, (sum(debit) - sum(credit)) as balance").
		Group("account_id").
		Scan(&results)

	for _, res := range results {
		cAcc := &ConsolidatedAccount{
			ReportID:  data.ID,
			AccountID: res.AccountID,
			Balance:   res.Balance,
		}
		config.DB.Create(cAcc)
	}
	return nil
}

func UpdateConsolidationReportService(data *ConsolidationReport) error {
	return UpdateConsolidationReport(data)
}

func DeleteConsolidationReportService(id uint) error {
	return DeleteConsolidationReport(id)
}
