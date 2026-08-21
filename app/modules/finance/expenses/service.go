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

func CreateExpenseSheetService(data *ExpenseSheet) error        { return CreateExpenseSheet(data) }
func GetAllExpenseSheetService() ([]ExpenseSheet, error)        { return GetAllExpenseSheet() }
func GetExpenseSheetByIDService(id uint) (*ExpenseSheet, error) { return GetExpenseSheetByID(id) }
func UpdateExpenseSheetService(data *ExpenseSheet) error        { return UpdateExpenseSheet(data) }
func DeleteExpenseSheetService(id uint) error                   { return DeleteExpenseSheet(id) }
