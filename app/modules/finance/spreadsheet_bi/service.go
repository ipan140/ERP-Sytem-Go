package spreadsheet_bi

func CreateSpreadsheetService(data *Spreadsheet) error {
	return CreateSpreadsheet(data)
}

func GetAllSpreadsheetService() ([]Spreadsheet, error) {
	return GetAllSpreadsheet()
}

func GetPaginatedSpreadsheetService(offset int, limit int, search string) ([]Spreadsheet, int64, error) {
	return GetPaginatedSpreadsheets(offset, limit, search)
}

func GetSpreadsheetByIDService(id uint) (*Spreadsheet, error) {
	return GetSpreadsheetByID(id)
}

func UpdateSpreadsheetService(data *Spreadsheet) error {
	return UpdateSpreadsheet(data)
}

func DeleteSpreadsheetService(id uint) error {
	return DeleteSpreadsheet(id)
}
