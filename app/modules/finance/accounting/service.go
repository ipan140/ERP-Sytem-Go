package accounting

import "ERP-System/config"

func CreateJournalEntryService(data *JournalEntry) error {
	return CreateJournalEntry(data)
}

func GetAllJournalEntryService() ([]JournalEntry, error) {
	return GetAllJournalEntry()
}

func GetJournalEntryByIDService(id uint) (*JournalEntry, error) {
	return GetJournalEntryByID(id)
}

func UpdateJournalEntryService(data *JournalEntry) error {
	return UpdateJournalEntry(data)
}

func DeleteJournalEntryService(id uint) error {
	return DeleteJournalEntry(id)
}

func CreateJournalItemService(data *JournalItem) error {
	return CreateJournalItem(data)
}

type LedgerResult struct {
	AccountID   uint    `json:"account_id"`
	AccountName string  `json:"account_name"`
	TotalDebit  float64 `json:"total_debit"`
	TotalCredit float64 `json:"total_credit"`
	Balance     float64 `json:"balance"`
}

func GetGeneralLedgerService() ([]LedgerResult, error) {
	var results []LedgerResult

	// Query GORM untuk Group By Account
	rows, err := config.DB.Table("journal_items").
		Select("journal_items.account_id, accounts.name as account_name, sum(journal_items.debit) as total_debit, sum(journal_items.credit) as total_credit, (sum(journal_items.debit) - sum(journal_items.credit)) as balance").
		Joins("left join accounts on accounts.id = journal_items.account_id").
		Group("journal_items.account_id, accounts.name").
		Rows()

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var res LedgerResult
		config.DB.ScanRows(rows, &res)
		results = append(results, res)
	}

	return results, nil
}
