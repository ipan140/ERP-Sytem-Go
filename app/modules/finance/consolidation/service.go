package consolidation

func CreateConsolidationEntryService(data *ConsolidationEntry) error {
	return CreateConsolidationEntry(data)
}

func GetAllConsolidationEntryService() ([]ConsolidationEntry, error) {
	return GetAllConsolidationEntry()
}

func GetConsolidationEntryByIDService(id uint) (*ConsolidationEntry, error) {
	return GetConsolidationEntryByID(id)
}

func UpdateConsolidationEntryService(data *ConsolidationEntry) error {
	return UpdateConsolidationEntry(data)
}

func DeleteConsolidationEntryService(id uint) error {
	return DeleteConsolidationEntry(id)
}
