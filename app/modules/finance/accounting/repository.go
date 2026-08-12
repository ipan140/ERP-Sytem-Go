package accounting

import (
	"ERP-System/config"
)

func CreateJournalEntry(data *JournalEntry) error {
	return config.DB.Create(data).Error
}

func GetAllJournalEntry() ([]JournalEntry, error) {
	var list []JournalEntry
	err := config.DB.Find(&list).Error
	return list, err
}

func GetJournalEntryByID(id uint) (*JournalEntry, error) {
	var data JournalEntry
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateJournalEntry(data *JournalEntry) error {
	return config.DB.Save(data).Error
}

func DeleteJournalEntry(id uint) error {
	return config.DB.Delete(&JournalEntry{}, id).Error
}
