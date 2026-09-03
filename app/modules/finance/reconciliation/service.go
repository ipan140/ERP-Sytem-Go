package reconciliation

import (
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
