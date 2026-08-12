package consolidation

import (
	"ERP-System/config"
)

func CreateConsolidationEntry(data *ConsolidationEntry) error {
	return config.DB.Create(data).Error
}

func GetAllConsolidationEntry() ([]ConsolidationEntry, error) {
	var list []ConsolidationEntry
	err := config.DB.Find(&list).Error
	return list, err
}

func GetConsolidationEntryByID(id uint) (*ConsolidationEntry, error) {
	var data ConsolidationEntry
	err := config.DB.First(&data, id).Error
	return &data, err
}

func UpdateConsolidationEntry(data *ConsolidationEntry) error {
	return config.DB.Save(data).Error
}

func DeleteConsolidationEntry(id uint) error {
	return config.DB.Delete(&ConsolidationEntry{}, id).Error
}
