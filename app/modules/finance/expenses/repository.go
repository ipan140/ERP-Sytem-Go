package expenses

import (
	"ERP-System/config"
)

func CreateExpense(data *Expense) error {
	return config.DB.Create(data).Error
}

func GetAllExpense() ([]Expense, error) {
	var list []Expense
	err := config.DB.Find(&list).Error
	return list, err
}

func GetExpenseByID(id uint) (*Expense, error) {
	var data Expense
	err := config.DB.First(&data, id).Error
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
	err := config.DB.Find(&list).Error
	return list, err
}
func GetExpenseSheetByID(id uint) (*ExpenseSheet, error) {
	var data ExpenseSheet
	err := config.DB.First(&data, id).Error
	return &data, err
}
func UpdateExpenseSheet(data *ExpenseSheet) error { return config.DB.Save(data).Error }
func DeleteExpenseSheet(id uint) error            { return config.DB.Delete(&ExpenseSheet{}, id).Error }
