package reconciliation

import (
	"ERP-System/config"
)

func GetAllBankStatementsRepo() ([]BankStatementItem, error) {
	var list []BankStatementItem
	err := config.DB.Order("date desc, id desc").Find(&list).Error
	return list, err
}

func GetBankStatementByIDRepo(id uint) (*BankStatementItem, error) {
	var item BankStatementItem
	err := config.DB.First(&item, id).Error
	return &item, err
}

func CreateBankStatementRepo(item *BankStatementItem) error {
	return config.DB.Create(item).Error
}

func UpdateBankStatementRepo(item *BankStatementItem) error {
	return config.DB.Save(item).Error
}

func DeleteBankStatementRepo(id uint) error {
	return config.DB.Delete(&BankStatementItem{}, id).Error
}

func GetUnreconciledStatementsRepo() ([]BankStatementItem, error) {
	var list []BankStatementItem
	err := config.DB.Where("is_reconciled = ?", false).Find(&list).Error
	return list, err
}
