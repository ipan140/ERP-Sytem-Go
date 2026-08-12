package invoicing

func CreateInvoiceService(data *Invoice) error {
	return CreateInvoice(data)
}

func GetAllInvoiceService() ([]Invoice, error) {
	return GetAllInvoice()
}

func GetInvoiceByIDService(id uint) (*Invoice, error) {
	return GetInvoiceByID(id)
}

func UpdateInvoiceService(data *Invoice) error {
	return UpdateInvoice(data)
}

func DeleteInvoiceService(id uint) error {
	return DeleteInvoice(id)
}
