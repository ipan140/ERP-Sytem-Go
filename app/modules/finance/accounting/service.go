package accounting

func CreateJournalEntryService(data *JournalEntry) error {
	return CreateJournalEntry(data)
}

func GetAllJournalEntryService() ([]JournalEntry, error) {
	return GetAllJournalEntry()
}

func GetJournalEntryByIDService(id uint) (*JournalEntry, error) {
	return GetJournalEntryByID(id)
}

func UpdateJournalEntryService(data *JournalEntry) error {
	return UpdateJournalEntry(data)
}

func DeleteJournalEntryService(id uint) error {
	return DeleteJournalEntry(id)
}
