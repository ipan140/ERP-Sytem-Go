package expenses

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateExpense(data *Expense) error {
	return config.DB.Create(data).Error
}

func GetAllExpense() ([]Expense, error) {
	var list []Expense
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedExpenses(offset int, limit int, search string, status string) ([]Expense, int64, error) {
	var list []Expense
	var total int64

	query := config.DB.Model(&Expense{}).Preload(clause.Associations)

	if status != "" && status != "all" && status != "All" && status != "Semua" {
		query = query.Where("finance.expenses.state = ?", status)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = finance.expenses.employee_id").
			Where("finance.expenses.name ILIKE ? OR hrd.employees.name ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("finance.expenses.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetExpenseByID(id uint) (*Expense, error) {
	var data Expense
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}

func UpdateExpense(data *Expense) error {
	return config.DB.Save(data).Error
}

func DeleteExpense(id uint) error {
	return config.DB.Delete(&Expense{}, id).Error
}

func CreateExpenseSheet(data *ExpenseSheet) error { return config.DB.Create(data).Error }
func GetAllExpenseSheet() ([]ExpenseSheet, error) {
	var list []ExpenseSheet
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}

func GetPaginatedExpenseSheets(offset int, limit int, search string, status string) ([]ExpenseSheet, int64, error) {
	var list []ExpenseSheet
	var total int64

	query := config.DB.Model(&ExpenseSheet{}).Preload(clause.Associations)

	if status != "" && status != "all" && status != "All" && status != "Semua" {
		query = query.Where("finance.expense_sheets.state = ?", status)
	}
	if search != "" {
		s := "%" + search + "%"
		query = query.Joins("LEFT JOIN hrd.employees ON hrd.employees.id = finance.expense_sheets.employee_id").
			Where("finance.expense_sheets.name ILIKE ? OR hrd.employees.name ILIKE ?", s, s)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("finance.expense_sheets.id DESC").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func GetExpenseSheetByID(id uint) (*ExpenseSheet, error) {
	var data ExpenseSheet
	err := config.DB.Preload(clause.Associations).First(&data, id).Error
	return &data, err
}
func UpdateExpenseSheet(data *ExpenseSheet) error { return config.DB.Save(data).Error }
func DeleteExpenseSheet(id uint) error            { return config.DB.Delete(&ExpenseSheet{}, id).Error }
