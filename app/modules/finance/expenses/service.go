package expenses

func CreateExpenseService(data *Expense) error {
	return CreateExpense(data)
}

func GetAllExpenseService() ([]Expense, error) {
	return GetAllExpense()
}

func GetExpenseByIDService(id uint) (*Expense, error) {
	return GetExpenseByID(id)
}

func UpdateExpenseService(data *Expense) error {
	return UpdateExpense(data)
}

func DeleteExpenseService(id uint) error {
	return DeleteExpense(id)
}
