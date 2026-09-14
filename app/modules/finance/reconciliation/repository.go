package reconciliation

import (
	"ERP-System/config"
)

func GetAllBankStatementsRepo() ([]BankStatementItem, error) {
	var list []BankStatementItem
	err := config.DB.Order("date desc, id desc").Find(&list).Error
	return list, err
}

func GetPaginatedBankStatementsRepo(offset int, limit int, search string, bank string, status string) ([]BankStatementItem, int64, error) {
	var list []BankStatementItem
	var total int64

	query := config.DB.Model(&BankStatementItem{})

	if bank != "" && bank != "all" && bank != "All" && bank != "Semua" {
		query = query.Where("bank_name = ?", bank)
	}

	if status == "reconciled" {
		query = query.Where("is_reconciled = ?", true)
	} else if status == "unreconciled" {
		query = query.Where("is_reconciled = ?", false)
	}

	if search != "" {
		s := "%" + search + "%"
		query = query.Where("description ILIKE ? OR ref_number ILIKE ? OR matched_invoice ILIKE ?", s, s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("date desc, id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
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
