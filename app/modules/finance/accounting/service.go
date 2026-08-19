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

func CreateAccountReconcileModelService(data *AccountReconcileModel) error { return CreateAccountReconcileModel(data) }
func GetAllAccountReconcileModelService() ([]AccountReconcileModel, error) { return GetAllAccountReconcileModel() }
func GetAccountReconcileModelByIDService(id uint) (*AccountReconcileModel, error) { return GetAccountReconcileModelByID(id) }
func UpdateAccountReconcileModelService(data *AccountReconcileModel) error { return UpdateAccountReconcileModel(data) }
func DeleteAccountReconcileModelService(id uint) error { return DeleteAccountReconcileModel(id) }

func CreateFollowupRuleService(data *FollowupRule) error { return CreateFollowupRule(data) }
func GetAllFollowupRuleService() ([]FollowupRule, error) { return GetAllFollowupRule() }
func GetFollowupRuleByIDService(id uint) (*FollowupRule, error) { return GetFollowupRuleByID(id) }
func UpdateFollowupRuleService(data *FollowupRule) error { return UpdateFollowupRule(data) }
func DeleteFollowupRuleService(id uint) error { return DeleteFollowupRule(id) }

func CreateAccountLockDateService(data *AccountLockDate) error { return CreateAccountLockDate(data) }
func GetAllAccountLockDateService() ([]AccountLockDate, error) { return GetAllAccountLockDate() }
func GetAccountLockDateByIDService(id uint) (*AccountLockDate, error) { return GetAccountLockDateByID(id) }
func UpdateAccountLockDateService(data *AccountLockDate) error { return UpdateAccountLockDate(data) }
func DeleteAccountLockDateService(id uint) error { return DeleteAccountLockDate(id) }

func CreatePaymentAcquirerService(data *PaymentAcquirer) error { return CreatePaymentAcquirer(data) }
func GetAllPaymentAcquirerService() ([]PaymentAcquirer, error) { return GetAllPaymentAcquirer() }
func GetPaymentAcquirerByIDService(id uint) (*PaymentAcquirer, error) { return GetPaymentAcquirerByID(id) }
func UpdatePaymentAcquirerService(data *PaymentAcquirer) error { return UpdatePaymentAcquirer(data) }
func DeletePaymentAcquirerService(id uint) error { return DeletePaymentAcquirer(id) }

func CreatePaymentTransactionService(data *PaymentTransaction) error { return CreatePaymentTransaction(data) }
func GetAllPaymentTransactionService() ([]PaymentTransaction, error) { return GetAllPaymentTransaction() }
func GetPaymentTransactionByIDService(id uint) (*PaymentTransaction, error) { return GetPaymentTransactionByID(id) }
func UpdatePaymentTransactionService(data *PaymentTransaction) error { return UpdatePaymentTransaction(data) }
func DeletePaymentTransactionService(id uint) error { return DeletePaymentTransaction(id) }

func CreateAccountIncotermsService(data *AccountIncoterms) error { return CreateAccountIncoterms(data) }
func GetAllAccountIncotermsService() ([]AccountIncoterms, error) { return GetAllAccountIncoterms() }
func GetAccountIncotermsByIDService(id uint) (*AccountIncoterms, error) { return GetAccountIncotermsByID(id) }
func UpdateAccountIncotermsService(data *AccountIncoterms) error { return UpdateAccountIncoterms(data) }
func DeleteAccountIncotermsService(id uint) error { return DeleteAccountIncoterms(id) }
