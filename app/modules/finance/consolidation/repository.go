package consolidation

import (
	"gorm.io/gorm/clause"
	"ERP-System/config"
)

func CreateConsolidationReport(data *ConsolidationReport) error {
	return config.DB.Create(data).Error
}

func GetAllConsolidationReports() ([]ConsolidationReport, error) {
	var list []ConsolidationReport
	err := config.DB.Preload(clause.Associations).Find(&list).Error
	return list, err
}
