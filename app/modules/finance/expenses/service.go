package expenses

import (
	"ERP-System/pkg/rabbitmq"
	"encoding/json"
	"log"
)

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

func ExportExpensesExcelService(period string, userID uint) error {
	if rabbitmq.Channel != nil {
		body, _ := json.Marshal(map[string]interface{}{"action": "export_expenses", "period": period, "user_id": userID})
		_ = rabbitmq.PublishEvent(rabbitmq.Channel, "finance_report_generator", body)
		log.Printf("💸 Event RabbitMQ: Export Excel Klaim Biaya %s dikirim ke antrean!", period)
	}
	return nil
}
