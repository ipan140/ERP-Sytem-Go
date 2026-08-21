package accounting

import (
	"ERP-System/config"
)

func CreateJournalEntry(data *JournalEntry) error {
	return config.DB.Create(data).Error
}

func GetAllJournalEntry() ([]JournalEntry, error) {
	var list []JournalEntry
	err := config.DB.Find(&list).Error
	return list, err
}

func GetJournalEntryByID(id uint) (*JournalEntry, error) {
	var data JournalEntry
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateJournalEntry(data *JournalEntry) error {
	return config.DB.Save(data).Error
}

func DeleteJournalEntry(id uint) error {
	return config.DB.Delete(&JournalEntry{}, id).Error
}

func CreateJournalItem(data *JournalItem) error {
	return config.DB.Create(data).Error
}

func CreateAccountReconcileModel(data *AccountReconcileModel) error {
	return config.DB.Create(data).Error
}
func GetAllAccountReconcileModel() ([]AccountReconcileModel, error) {
	var list []AccountReconcileModel
	err := config.DB.Find(&list).Error
	return list, err
}
func GetAccountReconcileModelByID(id uint) (*AccountReconcileModel, error) {
	var data AccountReconcileModel
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateAccountReconcileModel(data *AccountReconcileModel) error {
	return config.DB.Save(data).Error
}
func DeleteAccountReconcileModel(id uint) error {
	return config.DB.Delete(&AccountReconcileModel{}, id).Error
}

func CreateFollowupRule(data *FollowupRule) error { return config.DB.Create(data).Error }
func GetAllFollowupRule() ([]FollowupRule, error) {
	var list []FollowupRule
	err := config.DB.Find(&list).Error
	return list, err
}
func GetFollowupRuleByID(id uint) (*FollowupRule, error) {
	var data FollowupRule
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateFollowupRule(data *FollowupRule) error { return config.DB.Save(data).Error }
func DeleteFollowupRule(id uint) error            { return config.DB.Delete(&FollowupRule{}, id).Error }

func CreateAccountLockDate(data *AccountLockDate) error { return config.DB.Create(data).Error }
func GetAllAccountLockDate() ([]AccountLockDate, error) {
	var list []AccountLockDate
	err := config.DB.Find(&list).Error
	return list, err
}
func GetAccountLockDateByID(id uint) (*AccountLockDate, error) {
	var data AccountLockDate
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateAccountLockDate(data *AccountLockDate) error { return config.DB.Save(data).Error }
func DeleteAccountLockDate(id uint) error               { return config.DB.Delete(&AccountLockDate{}, id).Error }

func CreatePaymentAcquirer(data *PaymentAcquirer) error { return config.DB.Create(data).Error }
func GetAllPaymentAcquirer() ([]PaymentAcquirer, error) {
	var list []PaymentAcquirer
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPaymentAcquirerByID(id uint) (*PaymentAcquirer, error) {
	var data PaymentAcquirer
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePaymentAcquirer(data *PaymentAcquirer) error { return config.DB.Save(data).Error }
func DeletePaymentAcquirer(id uint) error               { return config.DB.Delete(&PaymentAcquirer{}, id).Error }

func CreatePaymentTransaction(data *PaymentTransaction) error { return config.DB.Create(data).Error }
func GetAllPaymentTransaction() ([]PaymentTransaction, error) {
	var list []PaymentTransaction
	err := config.DB.Find(&list).Error
	return list, err
}
func GetPaymentTransactionByID(id uint) (*PaymentTransaction, error) {
	var data PaymentTransaction
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdatePaymentTransaction(data *PaymentTransaction) error { return config.DB.Save(data).Error }
func DeletePaymentTransaction(id uint) error {
	return config.DB.Delete(&PaymentTransaction{}, id).Error
}

func CreateAccountIncoterms(data *AccountIncoterms) error { return config.DB.Create(data).Error }
func GetAllAccountIncoterms() ([]AccountIncoterms, error) {
	var list []AccountIncoterms
	err := config.DB.Find(&list).Error
	return list, err
}
func GetAccountIncotermsByID(id uint) (*AccountIncoterms, error) {
	var data AccountIncoterms
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateAccountIncoterms(data *AccountIncoterms) error { return config.DB.Save(data).Error }
func DeleteAccountIncoterms(id uint) error                { return config.DB.Delete(&AccountIncoterms{}, id).Error }
