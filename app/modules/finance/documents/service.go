package documents

func CreateFinanceDocumentService(data *FinanceDocument) error {
	return CreateFinanceDocument(data)
}

func GetAllFinanceDocumentService() ([]FinanceDocument, error) {
	return GetAllFinanceDocument()
}

func GetPaginatedFinanceDocumentService(offset int, limit int, search string) ([]FinanceDocument, int64, error) {
	return GetPaginatedFinanceDocuments(offset, limit, search)
}

func GetFinanceDocumentByIDService(id uint) (*FinanceDocument, error) {
	return GetFinanceDocumentByID(id)
}

func UpdateFinanceDocumentService(data *FinanceDocument) error {
	return UpdateFinanceDocument(data)
}

func DeleteFinanceDocumentService(id uint) error {
	return DeleteFinanceDocument(id)
}
